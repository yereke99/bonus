package main

import (
	"gopkg.in/gomail.v2"
)

func main() {
	// Создаем новое сообщение
	m := gomail.NewMessage()

	// Указываем отправителя
	m.SetHeader("From", "erkinbekuly99@yandex.kz")

	// Указываем получателя
	m.SetHeader("To", "erkinbekly@gmail.com")

	// Указываем тему письма
	m.SetHeader("Subject", "Hello from Go")

	// Текст сообщения
	m.SetBody("text/plain", "This is a test email sent using Go and Yandex Mail.")

	// Настройки SMTP сервера Yandex
	d := gomail.NewDialer("smtp.yandex.com", 465, "erkinbekuly99@yandex.kz", "Apqp6bX45S8rw5A")
	d.SSL = true

	// Отправляем сообщение
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
