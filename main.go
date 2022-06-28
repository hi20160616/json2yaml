package main

import (
	"flag"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

var fFlag = flag.String("f", "foo.json", "The json file u need convert to yaml.")
var outFlag = flag.String("o", "bar.yml", "Where to save converted yaml file?")

func main() {
	flag.Parse()
	j, err := os.ReadFile(*fFlag)
	if err != nil {
		log.Fatalf("Open json file error: %v", err)
	}
	// j := []byte(`{"name": "John", "age": 30}`)
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		log.Fatalf("JSONToYAML error: %v", err)
		return
	}
	if *outFlag == "" {
		*outFlag = *fFlag + ".yml"
	}
	out, err := os.OpenFile(*outFlag, os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Save yaml error: %v", err)
	}
	out.Write(y)
	// fmt.Println(string(y))
	/* Output:
	name: John
	age: 30
	*/
}
