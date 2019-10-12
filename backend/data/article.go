package data

type ArticleType int

const (
	Id ArticleType = iota
	Author
	Title
	Content
	InDraft
	Date
	Public
)

var (
	sortedArticleTypes = []ArticleType{Id, Author, Title, Content, InDraft, Date, Public}
)

func (t ArticleType) String() string {
	return map[ArticleType]string{
		Id:      "id",
		Author:  "author",
		Title:   "title",
		Content: "content",
		InDraft: "indraft",
		Date:    "date",
		Public:  "public",
	}[t]
}

type article map[ArticleType]string

func NewArticle() *article {
	return &article{
		Id:      "",
		Author:  "",
		Title:   "",
		Content: "",
		InDraft: "",
		Date:    "",
		Public:  "",
	}
}

func NewArticleWithArray(array []string) *article {
	a := NewArticle()
	for i, typ := range sortedArticleTypes {
		(*a)[typ] = array[i]
	}
	return a
}

func getArticleLength() int {
	return len(sortedArticleTypes)
}

func (a *article) getValuesArticle() []string {
	ret := make([]string, 0, getArticleLength())
	for _, typ := range sortedArticleTypes {
		ret = append(ret, (*a)[typ])
	}

	return ret
}

func getTypesArticle() []string {
	ret := make([]string, 0, getArticleLength())
	for _, typ := range sortedArticleTypes {
		ret = append(ret, typ.String())
	}

	return ret
}
