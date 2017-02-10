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

I wrote a very simple benchmarking suite in python `bench.py` that simulates what it would be like to compare a random sampling of items from python's system call method.

The file creates 100,000 items which can contain between 5 and 15 entries. It then chunks the items into groups of 100, 1000, 5000... up to 100k. As you can see from the results below, a comparison of less than ~5k items can be done in realtime.

My specs are: macOS, 2.5 Core i5, 16GB 1600 DDR3, and an SSD

    ************************************************************

    Building the pool of 100000 items

    done building the pool: 3.90105509758


    ****************************************

    Comparing 100 items

    	run 1 took 0.0213799476624
    	run 2 took 0.0232019424438
    	run 3 took 0.0260858535767


    ****************************************

    Comparing 1000 items

    	run 1 took 0.104200839996
    	run 2 took 0.0777702331543
    	run 3 took 0.0913891792297


    ****************************************

    Comparing 5000 items

    	run 1 took 0.385035037994
    	run 2 took 0.361575126648
    	run 3 took 0.383772134781


    ****************************************

    Comparing 10000 items

    	run 1 took 0.728197813034
    	run 2 took 0.706686019897
    	run 3 took 0.723525047302


    ****************************************

    Comparing 25000 items

    	run 1 took 1.84598708153
    	run 2 took 1.81941795349
    	run 3 took 1.80622196198


    ****************************************

    Comparing 50000 items

    	run 1 took 3.73297214508
    	run 2 took 3.61929798126
    	run 3 took 3.69822692871


    ****************************************

    Comparing 75000 items

    	run 1 took 5.52517104149
    	run 2 took 5.44108700752
    	run 3 took 5.32282090187


    ****************************************

    Comparing 100000 items

    	run 1 took 7.28775191307
    	run 2 took 7.29696202278
    	run 3 took 7.31481003761


    ============================================================

    finished running: 63.7240190506