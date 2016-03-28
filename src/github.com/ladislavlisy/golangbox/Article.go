package golangbox

type Article struct {
	code uint32
	catg uint32
	pendings []uint32
}

func NewArticle(code uint32, catg uint32, pendings []uint32) *Article {
	return &Article{code, catg, pendings}
}

