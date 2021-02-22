package kawaii

import (
	bolt "go.etcd.io/bbolt"
)

// DB is main database struct
type DB struct {
	db *bolt.DB
}

// New initializes database
func New(dbFilePath string) (*DB, error) {

	boltDB, err := bolt.Open(dbFilePath, 0666, nil)
	if err != nil {
		return nil, err
	}

	db := &DB{db: boltDB}

	return db, nil
}

// Close closes opened database
func (db *DB) Close() error {
	if db != nil && db.db != nil {
		return db.db.Close()
	}
	return nil
}

// Set adds new key value to databas
func (db *DB) Set(bucket, key, value string) error {

	// Start a write transaction.
	err := db.db.Update(func(tx *bolt.Tx) error {

		// Create a bucket.
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		// Set the key for the value.
		if err := b.Put([]byte(key), []byte(value)); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// Get retrieves value from database
func (db *DB) Get(bucket, key string) (string, error) {

	var dbValue []byte

	// Retrieve the key again.
	err := db.db.View(func(tx *bolt.Tx) error {
		dbValue = tx.Bucket([]byte(bucket)).Get([]byte(key))
		return nil
	})

	if err != nil {
		return "", err
	}

	return string(dbValue), nil

}

// GetAll retrieves all keys and values in a bucket from database
func (db *DB) GetAll(bucket string) (map[string]string, error) {

	bucketMap := make(map[string]string)

	// Retrieve the key again.
	err := db.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.ForEach(func(k, v []byte) error {
			bucketMap[string(k)] = string(v)
			return nil
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return bucketMap, nil
}
