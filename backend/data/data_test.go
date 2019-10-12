package data

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	testDatasourcePath = "./test_" + tableName + sqliteExtension
	testDatasource     *sqlite
	original           = sqliteDatabasePath
	testingContext     *testing.T
)

// test helper {{{

func createTmpCopy(sourceFile string) string {
	destinationFile := sourceFile + ".tmp"
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		testingContext.Fatal("Error reading", sourceFile)
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		testingContext.Fatal("Error creating", destinationFile)
	}
	return destinationFile
}

func setup(t *testing.T, sqlSourcePath string) {
	testingContext = t
	if len(sqlSourcePath) == 0 {
		sqliteDatabasePath = testDatasourcePath
	} else {
		sqliteDatabasePath = createTmpCopy(sqlSourcePath)
	}
	tmp, err := NewSqlite()
	testDatasource = tmp.(*sqlite)
	if err != nil {
		t.Error(err)
	}
}

func teardown() {
	if sqliteDatabasePath != original {
		os.Remove(sqliteDatabasePath)
	}
}

// }}}

// actual tests {{{

func TestNewSqlite_ShouldCreate_SqliteFile(t *testing.T) {
	setup(t, "")
	defer teardown()

	if _, err := os.Stat(sqliteDatabasePath); err != nil {
		t.Error("Expected sqlite3 file created, but couldn't read stat to it.")
	}
}

func TestNewSqlite_ShouldCreate_TriggerInSqliteDb(t *testing.T) {
	setup(t, "")
	defer teardown()

	created, err := testDatasource.checkTriggerExists(triggerNameUpdateDate)
	if err != nil {
		t.Error(err)
	}
	if !created {
		t.Error("According to <checkTriggerExists>, trigger was not created.")
	}
}

func TestInsertArticle_ShouldIncrement_Datarows(t *testing.T) {
	setup(t, "")
	defer teardown()

	a := NewArticleWithArray([]string{
		"1", "George", "Return of the jedis", "In a galaxy far far away ...", "1", "", "1"})
	// t.Log(p_InsertArticle(a))

	testDatasource.InsertArticle(a)
}

func TestGetArticleWhere_ShouldReturn_ArrayLengthThree(t *testing.T) {

}

// func TestUdateArticle_ShouldUpdate_DateThroughTrigger(t *testing.T) {
// 	panic("not imlemented")
// }

// }}}

// vim: fdn=1 fdm=marker:
