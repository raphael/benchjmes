package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	jmespath "github.com/jmespath/go-jmespath"
)

func main() {
	var (
		input = flag.String("file", "data.json", "file containing JSON")
		expr  = flag.String("expr", "", "JMESPath expression")
	)
	flag.Parse()
	if input == nil {
		log.Fatal("--file flag missing")
	}
	if expr == nil {
		log.Fatal("--expr flag missing")
	}
	data, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}
	out, err := Apply(*expr, data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", out)
}

func Apply(expr string, data []byte) (interface{}, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}
	res, err := jmespath.Search(expr, v)
	return res, err
}
