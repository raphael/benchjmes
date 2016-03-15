# go-jmespath Benchmark

This test suite benchmarks 
[go-jmespath](https://github.com/jmespath/go-jmespath) using the examples from
[http://jmespath.org/examples.html](http://jmespath.org/examples.html).

### Usage and results:
```
$ sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i7-2675QM CPU @ 2.20GHz

$ go test -bench=.
2016/03/14 16:39:10 Running 21 benchmarks
PASS
BenchmarkBaseline0-8    500000        2636 ns/op
BenchmarkJMES0-8        100000       15477 ns/op
BenchmarkBaseline1-8    300000        3949 ns/op
BenchmarkJMES1-8        100000       19344 ns/op
BenchmarkBaseline2-8    500000        3090 ns/op
BenchmarkJMES2-8        100000       17193 ns/op
BenchmarkBaseline3-8    200000        8405 ns/op
BenchmarkJMES3-8         50000       28547 ns/op
BenchmarkBaseline4-8    300000        4532 ns/op
BenchmarkJMES4-8        100000       22110 ns/op
BenchmarkBaseline5-8    300000        4603 ns/op
BenchmarkJMES5-8        100000       21731 ns/op
BenchmarkBaseline6-8    300000        4441 ns/op
BenchmarkJMES6-8        100000       21583 ns/op
BenchmarkBaseline7-8    300000        4466 ns/op
BenchmarkJMES7-8        100000       21891 ns/op
BenchmarkBaseline8-8    300000        4468 ns/op
BenchmarkJMES8-8        100000       22417 ns/op
BenchmarkBaseline9-8    200000        9900 ns/op
BenchmarkJMES9-8         50000       26617 ns/op
BenchmarkBaseline10-8   200000       10039 ns/op
BenchmarkJMES10-8        50000       26763 ns/op
BenchmarkBaseline11-8   200000        6691 ns/op
BenchmarkJMES11-8       100000       23683 ns/op
BenchmarkBaseline12-8   200000       10508 ns/op
BenchmarkJMES12-8        50000       28987 ns/op
BenchmarkBaseline13-8   300000        5968 ns/op
BenchmarkJMES13-8       100000       21385 ns/op
BenchmarkBaseline14-8   200000        7717 ns/op
BenchmarkJMES14-8        50000       26349 ns/op
BenchmarkBaseline15-8   200000        9999 ns/op
BenchmarkJMES15-8        50000       28395 ns/op
BenchmarkBaseline16-8   200000       10078 ns/op
BenchmarkJMES16-8        50000       29548 ns/op
BenchmarkBaseline17-8   200000       10094 ns/op
BenchmarkJMES17-8        50000       31823 ns/op
BenchmarkBaseline18-8   200000       12035 ns/op
BenchmarkJMES18-8        50000       27628 ns/op
BenchmarkBaseline19-8   200000        7866 ns/op
BenchmarkJMES19-8        50000       25852 ns/op
BenchmarkBaseline20-8   300000        5268 ns/op
BenchmarkJMES20-8        50000       32575 ns/op
ok    github.com/raphael/benchjmes  78.074s
```

