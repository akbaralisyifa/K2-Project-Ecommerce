package utils

import (
	"bytes"
	"ecommerce/config"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)


type DataEmployee struct{
	Name string
	
}

func SendGomail(templatePath string) {
	PwMail := config.ImportSetting().PwMail

	// Mendapatkan HTML
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		// Menangani kesalahan setelah parsing template
		fmt.Println("Error parsing template:", err)
		return
	}

	// Mengeksekusi template jika parsing berhasil
	if t != nil {
		err = t.Execute(&body, struct{Name string}{Name: "Employee Name"})
		if err != nil {
			// Menangani kesalahan setelah mengeksekusi template
			fmt.Println("Error executing template:", err)
			return
		}
	} else {
		fmt.Println("Template is nil after parsing")
		return
	}

	// Mengirim dengan GoMail
	m := gomail.NewMessage()
	m.SetHeader("From", "akbaralisyifa@gmail.com")
	m.SetHeader("To", "akbaralisyifa22@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "akbaralisyifa22@gmail.com", PwMail)

	// Mengirim email
	if err := d.DialAndSend(m); err != nil {
		// Menangani kesalahan pengiriman email
		fmt.Println("Error sending email:", err)
		panic(err)
	}
}
