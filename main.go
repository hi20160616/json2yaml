package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
)

var in = flag.String("f", "target.json", "The json file u need convert to yaml.")
var out = flag.String("o", "result.yml", "Where to save converted yaml file?")
var header = flag.String("H", "header.json", "Write `header.json` to yml first by default.\nOr, you can set \"\" for no header.")

func j2y(j, y string) error {
	jf, err := os.ReadFile(j)
	if err != nil {
		return err
	}
	// j := []byte(`{"name": "John", "age": 30}`)
	yf, err := yaml.JSONToYAML(jf)
	if err != nil {
		return err
	}
	out, err := os.OpenFile(y, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	out.Write(yf)
	return nil
}

func main() {
	flag.Parse()
	if *out == "result.yml" {
		*out = strings.TrimSuffix(*in, ".json") + ".yaml"
	}
	if *header != "" {
		if err := j2y("header.json", *out); err != nil {
			log.Fatalln(err)
		}
	}
	if err := j2y(*in, *out); err != nil {
		log.Fatalln(err)
	}
}
