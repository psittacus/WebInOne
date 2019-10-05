package datasource

type table struct {
	tableName   string
	definitions [][]string
}

var (
	sqliteArticleTableDefinition = table{
		tableName: "DataBlog",
		definitions: [][]string{
			{"id", "INTEGER PRIMARY KEY"},
			{"author", "TEXT NOT NULL"},
			{"title", "TEXT NOT NULL"},
			{"content", "TEXT"},
			{"inDraft", "INTEGER DEFAULT 1 NOT NULL"},
			{"date", "DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL"},
			{"public", "INTEGER DEFAULT 1"}}}
)

func (t *table) GetDefinition() (string){
	var ret string
	for _, def := range table.defintions[1:] {
		ret +=
	}
}
