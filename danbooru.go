package gobooru

import (
	"net/http"
)

const (
	DANBOORU = "https://danbooru.donmai.us"
	KONACHAN = "http://konachan.com"
	SANKAKU  = "https://chan.sankakucomplex.com"
)

type DbAPI struct {
	httpClient *http.Client
	prefix     string
}
