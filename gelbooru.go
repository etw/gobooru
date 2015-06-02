package gelbooru

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const apiPrefix = "http://gelbooru.com/index.php?page=dapi&s=post&q=index&id="

type API struct {
	httpClient *http.Client
	auth       *string
}

type Posts struct {
	Count  int `xml:"count,attr"`
	Offset int `xml:"offset,attr"`
	List   []Post
}

type Post struct {
	Height        int       `xml:"height,attr"`
	Width         int       `xml:"width,attr"`
	ParentId      int       `xml:"parent_id,attr"`
	FileUrl       url.URL   `xml:"file_url,attr"`
	SampleUrl     url.URL   `xml:"sample_url,attr"`
	SampleHeight  int       `xml:"sample_height,attr"`
	SampleWidth   int       `xml:"sample_widtht,attr"`
	Score         int       `xml:"score,attr"`
	PreviewUrl    url.URL   `xml:"preview_url,attr"`
	PreviewHeight int       `xml:"preview_height,attr"`
	PreviewWidth  int       `xml:"preview_width,attr"`
	Rating        string    `xml:"rating,attr"`
	Id            int       `xml:"id,attr"`
	Tags          []string  `xml:"tags,attr"`
	Change        time.Time `xml:"change,attr"`
	Md5           string    `xml:"md5,attr"`
	CreatorId     int       `xml:"creator_id,attr"`
	CreatedAt     time.Time `xml:"created_at,attr"`
	Status        string    `xml:"status,attr"`
	Source        url.URL   `xml:"source,attr"`
	HasNotes      bool      `xml:"has_notes,attr"`
	HasComments   bool      `xml:"has_comments,attr"`
	HasChildren   bool      `xml:"has_children,attr"`
}

func New(c *http.Client) *API {
	api := new(API)
	api.httpClient = c
	return api
}

func (api *API) GetPics(id *string) (*Posts, error) {
	var p Posts

	req, err := http.NewRequest("GET", fmt.Sprint(apiPrefix, *id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
