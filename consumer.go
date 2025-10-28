package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch <-chan Receipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for receipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"
		formattedMessage := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", receipient.Email, "Just testing email campaign")
		msg := []byte(formattedMessage)

		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "chandra@codersguy.com", []string{receipient.Email}, msg) // in built go package

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(1 * time.Second) // just a rate limiter to not make service reach limit
		fmt.Printf("Worker %d: Sending email to %s \n", id, receipient.Email)
	}
}
