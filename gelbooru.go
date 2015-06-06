package gobooru

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const GELBOORU = "http://gelbooru.com/index.php?page=dapi&q=index"

type GbAPI struct {
	httpClient *http.Client
	prefix     string
}

type GbPosts struct {
	XMLName xml.Name `xml:"posts"`
	Count   int      `xml:"count,attr"`
	Offset  int      `xml:"offset,attr"`
	List    []GbPost `xml:"post"`
}

type GbPost struct {
	Height        int    `xml:"height,attr"`
	Width         int    `xml:"width,attr"`
	ParentId      int    `xml:"parent_id,attr"`
	FileUrl       string `xml:"file_url,attr"`
	SampleUrl     string `xml:"sample_url,attr"`
	SampleHeight  int    `xml:"sample_height,attr"`
	SampleWidth   int    `xml:"sample_widtht,attr"`
	Score         int    `xml:"score,attr"`
	PreviewUrl    string `xml:"preview_url,attr"`
	PreviewHeight int    `xml:"preview_height,attr"`
	PreviewWidth  int    `xml:"preview_width,attr"`
	Rating        string `xml:"rating,attr"`
	Id            int    `xml:"id,attr"`
	Tags          string `xml:"tags,attr"`
	Change        int    `xml:"change,attr"`
	Md5           string `xml:"md5,attr"`
	CreatorId     int    `xml:"creator_id,attr"`
	CreatedAt     string `xml:"created_at,attr"`
	Status        string `xml:"status,attr"`
	Source        string `xml:"source,attr"`
	HasNotes      bool   `xml:"has_notes,attr"`
	HasComments   bool   `xml:"has_comments,attr"`
	HasChildren   bool   `xml:"has_children,attr"`
}

type GbComments struct {
	XMLName xml.Name    `xml:"comments"`
	Type    string      `xml:"type,attr"`
	List    []GbComment `xml:"comment"`
}

type GbComment struct {
	CreatedAt time.Time `xml:"created_at,attr"`
	PostId    int       `xml:"post_id,attr"`
	Body      string    `xml:"body,attr"`
	Creator   string    `xml:"creator,attr"`
	Id        int       `xml:"id,attr"`
	CreatorId int       `xml:"creator_id,attr"`
}

func NewGb(c *http.Client, p string) *GbAPI {
	api := new(GbAPI)
	api.httpClient = c
	api.prefix = p
	return api
}

func (api *GbAPI) metaGet(u *string) ([]byte, error) {
	if req, err := http.NewRequest("GET", *u, nil); err != nil {
		return nil, err
	} else {
		if resp, err := api.httpClient.Do(req); err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return nil, errors.New(resp.Status)
			}

			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				return nil, err
			} else {
				return body, nil
			}
		}
	}
}

func (api *GbAPI) GetByIdRaw(id int) (p *GbPosts, e error) {
	p = new(GbPosts)

	path := fmt.Sprintf("%s&s=post&id=%d", api.prefix, id)

	defer func() {
		if r := recover(); r != nil {
			p, e = nil, errors.New(fmt.Sprintf("Unknown error while getting %s", path))
		}
	}()

	if data, err := api.metaGet(&path); err != nil {
		return nil, err
	} else {
		if err := xml.Unmarshal(data, p); err != nil {
			return nil, err
		} else {
			return p, nil
		}
	}
}

func (api *GbAPI) GetByTagsRaw(t []string, n int) (p *GbPosts, e error) {
	p = new(GbPosts)

	path := fmt.Sprintf("%s&s=post&tags=%s", api.prefix, strings.Join(t, " "))

	defer func() {
		if r := recover(); r != nil {
			p, e = nil, errors.New(fmt.Sprintf("Unknown error while getting %s", path))
		}
	}()

	if data, err := api.metaGet(&path); err != nil {
		return nil, err
	} else {
		if err := xml.Unmarshal(data, p); err != nil {
			return nil, err
		} else {
			return p, nil
		}
	}
}

func (api *GbAPI) GetCommRaw(id int) (p *GbComments, e error) {
	p = new(GbComments)

	path := fmt.Sprintf("%s&s=comment&post_id=%d", api.prefix, id)

	defer func() {
		if r := recover(); r != nil {
			p, e = nil, errors.New(fmt.Sprintf("Unknown error while getting %s", path))
		}
	}()

	if data, err := api.metaGet(&path); err != nil {
		return nil, err
	} else {
		if err := xml.Unmarshal(data, p); err != nil {
			return nil, err
		} else {
			return p, nil
		}
	}
}

func (api *GbAPI) GetById(id int) (*Post, error) {
	if tmp, err := api.GetByIdRaw(id); err != nil {
		return nil, err
	} else {
		if len(tmp.List) == 1 {
			return &Post{
				Height:   tmp.List[0].Height,
				Width:    tmp.List[0].Width,
				Url:      tmp.List[0].FileUrl,
				Sample:   tmp.List[0].SampleUrl,
				Preview:  tmp.List[0].PreviewUrl,
				Rating:   tmp.List[0].Rating,
				Id:       tmp.List[0].Id,
				Tags:     tmp.List[0].Tags,
				Comments: []Comment{},
			}, nil
		} else {
			return nil, errors.New("No posts")
		}
	}
}
