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