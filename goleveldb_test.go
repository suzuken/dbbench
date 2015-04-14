package dbbench_test

import (
	"github.com/suzuken/dummy"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"testing"
)

func BenchmarkGoLevelDBInsert(b *testing.B) {
	os.Remove("./goleveldb_bench.db")
	d := dummy.NewGenerator()

	db, err := leveldb.OpenFile("./goleveldb_bench.db", nil)
	if err != nil {
		log.Fatalf("goleveldb initialization failed %s", err)
	}
	defer db.Close()

	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		db.Put([]byte(d.String(100)), []byte(d.String(100) + d.Int(13)), nil)
	}
}

func BenchmarkGoLevelDBBatchInsert(b *testing.B) {
	os.Remove("./goleveldb_bench.db")
	d := dummy.NewGenerator()

	db, err := leveldb.OpenFile("./goleveldb_bench.db", nil)
	if err != nil {
		log.Fatalf("goleveldb initialization failed %s", err)
	}
	defer db.Close()

	b.ResetTimer()
	batch := new(leveldb.Batch)
	for i:=0; i < b.N; i++ {
		batch.Put([]byte(d.String(100)), []byte(d.String(100) + d.Int(13)))
	}
	err = db.Write(batch, nil)
	if err != nil {
		log.Fatalf("batch write failed %s", err)
	}
}

