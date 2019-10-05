package datasource

import (
	"os"
	"testing"
)

var (
	testDatabase = sqlite{path: "./test_" + dataRootName + sqliteExtension}
)

func CleanupTest() {
	os.Remove(testDatabase.path)
}

func TestCreateDatabaseShouldCreateFile(t *testing.T) {
	defer CleanupTest()
	err := testDatabase.Setup()
	t.Log(sqliteArticleTableDefinition.getSqliteTableCreationStatement())
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(testDatabase.path); err != nil {
		t.Error(err)
	}
}
