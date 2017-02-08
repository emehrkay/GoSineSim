package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"sort"
)

type Result struct {
	Similarity float64
	Data       map[string]float64
}

type GoSignSimResults []Result

func (slice GoSignSimResults) Len() int {
	return len(slice)
}

func (slice GoSignSimResults) Less(i, j int) bool {
	return slice[i].Similarity < slice[j].Similarity
}

func (slice GoSignSimResults) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func norm(obj map[string]float64) float64 {
	var norm float64 = 0

	for _, v := range obj {
		norm += v * v
	}

	return math.Sqrt(norm)
}

func dotProduct(source map[string]float64, other map[string]float64) float64 {
	var product float64 = 0

	for k, v := range source {
		product += v * other[k]
	}

	return product
}

func pad(source map[string]float64, other map[string]float64) (map[string]float64, map[string]float64) {
	for k, _ := range source {
		_, okay := other[k]

		if okay == false {
			other[k] = 0
		}
	}

	for k, _ := range other {
		_, okay := source[k]

		if okay == false {
			source[k] = 0
		}
	}

	return source, other
}

func getScore(source map[string]float64, other map[string]float64) float64 {
	source, other = pad(source, other)
	dem := norm(source) * norm(other)

	if dem > 0 {
		return dotProduct(source, other) / dem
	}

	return 0
}

func CoseineSimilarity(source map[string]float64, pool []map[string]float64) GoSignSimResults {
	var results GoSignSimResults

	for _, other := range pool {
		score := getScore(source, other)
		res := Result{Similarity: score, Data: other}
		results = append(results, res)
	}

	sort.Sort(results)

	return results
}

func main() {
	source := flag.String("source", "", "The source JSON object to compare")
	pool := flag.String("pool", "", "The data that will be compared against to source")

	flag.Parse()

	var obj map[string]float64
	string_bytes := []byte(*source)
	err := json.Unmarshal(string_bytes, &obj)

	if err != nil {
		log.Fatal(err)
	}

	var pool_obj []map[string]float64
	pool_bytes := []byte(*pool)
	pool_err := json.Unmarshal(pool_bytes, &pool_obj)

	if pool_err != nil {
		log.Fatal(pool_err)
	}

	results := CoseineSimilarity(obj, pool_obj)
	results_json, _ := json.Marshal(results)

	fmt.Printf("Results: %s \n", results_json)
}
