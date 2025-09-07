package models

import "net/http"

type Client struct {
	PublicKey  string
	PrivateKey string
	BaseUrl    string
	HTTPClient *http.Client
}
