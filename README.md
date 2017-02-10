# GoSineSim
Cosine Similarity of two or more shallow JSON objects in Go.

## Usage

```
gosignsim --source=$JSON_OBJ_LITERAL --pool=[$JSON_OBJ_LITERAL,...] [--threshold=float]
```

GoSineSim works by comparing one data structure against one or more. Each item should be made up like this:

```javascript
{
    "id": "15",
    "data": {
        "cars": 30,
        "money": 99
    }
}

[{
    "id": "44",
    "data": {
        "cars": 87,
        "money": 40
    }
}]
````

### Data Rules

1. Each item set must have a string `id` field and an object literal `data` field. The `data` is a simple key: value pair where the keys are stings and the values are floats
    
    ```go
    type Item struct {
    	Id   string
    	Data map[string]float64
    }
    ```

2. The `source` argument is a single item
3. The `pool` argument is a list of items

### Example

```
./gosinesim -source='{"id": "15", "data": {"cars": 30, "money": 99}}' --pool='[{"id": "44", "data": {"cars": 87, "money": 40}}]'
```

Which would produce the result

```
[{"Similarity":66.32728204403626,"Id":"44","Data":{"cars":87,"money":40}}]
```

### Benchmarks

I wrote a very simple benchmarking suite in python `bench.py` that simulates what it would be like to compare a random sampling of items from another language's system calling method.

The file creates 100,000 items which can contain between 5 and 15 entries. It then chunks the items into groups of 100, 1000, 5000... up to 100k. As you can see a comparison of less than ~5k items can be done in realtime.

My specs are: macOS, 2.5 Core i5, 16GB 1600 DDR3, and an SSD

    ************************************************************

    Building the pool of 100000 items

    done building the pool: 4.08293294907


    ****************************************

    Running 100 tags

    	run 1 took 0.0148870944977
    	run 2 took 0.01486992836
    	run 3 took 0.0157899856567


    ****************************************

    Running 1000 tags

    	run 1 took 0.0886211395264
    	run 2 took 0.0958361625671
    	run 3 took 0.0808701515198


    ****************************************

    Running 5000 tags

    	run 1 took 0.397789955139
    	run 2 took 0.366457939148
    	run 3 took 0.390224933624


    ****************************************

    Running 10000 tags

    	run 1 took 0.731570959091
    	run 2 took 0.7547519207
    	run 3 took 0.756834030151


    ****************************************

    Running 25000 tags

    	run 1 took 1.84421396255
    	run 2 took 1.81892490387
    	run 3 took 1.84610414505


    ****************************************

    Running 50000 tags

    	run 1 took 3.67140698433
    	run 2 took 3.70125007629
    	run 3 took 3.64767408371


    ****************************************

    Running 75000 tags

    	run 1 took 5.64399909973
    	run 2 took 5.55435490608
    	run 3 took 5.80814218521


    ****************************************

    Running 100000 tags

    	run 1 took 7.6290500164
    	run 2 took 7.41003990173
    	run 3 took 7.53508520126


    ============================================================

    finished running: 65.6274158955