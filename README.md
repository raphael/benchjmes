# go-jmespath Benchmark

This test suite benchmarks 
[go-jmespath](https://github.com/jmespath/go-jmespath) using the examples from
[http://jmespath.org/examples.html](http://jmespath.org/examples.html).

### Usage and results:
```
$ sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i7-2675QM CPU @ 2.20GHz

$ go test -bench=.
2016/03/14 14:03:44 Running 21 benchmarks
PASS
BenchmarkJMES0-8    100000       15563 ns/op
BenchmarkJMES1-8    100000       19176 ns/op
BenchmarkJMES2-8    100000       16795 ns/op
BenchmarkJMES3-8     50000       28577 ns/op
BenchmarkJMES4-8    100000       21822 ns/op
BenchmarkJMES5-8    100000       21883 ns/op
BenchmarkJMES6-8    100000       21757 ns/op
BenchmarkJMES7-8    100000       21927 ns/op
BenchmarkJMES8-8    100000       22646 ns/op
BenchmarkJMES9-8     50000       26533 ns/op
BenchmarkJMES10-8    50000       27894 ns/op
BenchmarkJMES11-8   100000       24312 ns/op
BenchmarkJMES12-8    50000       29632 ns/op
BenchmarkJMES13-8   100000       21307 ns/op
BenchmarkJMES14-8    50000       26598 ns/op
BenchmarkJMES15-8    50000       28448 ns/op
BenchmarkJMES16-8    50000       30182 ns/op
BenchmarkJMES17-8    50000       33460 ns/op
BenchmarkJMES18-8    50000       28851 ns/op
BenchmarkJMES19-8    50000       27373 ns/op
BenchmarkJMES20-8    50000       33420 ns/op
ok    github.com/raphael/benchjmes  42.156s
```

