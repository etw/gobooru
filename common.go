package gobooru

import (
	"net/http"
)

type API struct {
	httpClient *http.Client
	format     string
}

func New(c *http.Client, f string) *API {
	api := new(API)
	api.httpClient = c
	api.format = f
	return api
}
