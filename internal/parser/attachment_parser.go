package parser

import (
	"mime"
)

func ExtractAttachmentNames(
	contentType string,
) []string {

	attachments := []string{}

	_, params, err :=
		mime.ParseMediaType(
			contentType,
		)

	if err != nil {
		return attachments
	}

	filename := params["filename"]

	if filename != "" {

		attachments = append(
			attachments,
			filename,
		)
	}

	return attachments
}