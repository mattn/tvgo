package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

func has(m []string, name string) bool {
	for _, s := range m {
		if s == name {
			return true
		}
	}
	return false
}

func main() {
	var v interface{}
	dec := json.NewDecoder(os.Stdin)
	dec.UseNumber()
	err := dec.Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	arr, ok := v.([]interface{})
	if !ok {
		log.Fatal("input stream should be array")
	}
	keys := []string{}
	for _, ai := range arr {
		elem, ok := ai.(map[string]interface{})
		if !ok {
			log.Fatal("input stream should be array of objects")
		}
		for k := range elem {
			if !has(keys, k) {
				keys = append(keys, k)
			}
		}
	}

	sort.Strings(keys)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(keys)
	table.SetAutoFormatHeaders(false)
	for _, ai := range arr {
		items := make([]string, len(keys))
		elem, _ := ai.(map[string]interface{})
		for i, k := range keys {
			if vv, ok := elem[k]; ok {
				if _, ok := vv.(map[string]interface{}); ok {
					items[i] = "..."
				} else if _, ok := vv.([]interface{}); ok {
					items[i] = "..."
				} else {
					items[i] = fmt.Sprint(vv)
				}
			}
		}
		table.Append(items)
	}
	table.Render()
}
