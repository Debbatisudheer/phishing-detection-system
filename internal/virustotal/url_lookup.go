package virustotal

import (
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strings"
)

func QueryURL(
	url string,
) ([]byte, error) {

	apiKey :=
		os.Getenv(
			"VT_API_KEY",
		)

	urlID :=
		base64.RawURLEncoding.EncodeToString(
			[]byte(url),
		)

	apiURL :=
		"https://www.virustotal.com/api/v3/urls/" +
			strings.TrimSpace(urlID)

	req, err :=
		http.NewRequest(
			"GET",
			apiURL,
			nil,
		)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"x-apikey",
		apiKey,
	)

	client := &http.Client{}

	resp, err :=
		client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(
		resp.Body,
	)
}