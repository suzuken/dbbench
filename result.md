# benchmark

spec

* MacBook Pro, Late 2013
* 2.6 GHz Intel Core i7
* 16GB 1600 MHz DDR3
* SSD
* sqlite3: 3.7.13
* bolt: based on https://github.com/boltdb/bolt/commit/3b449559cf34cbcc74460b59041a4399d3226e5a

```
	-> % sw_vers
	ProductName:    Mac OS X
	ProductVersion: 10.9.5
	BuildVersion:   13F34

	-> % make bench
	go test -bench .
	testing: warning: no tests to run
	PASS
	BenchmarkBoltInsert       100000             85617 ns/op
	BenchmarkBoltBatchInsert          100000             84179 ns/op
	BenchmarkBoltInsertManualTx       100000             84444 ns/op
	BenchmarkBoltCoalescerInsert        3000            484263 ns/op
	BenchmarkGoLevelDBInsert          100000             29371 ns/op
	BenchmarkGoLevelDBBatchInsert      50000             24801 ns/op
	BenchmarkSQLiteReplaceInsertWithTx         50000             32904 ns/op
	BenchmarkSQLiteReplaceInsert        3000            565700 ns/op
	BenchmarkSQLiteInsertWithTx       100000             27211 ns/op
	BenchmarkSQLiteInsert       3000            550811 ns/op
	BenchmarkSQLiteReadPrimaryKey      50000             22743 ns/op
	BenchmarkSQLiteReadKeyWithoutIndex           300           4998885 ns/op
	BenchmarkSQLiteReadKeyWithIndex   100000             24240 ns/op
	BenchmarkSQLiteReadRangeWithoutIndex       50000             27262 ns/op
	BenchmarkSQLiteReadRangeWithIndex          50000             33665 ns/op
	ok      github.com/suzuken/dbbench      164.240s
```
