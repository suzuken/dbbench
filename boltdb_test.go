package dbbench_test

import (
	"github.com/boltdb/bolt"
	"github.com/boltdb/coalescer"
	"github.com/suzuken/dummy"
	"log"
	"os"
	"testing"
	"time"
)

func BenchmarkBoltInsert(b *testing.B) {
	os.Remove("./bolt_bench.db")
	d := dummy.NewGenerator()

	db, err := bolt.Open("bolt_bench.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	// RW
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("TestBucket"))
		if err != nil {
			return err
		}
		for i := 0; i < b.N; i++ {
			err := bucket.Put([]byte(d.String(100)), []byte(d.Int(42)))
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkBoltBatchInsert(b *testing.B) {
	os.Remove("./bolt_bench.db")
	d := dummy.NewGenerator()

	db, err := bolt.Open("bolt_bench.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()
	// RW
	err = db.Batch(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("TestBucket"))
		if err != nil {
			return err
		}
		for i := 0; i < b.N; i++ {
			err := bucket.Put([]byte(d.String(100)), []byte(d.Int(42)))
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkBoltInsertManualTx(b *testing.B) {
	os.Remove("./bolt_bench.db")
	d := dummy.NewGenerator()

	db, err := bolt.Open("bolt_bench.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	b.ResetTimer()

	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte("TestBucket"))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		err := bucket.Put([]byte(d.String(100)), []byte(d.Int(42)))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func BenchmarkBoltCoalescerInsert(b *testing.B) {
	os.Remove("./bolt_bench.db")
	d := dummy.NewGenerator()

	db, err := bolt.Open("bolt_bench.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// FIXME heuristic parameters for limit and interval.
	//
	// coalescer behavior is:
	//
	// * flushing by interval ( here is by 1 nano sec.
	// * flushing when buffer reached limit. ( here is 1000 rec.
	//
	// but flusher seems not working immediately when reaching limit.
	c, err := coalescer.New(db, 1000, time.Nanosecond*1)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	err = c.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("TestBucket"))
		if err != nil {
			return err
		}
		return nil
	})
	// RW
	for i := 0; i < b.N; i++ {
		err = c.Update(func(tx *bolt.Tx) error {
			err = tx.Bucket([]byte("TestBucket")).Put([]byte(d.String(100)), []byte(d.Int(42)))
			if err != nil {
				log.Fatal(err)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
