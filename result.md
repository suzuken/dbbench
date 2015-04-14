# 2015/04/14

	-> % make bench
	go test -bench .
	testing: warning: no tests to run
	PASS
	BenchmarkBoltInsert       100000             90119 ns/op
	BenchmarkBoltBatchInsert          100000             88330 ns/op
	BenchmarkBoltInsertManualTx       100000             97870 ns/op
	BenchmarkBoltCoalescerInsert        3000            507014 ns/op
	BenchmarkSQLiteReplaceInsertWithTx         50000             38388 ns/op
	BenchmarkSQLiteReplaceInsert        2000            764693 ns/op
	BenchmarkSQLiteInsertWithTx       100000             27249 ns/op
	BenchmarkSQLiteInsert       2000            801088 ns/op
	ok      github.com/suzuken/dbbench      38.410s
