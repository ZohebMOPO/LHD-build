package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	toList := []string{"zohebkhan@zoheb.com"}

	host := "smtp.gmail.com"

	port := "587"

	msg := "Hacking in MLH Build :O"

	body := []byte(msg)

	auth := smtp.PlainAuth(" ", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail")
}
