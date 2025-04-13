package main

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Messsage struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
}

func (m *Mail) SendSMTPMessage(msg Messsage) error {
	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}

	msg.DataMap = data

	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		log.Println("Error while Getting formattedMessage")
		return err
	}

	plainMessage, err := m.buildPlainTextMessage(msg)
	if err != nil {
		log.Println("Error while Getting buildPlainTextMessage")
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		// Add debug output before connection
		log.Printf("Attempting to connect to %s:%d with username %s", m.Host, m.Port, m.Username)
		log.Println("Error while starting SMTP server")
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)

	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, val := range msg.Attachments {
			email.AddAttachment(val)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Println("Error while sending email")
		return err
	}

	return nil
}

func (m *Mail) getEncryption(s string) mail.Encryption {
	switch s {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}

func (m *Mail) buildHTMLMessage(msg Messsage) (string, error) {

	templateToRender := "./templates/mail.html.gohtml"

	t, err := template.New("email-html").ParseFiles(templateToRender)

	if err != nil {
		log.Println("Error while reading template")
		return "", err
	}

	var tpl bytes.Buffer

	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		log.Println("Some Execute template error")
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		log.Println("Error while inlineCSS func ")
		return "", err
	}

	return formattedMessage, nil

}
func (m *Mail) buildPlainTextMessage(msg Messsage) (string, error) {

	templateToRender := "./templates/mail.plain.gohtml"

	t, err := template.New("email-plain").ParseFiles(templateToRender)

	if err != nil {
		log.Println("Error while reading template")
		return "", err
	}

	var tpl bytes.Buffer

	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		log.Println("Some Execute template error")
		return "", err
	}

	plainMessage := tpl.String()

	return plainMessage, nil

}

func (m *Mail) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		log.Println("Error while creating Premailer")
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		log.Println("Error while Prem tranform")
		return "", err
	}

	return html, nil

}
