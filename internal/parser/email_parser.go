package parser

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/mail"
	"os"
	"path/filepath"
)

type ParsedEmail struct {
	From        string
	ReplyTo     string
	ReturnPath  string
	Subject     string
	Body        string
	Attachments []string
}

func ParseRawEmail(
	r io.Reader,
) (*ParsedEmail, error) {

	buf := new(bytes.Buffer)

	_, err := buf.ReadFrom(r)

	if err != nil {
		return nil, err
	}

	rawEmail := buf.String()

	fmt.Println("===================================")
	fmt.Println("RAW EMAIL START")
	fmt.Println(rawEmail)
	fmt.Println("RAW EMAIL END")
	fmt.Println("===================================")

	msg, err := mail.ReadMessage(
		bytes.NewReader([]byte(rawEmail)),
	)

	if err != nil {
		return nil, err
	}

	contentType := msg.Header.Get(
		"Content-Type",
	)

	var attachments []string
	var body string

	mediaType, params, err :=
		mime.ParseMediaType(
			contentType,
		)

	if err == nil &&
		mediaType == "multipart/mixed" {

		reader :=
			multipart.NewReader(
				msg.Body,
				params["boundary"],
			)

		for {

			part, err :=
				reader.NextPart()

			if err != nil {
				break
			}

			filename := part.FileName()

if filename != "" {

	os.MkdirAll(
		"uploads",
		0755,
	)

	savePath :=
		filepath.Join(
			"uploads",
			filename,
		)

	fileData, err :=
		io.ReadAll(
			part,
		)

	if err == nil {

		os.WriteFile(
			savePath,
			fileData,
			0644,
		)

		fmt.Println(
			"Saved Attachment:",
			savePath,
		)

		attachments = append(
			attachments,
			savePath,
		)
	}

	continue
}

			partBytes := new(bytes.Buffer)

			_, _ = partBytes.ReadFrom(
				part,
			)

			body += partBytes.String()
		}

	} else {

		bodyBytes := new(bytes.Buffer)

		_, err = bodyBytes.ReadFrom(
			msg.Body,
		)

		if err != nil {
			return nil, err
		}

		body = bodyBytes.String()
	}

	parsed := &ParsedEmail{
	From:        msg.Header.Get("From"),
	ReplyTo:     msg.Header.Get("Reply-To"),
	ReturnPath:  msg.Header.Get("Return-Path"),
	Subject:     msg.Header.Get("Subject"),
	Body:        body,
	Attachments: attachments,
}
fmt.Println(
	"Reply-To:",
	parsed.ReplyTo,
)

fmt.Println(
	"Return-Path:",
	parsed.ReturnPath,
)

	return parsed, nil
}