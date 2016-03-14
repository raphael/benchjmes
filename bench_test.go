package main

import (
	"encoding/json"
	"log"
	"testing"
)

var benchmarks = []struct {
	Data string
	Expr string
}{
	{`{"a": "foo", "b": "bar", "c": "baz"}`, `a`},
	{`{"a": {"b": {"c": {"d": "value"}}}}`, `a.b.c.d`},
	{`["a", "b", "c", "d", "e", "f"]`, `[1]`},
	{`{"a": {
		"b": {
			"c": [
			{"d": [0, [1, 2]]},
			{"d": [3, 4]}
			]
		}
	}}`, `a.b.c[0].d[1][0]`},
	{`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`, `[0:5]`},
	{`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`, `[5:10]`},
	{`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`, `[:5]`},
	{`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`, `[::2]`},
	{`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`, `[::-1]`},
	{`{
		"people": [
			{"first": "James", "last": "d"},
			{"first": "Jacob", "last": "e"},
			{"first": "Jayden", "last": "f"},
			{"missing": "different"}
		],
		"foo": {"bar": "baz"}
	}`, `people[*].first`},
	{`{
		"people": [
			{"first": "James", "last": "d"},
			{"first": "Jacob", "last": "e"},
			{"first": "Jayden", "last": "f"},
			{"missing": "different"}
		],
		"foo": {"bar": "baz"}
	}`, `people[:2].first`},
	{`{
		"ops": {
			"functionA": {"numArgs": 2},
			"functionB": {"numArgs": 3},
			"functionC": {"variadic": true}
		}
	}`, `ops.*.numArgs`},
	{`{
		"reservations": [
		{
			"instances": [
			{"state": "running"},
			{"state": "stopped"}
			]
		},
		{
			"instances": [
			{"state": "terminated"},
			{"state": "runnning"}
			]
		}
		]
	}`, `reservations[*].instances[*].state`},
	{`[
		[0, 1],
		2,
		[3],
		4,
		[5, [6, 7]]
	]`, `[]`},
	{`{
		"machines": [
		{"name": "a", "state": "running"},
		{"name": "b", "state": "stopped"},
		{"name": "b", "state": "running"}
		]
	}`, `machines[?state=='running'].name`},
	{`{
		"people": [
		{"first": "James", "last": "d"},
		{"first": "Jacob", "last": "e"},
		{"first": "Jayden", "last": "f"},
		{"missing": "different"}
		],
		"foo": {"bar": "baz"}
	}`, `people[*].first | [0]`},
	{`{
		"people": [
		{
			"name": "a",
			"state": {"name": "up"}
		},
		{
			"name": "b",
			"state": {"name": "down"}
		},
		{
			"name": "c",
			"state": {"name": "up"}
		}
		]
	}`, `people[].[name, state.name]`},
	{`{
		"people": [
		{
			"name": "a",
			"state": {"name": "up"}
		},
		{
			"name": "b",
			"state": {"name": "down"}
		},
		{
			"name": "c",
			"state": {"name": "up"}
		}
		]
	}`, `people[].{Name: name, State: state.name}`},
	{`{
		"people": [
		{
			"name": "b",
			"age": 30,
			"state": {"name": "up"}
		},
		{
			"name": "a",
			"age": 50,
			"state": {"name": "down"}
		},
		{
			"name": "c",
			"age": 40,
			"state": {"name": "up"}
		}
		]
	}`, `length(people)`},
	{`{
		"people": [
		{
			"name": "b",
			"age": 30
		},
		{
			"name": "a",
			"age": 50
		},
		{
			"name": "c",
			"age": 40
		}
		]
	}`, `max_by(people, &age).name`},
	{`{
		"myarray": [
		"foo",
		"foobar",
		"barfoo",
		"bar",
		"baz",
		"barbaz",
		"barfoobaz"
		]
	}`, `myarray[?contains(@, 'foo') == ` + "`" + `true` + "`" + `]`},
}

func benchmarkBaseline(expr, data string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var v interface{}
		err := json.Unmarshal([]byte(data), &v)
		if err != nil {
			panic(err)
		}
		if v == nil {
			log.Fatalf("%s on %v returned nil", expr, data)
		}
	}
}

func benchmarkJMES(expr, data string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		v, err := Apply(expr, []byte(data))
		if err != nil {
			panic(err)
		}
		if v == nil {
			log.Fatalf("%s on %v returned nil", expr, data)
		}
	}
}

func TestPrintBenchmarkTitle(t *testing.T) { log.Printf("Running %d benchmarks", len(benchmarks)) }

func BenchmarkBaseline0(b *testing.B) { benchmarkBaseline(benchmarks[0].Expr, benchmarks[0].Data, b) }
func BenchmarkJMES0(b *testing.B)     { benchmarkJMES(benchmarks[0].Expr, benchmarks[0].Data, b) }

func BenchmarkBaseline1(b *testing.B) { benchmarkBaseline(benchmarks[1].Expr, benchmarks[1].Data, b) }
func BenchmarkJMES1(b *testing.B)     { benchmarkJMES(benchmarks[1].Expr, benchmarks[1].Data, b) }

func BenchmarkBaseline2(b *testing.B) { benchmarkBaseline(benchmarks[2].Expr, benchmarks[2].Data, b) }
func BenchmarkJMES2(b *testing.B)     { benchmarkJMES(benchmarks[2].Expr, benchmarks[2].Data, b) }

