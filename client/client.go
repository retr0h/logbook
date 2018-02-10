// Copyright (c) 2018 John Dewey

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

// Package client provides primitives for accessing boltdb.
package client

import (
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/retr0h/logbook/entry"
)

var (
	db         *bolt.DB
	database   string
	timeout    time.Duration
	fileMode   os.FileMode
	bucketName []byte
)

// BoltPair is a struct to store key/value pair data.
type BoltPair struct {
	Key   []byte
	Value []byte
}

func init() {
	database = "my.db"
	timeout = 1 * time.Second
	fileMode = 0600
	bucketName = []byte("users")
}

// New initializes a bolt database, and creates the file if does not exist.
// By default it opens the file in 0600 mode, with a timeout period.
func New() error {
	var err error
	db, err = bolt.Open(database, fileMode, &bolt.Options{Timeout: timeout})
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucketName)

		return err
	})
	return err
}

// Close bolt db
func Close() {
	db.Close()
}

// List all key/value pairs from a bucket and returns a BoltPair struct.
func List() []BoltPair {
	var pairs []BoltPair

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		b.ForEach(func(k, v []byte) error {
			// Due to Byte slices returned from Bolt are only valid during a transaction.
			// Once the transaction has been committed or rolled back then the memory they
			// point to can be reused by a new page or can be unmapped from virtual memory
			// and you'll see an unexpected fault address panic when accessing it.  We copy
			// the slice to retain it
			dstk := make([]byte, len(k))
			dstv := make([]byte, len(v))
			copy(dstk, k)
			copy(dstv, v)

			pair := BoltPair{dstk, dstv}
			pairs = append(pairs, pair)

			return nil
		})

		return nil
	})

	return pairs
}

// Get value from bucket by key.
func Get(key []byte) ([]byte, error) {
	var value []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}
		v := b.Get(key)
		if v != nil {
			value = v
		}

		return nil
	})
	return value, err
}

// Put a key/value pair into target bucket.
func Put(key []byte, e *entry.Entry) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			return fmt.Errorf("bucket %q not found", bucketName)
		}

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		e.ID = int(id)

		encoded, err := e.Marshal()
		if err != nil {
			return err
		}

		err = b.Put(key, encoded)

		return err

	})
	return err
}
