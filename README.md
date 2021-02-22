# kawaii

### Installing

To start using Kawaii, install Go and run `go get`:

```sh
$ go get github.com/yakuter/kawaii
```
### Usage

To use bbolt as an embedded key-value store, import as:

```go
import (
	"log"

	"github.com/yakuter/kawaii"
)

func main() {
	dbPath := "./mydb"
	db, err := kawaii.New(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	testBucket := "BucketA"
	testKey := "KeyA"
	testValue := "ValueA"

	err = db.Set(testBucket, testKey, testValue)
	if err != nil {
		log.Println(err)
	}

	value, err := db.Get(testBucket, testKey)
	if err != nil {
		log.Println(err)
	}

	log.Println(value)
}
```