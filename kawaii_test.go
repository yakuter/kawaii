package kawaii

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "_db_dir_*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	dbPath := filepath.Join(tempDir, "testDB.db")

	db, err := New(dbPath)
	require.NoError(t, err)
	require.NotNil(t, db)
	defer db.Close()
}

func TestInitError(t *testing.T) {
	db, err := New("/wrong/path")
	require.Error(t, err)
	require.Nil(t, db)
	defer db.Close()
}

func TestSet(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "_db_dir_*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	dbPath := filepath.Join(tempDir, "testDB.db")

	db, err := New(dbPath)
	require.NoError(t, err)
	require.NotNil(t, db)
	defer db.Close()

	testBucket := "BucketA"
	testKey := "KeyA"
	testValue := "ValueA"

	err = db.Set(testBucket, testKey, testValue)
	require.NoError(t, err)

	actualValue, err := db.Get(testBucket, testKey)
	require.NoError(t, err)
	require.Equal(t, testValue, actualValue)
}

func TestGetAll(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "_db_dir_*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	dbPath := filepath.Join(tempDir, "agetnTestDB.db")

	db, err := New(dbPath)
	require.NoError(t, err)

	defer db.Close()

	testBucket := "TestBucket"
	testMap := map[string]string{
		"KeyA": "ValueA",
		"KeyB": "ValueB",
		"KeyC": "ValueC",
	}

	for k, v := range testMap {
		err = db.Set(testBucket, k, v)
		require.NoError(t, err)
	}

	valueMap, err := db.GetAll(testBucket)
	require.NoError(t, err)

	if diff := deep.Equal(testMap, valueMap); diff != nil {
		t.Error(diff)
	}

	for k, v := range valueMap {
		t.Logf("Key: %q Value: %q\n", k, v)
	}
}
