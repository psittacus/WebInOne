package datasource

import "fmt"

type table struct {
	name              string
	columnDefinitions [][]string
}

var (
	sqliteArticleTableDefinition = table{
		name: "DataBlog",
		columnDefinitions: [][]string{
			{"id", "INTEGER PRIMARY KEY"},
			{"author", "TEXT NOT NULL"},
			{"title", "TEXT NOT NULL"},
			{"content", "TEXT"},
			{"inDraft", "INTEGER DEFAULT 1 NOT NULL"},
			{"date", "DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL"},
			{"public", "INTEGER DEFAULT 1"}}}
)

func (t *table) getSqliteTableCreationStatement() string {
	tableDefinition := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", t.name)
	if len(t.columnDefinitions) > 0 {
		tableDefinition += fmt.Sprintf("%s %s", t.columnDefinitions[0][0], t.columnDefinitions[0][1])
		for _, def := range t.columnDefinitions[1:] {
			tableDefinition += fmt.Sprintf(", %s %s", def[0], def[1])
		}
	}

	tableDefinition += ")"
	return tableDefinition
}
