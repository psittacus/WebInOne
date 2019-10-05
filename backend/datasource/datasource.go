package datasource

import (
	_ "github.com/mattn/go-sqlite3"
)

const (
	tableName    = "blogTable"
	databasePath = "."
)

func GetArticleWithID(id string) string {
	return "das ist unser toller artikel!"
}
