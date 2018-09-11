// This self-contained program uses gomail to send email.
// Dependencies:
//   https://github.com/go-gomail/gomail

// How to use?
// IITB firewall blocks the SMTP request. It worked on Airtel connection.
// Make sure that in google account security dashboard you
// enable "Allow less secure apps:" to ON
// Also enable IMAP in the gmail settings (not sure if this is necessary -- it did work without it)

// Successfully Tested!!

package main

import "fmt"

//import "github.com/gopkg.in/gomail.v2"
import "github.com/go-gomail/gomail"

func sendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "itsoflife@gmail.com")
	m.SetHeader("To", "anshumandhuliya@gmail.com", "lazynintel@gmail.com")
	m.SetAddressHeader("Cc", "iampriyanka.dhuliya@gmail.com", "Phulki")
	m.SetHeader("Subject", "Hello There Again!")
	m.SetBody("text/html", "Hello <b>Neo</b> and <b>Phulki</b>! This email is from a golang code.")
	m.Attach("myfavicon.ico")

	d := gomail.NewDialer("smtp.gmail.com", 587, "itsoflife", "itsoflife008")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main() {
	sendEmail()
	fmt.Println("vim-go")
}
