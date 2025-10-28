package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Receipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannels := make(chan Receipient)
	// Producer will read emails from csv file.
	// And send to consumers

	workerCount := 5 // consumers you want to receive emails and process to
	go loadReceipient("./emails.csv", recipientChannels)
	wg := sync.WaitGroup{}
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannels, &wg)
	}
	wg.Wait()
}

func executeTemplate(r Receipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
