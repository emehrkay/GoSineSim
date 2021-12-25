package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	// "sort"
)

type Item struct {
	ID   string             `json:"id"`
	Data map[string]float64 `json:"data"`
}

type Items []Item

type Result struct {
	ID         string             `json:"id"`
	Similarity float64            `json:"similarity"`
	Data       map[string]float64 `json:"data"`
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
	newSource := source
	newSource.Data = map[string]float64{}

	for k, v := range source.Data {
		newSource.Data[k] = v
	}

	newOther := other
	newOther.Data = map[string]float64{}

	for k, v := range other.Data {
		newOther.Data[k] = v
	}

	for k := range newSource.Data {
		_, okay := newOther.Data[k]

		if okay == false {
			newOther.Data[k] = 0
		}
	}

	for k := range newOther.Data {
		_, okay := newSource.Data[k]

		if okay == false {
			newSource.Data[k] = 0
		}
	}

	return newSource, newOther
}

func getResultScore(source, other Item, resultChan chan Result) {
	source, other = pad(source, other)
	dem := norm(source) * norm(other)
	var score float64

	if dem > 0 {
		score = (dotProduct(source, other) / dem) * 100
	}
	// fmt.Printf("-----DEM: %v SCORE: %v\n\n", dem, score)
	resultChan <- Result{
		ID:         other.ID,
		Similarity: score,
		Data:       other.Data,
	}
}

func CoseineSimilarityWorker(source Item, pool Items, threshold float64) GoSignSimResults {
	results := make([]Result, 0)
	resChan := make(chan Result, len(pool))

	go func() {
		for _, item := range pool {
			sourceCopy := source
			itemCopy := item
			go getResultScore(sourceCopy, itemCopy, resChan)
		}
	}()

	for i := 0; i < len(pool); i++ {
		select {
		case res := <-resChan:
			if res.Similarity >= threshold {
				results = append(results, res)
			}
		}
	}

	return results
}

func main() {
	source := flag.String("source", "", "The source JSON object to compare")
	pool := flag.String("pool", "", "The data that will be compared against to source")
	pool_file := flag.String("pool_file", "", "An optional file to read the pool data from")
	threshold := flag.Float64("threshold", 0.0, "The lower limit ")
	output_file := flag.String("output_file", "", "The file to save the resulting JSON")
	verbose := flag.Bool("verbose", false, "Verbose mode")

	flag.Parse()

	var obj Item
	var results GoSignSimResults
	string_bytes := []byte(*source)
	err := json.Unmarshal(string_bytes, &obj)

	if err != nil {
		log.Fatal(err)
	}

	pf := *pool_file
	var pool_obj Items

	if pf != "" {
		pool_file_bytes, pool_file_err := ioutil.ReadFile(string(*pool_file))

		if pool_file_err != nil {
			fmt.Print(pool_file_err)
		}

		pool_err := json.Unmarshal(pool_file_bytes, &pool_obj)

		if pool_err != nil {
			log.Fatal(pool_err)
		}
	} else {
		pool_bytes := []byte(*pool)
		pool_err := json.Unmarshal(pool_bytes, &pool_obj)

		if pool_err != nil {
			log.Fatal(pool_err)
		}
	}

	results = CoseineSimilarityWorker(obj, pool_obj, float64(*threshold))
	results_json, results_err := json.Marshal(results)

	if results_err != nil {
		log.Fatal(results_err)
	}

	of := *output_file

	if of != "" {
		file, file_err := os.Create(of)

		if file_err != nil {
			log.Fatal("Could not create file: %s" + of)
			fmt.Printf("%s\n", results_json)
		}

		_, write_err := file.Write(results_json)

		if write_err != nil {
			log.Fatal("Unable to write to file: %s" + of)
			fmt.Printf("%s\n", results_json)
		}

		if *verbose {
			fmt.Printf("Done writing to file: %s\n\n", of)
			fmt.Printf("%s\n", results_json)
		}
	} else {
		fmt.Printf("%s\n", results_json)
	}
}
