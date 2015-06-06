package gobooru

type Post struct {
	Height   int
	Width    int
	Url      string
	Sample   string
	Preview  string
	Rating   string
	Id       int
	Tags     string
	Comments []Comment
}

type Comment struct {
	Id     int
	PostId int
	Body   string
}

type cursor struct {
	page int
	pos  int
}

type Booru interface {
	GetById(int) (*Post, error)

	GetByTags([]string, int, int) ([]Post, error)
	CurByTags([]string) (func(int) ([]Post, error), error)

	GetComm(int) ([]Comment, error)
	CurComm(int) (func(int) ([]Comment, error), error)
}
