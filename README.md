# kawaii

### Installing

To start using Kawaii, install Go and run `go get`:

```sh
$ go get github.com/yakuter/kawaii
```
### Usage

To use bbolt as an embedded key-value store, import as:

```go
import bolt "github.com/yakuter/kawaii"

dbPath := "./mydb"
db, err := New(dbPath)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

err = db.Set(testBucket, testKey, testValue)
if err != nil {
    log.Println(err)
}

value, err := db.Get(testBucket, testKey)
if err != nil {
    log.Println(err)
}

log.Println(value)
```