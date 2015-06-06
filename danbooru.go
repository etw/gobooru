package gobooru

import (
	"net/http"
)

const (
	DANBOORU = "https://danbooru.donmai.us"
	KONACHAN = "http://konachan.com"
	SANKAKU  = "https://chan.sankakucomplex.com"
	YANDERE  = "https://yande.re"
)

type DbAPI struct {
	httpClient *http.Client
	prefix     string
}

func NewDb(c *http.Client, p string) *DbAPI {
	api := new(DbAPI)
	api.httpClient = c
	api.prefix = p
	return api
}
