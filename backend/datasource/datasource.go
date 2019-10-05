package datasource

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dataRootName       = "DataBlog"
	sqliteExtension    = "sqlite3"
	sqliteDatabasePath = "./" + dataRootName + sqliteExtension
)

type Source interface {
	Setup() error
	// GetArticleWithId(id string) Article
	GetArticleWithId(id string) string
}

type sqlite struct {
	path string
}

func NewSqlite() Source {
	return &sqlite{path: sqliteDatabasePath}
}

func (s *sqlite) GetArticleWithId(id string) string {
	return "das ist unser toller artikel!"
}

func (s *sqlite) Setup() error {
	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return err
	}
	defer db.Close()

	statement, err := db.Prepare(sqliteArticleTableDefinition.getSqliteTableCreationStatement())
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	return err
}