func BenchmarkBaseline3(b *testing.B) { benchmarkBaseline(benchmarks[3].Expr, benchmarks[3].Data, b) }
func BenchmarkJMES3(b *testing.B)     { benchmarkJMES(benchmarks[3].Expr, benchmarks[3].Data, b) }

func BenchmarkBaseline4(b *testing.B) { benchmarkBaseline(benchmarks[4].Expr, benchmarks[4].Data, b) }
func BenchmarkJMES4(b *testing.B)     { benchmarkJMES(benchmarks[4].Expr, benchmarks[4].Data, b) }

func BenchmarkBaseline5(b *testing.B) { benchmarkBaseline(benchmarks[5].Expr, benchmarks[5].Data, b) }
func BenchmarkJMES5(b *testing.B)     { benchmarkJMES(benchmarks[5].Expr, benchmarks[5].Data, b) }

func BenchmarkBaseline6(b *testing.B) { benchmarkBaseline(benchmarks[6].Expr, benchmarks[6].Data, b) }
func BenchmarkJMES6(b *testing.B)     { benchmarkJMES(benchmarks[6].Expr, benchmarks[6].Data, b) }

func BenchmarkBaseline7(b *testing.B) { benchmarkBaseline(benchmarks[7].Expr, benchmarks[7].Data, b) }
func BenchmarkJMES7(b *testing.B)     { benchmarkJMES(benchmarks[7].Expr, benchmarks[7].Data, b) }

func BenchmarkBaseline8(b *testing.B) { benchmarkBaseline(benchmarks[8].Expr, benchmarks[8].Data, b) }
func BenchmarkJMES8(b *testing.B)     { benchmarkJMES(benchmarks[8].Expr, benchmarks[8].Data, b) }

func BenchmarkBaseline9(b *testing.B) { benchmarkBaseline(benchmarks[9].Expr, benchmarks[9].Data, b) }
func BenchmarkJMES9(b *testing.B)     { benchmarkJMES(benchmarks[9].Expr, benchmarks[9].Data, b) }

func BenchmarkBaseline10(b *testing.B) { benchmarkBaseline(benchmarks[10].Expr, benchmarks[10].Data, b) }
func BenchmarkJMES10(b *testing.B)     { benchmarkJMES(benchmarks[10].Expr, benchmarks[10].Data, b) }

func BenchmarkBaseline11(b *testing.B) { benchmarkBaseline(benchmarks[11].Expr, benchmarks[11].Data, b) }
func BenchmarkJMES11(b *testing.B)     { benchmarkJMES(benchmarks[11].Expr, benchmarks[11].Data, b) }

func BenchmarkBaseline12(b *testing.B) { benchmarkBaseline(benchmarks[12].Expr, benchmarks[12].Data, b) }
func BenchmarkJMES12(b *testing.B)     { benchmarkJMES(benchmarks[12].Expr, benchmarks[12].Data, b) }

func BenchmarkBaseline13(b *testing.B) { benchmarkBaseline(benchmarks[13].Expr, benchmarks[13].Data, b) }
func BenchmarkJMES13(b *testing.B)     { benchmarkJMES(benchmarks[13].Expr, benchmarks[13].Data, b) }

func BenchmarkBaseline14(b *testing.B) { benchmarkBaseline(benchmarks[14].Expr, benchmarks[14].Data, b) }
func BenchmarkJMES14(b *testing.B)     { benchmarkJMES(benchmarks[14].Expr, benchmarks[14].Data, b) }

func BenchmarkBaseline15(b *testing.B) { benchmarkBaseline(benchmarks[15].Expr, benchmarks[15].Data, b) }
func BenchmarkJMES15(b *testing.B)     { benchmarkJMES(benchmarks[15].Expr, benchmarks[15].Data, b) }

func BenchmarkBaseline16(b *testing.B) { benchmarkBaseline(benchmarks[16].Expr, benchmarks[16].Data, b) }
func BenchmarkJMES16(b *testing.B)     { benchmarkJMES(benchmarks[16].Expr, benchmarks[16].Data, b) }

func BenchmarkBaseline17(b *testing.B) { benchmarkBaseline(benchmarks[17].Expr, benchmarks[17].Data, b) }
func BenchmarkJMES17(b *testing.B)     { benchmarkJMES(benchmarks[17].Expr, benchmarks[17].Data, b) }

func BenchmarkBaseline18(b *testing.B) { benchmarkBaseline(benchmarks[18].Expr, benchmarks[18].Data, b) }
func BenchmarkJMES18(b *testing.B)     { benchmarkJMES(benchmarks[18].Expr, benchmarks[18].Data, b) }

func BenchmarkBaseline19(b *testing.B) { benchmarkBaseline(benchmarks[19].Expr, benchmarks[19].Data, b) }
func BenchmarkJMES19(b *testing.B)     { benchmarkJMES(benchmarks[19].Expr, benchmarks[19].Data, b) }

func BenchmarkBaseline20(b *testing.B) { benchmarkBaseline(benchmarks[20].Expr, benchmarks[20].Data, b) }
func BenchmarkJMES20(b *testing.B)     { benchmarkJMES(benchmarks[20].Expr, benchmarks[20].Data, b) }
