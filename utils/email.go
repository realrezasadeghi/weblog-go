package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/fs"
	"path/filepath"
	"weblog/configs"
	"weblog/models"
)

func ParseTemplateDirectory(dir string) (*template.Template, error) {
	var paths []string

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(user *models.User, data *models.EmailData, email string) error {
	config, err := configs.LoadSmtpConfigs()

	if err != nil {
		fmt.Println("[SendEmail] could not load env smtp files", err.Error())
		return err
	}

	// sender email
	senderEmail := map[string]string{
		"to":           user.Email,
		"smtpPort":     config.SMTPPort,
		"smtpUser":     config.SMTPUser,
		"smtpHost":     config.SMTPHost,
		"from":         config.EmailFrom,
		"smtpPassword": config.SMTPPassword,
	}

	// body templateDirectory html
	var body bytes.Buffer

	templateDirectory, errParseTemplate := ParseTemplateDirectory("templates")

	if errParseTemplate != nil {
		fmt.Println("[SendEmail] error in parse templateDirectory directory", errParseTemplate.Error())
		return errParseTemplate
	}

	errExecuteTemplate := templateDirectory.ExecuteTemplate(&body, email, &data)

	if errExecuteTemplate != nil {
		fmt.Println("[SendEmail] error in execute template", errExecuteTemplate.Error())
		return errExecuteTemplate
	}

	// receive email
	messageEmail := gomail.NewMessage()

	messageEmail.SetHeader("To", senderEmail["to"])
	messageEmail.SetHeader("Subject", data.Subject)
	messageEmail.SetHeader("From", senderEmail["from"])
	messageEmail.SetBody("text/html", body.String())
	messageEmail.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	dailer := gomail.NewDialer(senderEmail["smtpHost"], ToInt(senderEmail["smtpPort"]), senderEmail["smtpUser"], senderEmail["smtpPassword"])

	dailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dailer.DialAndSend(messageEmail); err != nil {
		fmt.Println("[SendEmail] could not send email", err.Error())
		return err
	}

	return nil
}
