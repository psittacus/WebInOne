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

// TODO: test if statement creates database
func (t *table) getSqliteTableCreationStatement() string {
	tableDefinition := fmt.Sprintf("CREATE TABLE IF NOT EXIST %s (", t.name)
	for _, def := range t.columnDefinitions {
		tableDefinition += fmt.Sprintf("%s %s,", def[0], def[1])
	}

	tableDefinition += ")"
	return tableDefinition
}
