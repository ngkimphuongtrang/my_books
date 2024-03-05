package main

import (
	"crypto/tls"
	"github.com/emersion/go-imap"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

func main() {
	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", &tls.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("trang.nguyen3@trustingsocial.com", "lxxlfrlturopceti"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Select INBOX
	_, err = c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	// Search for unseen messages
	criteria := imap.NewSearchCriteria()
	//criteria.WithoutFlags = []string{imap.SeenFlag}
	criteria.Since = time.Date(2024, 03, 01, 0, 0, 0, 0, time.Local)
	criteria.Text = []string{"Grab"}

	seqNums, err := c.Search(criteria)
	if err != nil {
		log.Fatal(err)
	}

	if len(seqNums) == 0 {
		log.Println("No new unseen emails.")
		return
	}

	seqSet := new(imap.SeqSet)
	seqSet.AddNum(seqNums...)

	// Get the whole message body
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal(err)
		}
	}()

	for msg := range messages {
		if msg == nil {
			log.Fatal("Server didn't returned message")
		}

		r := msg.GetBody(section)
		if r == nil {
			log.Fatal("Server didn't returned message body")
		}

		// Create a new mail reader
		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Fatal(err)
		}

		// Process each message's part
		// Check if the message is multipart
		if mr.Header.Get("Content-Type") == "multipart/mixed" {
			log.Println("ok")
			// Process each part
			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break // No more parts
				}
				if err != nil {
					log.Fatal(err)
				}

				// Check the content type of the part
				contentType := p.Header.Get("Content-Type")

				if contentType == "image/jpeg" || contentType == "image/png" {
					// We found an image
					filename := p.Header.Get("file-name")
					if filename == "" {
						filename = "image_attachment" // You might want to generate a unique filename
					}
					log.Printf("Found an image: %v\n", filename)

					// Save the image
					body, _ := ioutil.ReadAll(p.Body)
					err = ioutil.WriteFile(filename, body, 0644)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("Saved image %v to disk.\n", filename)
				}
			}
		}

	}

	log.Println("Done!")
}
