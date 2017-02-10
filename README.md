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

    done building the pool: 3.93292379379


    ****************************************

    Comparing 100 items

    	run 1 took 0.0157909393311
    	run 2 took 0.0160090923309
    	run 3 took 0.0309271812439


    ****************************************

    Comparing 1000 items

    	run 1 took 0.0655910968781
    	run 2 took 0.0748629570007
    	run 3 took 0.0511500835419


    ****************************************

    Comparing 5000 items

    	run 1 took 0.22353386879
    	run 2 took 0.204764127731
    	run 3 took 0.197242021561


    ****************************************

    Comparing 10000 items

    	run 1 took 0.291663885117
    	run 2 took 0.608170032501
    	run 3 took 0.627039909363


    ****************************************

    Comparing 25000 items

    	run 1 took 0.51177406311
    	run 2 took 0.534518003464
    	run 3 took 0.545824050903


    ****************************************

    Comparing 50000 items

    	run 1 took 0.905314922333
    	run 2 took 0.918902873993
    	run 3 took 0.902179002762


    ****************************************

    Comparing 75000 items

    	run 1 took 1.37646508217
    	run 2 took 1.3969039917
    	run 3 took 1.66373586655


    ****************************************

    Comparing 100000 items

    	run 1 took 1.83165812492
    	run 2 took 1.93080997467
    	run 3 took 5.73800802231


    ============================================================

    finished running: 26.3283598423