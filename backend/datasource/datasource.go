package datasource

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	dataTable *sql.DB
)

func init() {
	dataTable, err := setupDatasource()
	if err != nil {
		log.Fatal(err)
	}
}

const (
	databasePath = "."
)

func GetArticleWithID(id string) string {
	return "das ist unser toller artikel!"
}

func setupDatasource() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/%s", databasePath, sqliteArticleTableDefinition.name))
	if err != nil {
		return nil, err
	}

	statement := sqliteArticleTableDefinition.getSqliteTableCreationStatement()

	s, err := db.Prepare(statement)
	if err != nil {
		return nil, err
	}

	_, err = s.Exec()
	return db, err
}
