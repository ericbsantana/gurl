package parser

import (
	"net/url"
)

func ParseURL(urlToParse string) (string, string, string, error) {
	u, err := url.Parse(urlToParse)

	if err != nil {
		return "", "", "", err
	}

	host := u.Hostname()
	port := u.Port()
	path := u.Path

	if port == "" {
		port = "80"
	}

	return host, port, path, err
}
