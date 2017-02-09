package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"sort"
)

type Item struct {
	Id   string
	Data map[string]float64
}

type Items []Item

type Result struct {
	Similarity float64
	Id         string
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

func norm(obj Item) float64 {
	var norm float64 = 0

	for _, v := range obj.Data {
		norm += v * v
	}

	return math.Sqrt(norm)
}

func dotProduct(source, other Item) float64 {
	var product float64 = 0

	for k, v := range source.Data {
		product += v * other.Data[k]
	}

	return product
}

func pad(source, other Item) (Item, Item) {
	for k, _ := range source.Data {
		_, okay := other.Data[k]

		if okay == false {
			other.Data[k] = 0
		}
	}

	for k, _ := range other.Data {
		_, okay := source.Data[k]

		if okay == false {
			source.Data[k] = 0
		}
	}

	return source, other
}

func getScore(source, other Item) float64 {
	source, other = pad(source, other)
	dem := norm(source) * norm(other)

	if dem > 0 {
		return dotProduct(source, other) / dem
	}

	return 0
}

func CoseineSimilarity(source Item, pool Items) GoSignSimResults {
	var results GoSignSimResults

	for _, other := range pool {
		score := getScore(source, other)
		res := Result{Similarity: score, Data: other.Data, Id: other.Id}
		results = append(results, res)
	}

	sort.Sort(results)

	return results
}

func main() {
	source := flag.String("source", "", "The source JSON object to compare")
	pool := flag.String("pool", "", "The data that will be compared against to source")

	flag.Parse()

	var obj Item
	string_bytes := []byte(*source)
	err := json.Unmarshal(string_bytes, &obj)

	if err != nil {
		log.Fatal(err)
	}

	var pool_obj Items
	pool_bytes := []byte(*pool)
	pool_err := json.Unmarshal(pool_bytes, &pool_obj)

	if pool_err != nil {
		log.Fatal(pool_err)
	}

	results := CoseineSimilarity(obj, pool_obj)
	results_json, _ := json.Marshal(results)

	fmt.Printf("%s", results_json)
}
