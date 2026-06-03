package smtpserver

import (
	"fmt"
	"io"
	"log"

	"github.com/emersion/go-smtp"

	"phishing-platform/internal/parser"
	"phishing-platform/internal/pipeline"
)

type Backend struct{}

func (bkd *Backend) NewSession(
	c *smtp.Conn,
) (smtp.Session, error) {

	return &Session{}, nil
}

type Session struct{}

func (s *Session) AuthPlain(
	username,
	password string,
) error {

	return nil
}

func (s *Session) Mail(
	from string,
	opts *smtp.MailOptions,
) error {

	fmt.Println(
		"Mail From:",
		from,
	)

	return nil
}

func (s *Session) Rcpt(
	to string,
	opts *smtp.RcptOptions,
) error {

	fmt.Println(
		"Rcpt To:",
		to,
	)

	return nil
}

func (s *Session) Data(
	r io.Reader,
) error {

	parsedEmail, err :=
		parser.ParseRawEmail(r)

	if err != nil {

		fmt.Println(
			"Email Parse Error:",
			err,
		)

		return err
	}

	fmt.Println(
		"Parsed Email:",
	)

	fmt.Println(
		"From:",
		parsedEmail.From,
	)

	fmt.Println(
		"Subject:",
		parsedEmail.Subject,
	)

	fmt.Println(
		"Body:",
		parsedEmail.Body,
	)

	fmt.Println(
	"Attachments:",
	parsedEmail.Attachments,
)

	pipeline.ProcessEmail(
	parsedEmail.From,
	parsedEmail.ReplyTo,
	parsedEmail.ReturnPath,
	parsedEmail.Subject,
	parsedEmail.Body,
	parsedEmail.Attachments,
)

	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {

	return nil
}

func StartSMTPServer() {

	backend := &Backend{}

	server := smtp.NewServer(
		backend,
	)

	server.Addr = ":2525"

	server.Domain = "localhost"

	server.AllowInsecureAuth = true

	fmt.Println(
		"SMTP Server running on port 2525",
	)

	log.Fatal(
		server.ListenAndServe(),
	)
}