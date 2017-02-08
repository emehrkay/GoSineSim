# GoSineSim
Cosine Similarity of two or more shallow JSON objects in Go.

## Usage

```
gosignsim --source=$JSON_OBJ_LITERAL --pool=[$JSON_OBJ_LITERAL,...]
```

Let's say you had two simple maps `<String:Float>` and wanted to see the cosine similiary between the two

```javascript
{
    "cars": 30,
    "money": 99
}

[{
    "cars": 87,
    "money": 40
}]
````

You would pass them into the compiled app:

```
./gosinesim -source='{"cars": 30, "money": 99}' --pool='[{"cars": 87, "money": 40}]'
```

> the `source` argument is a single object literal while the `pool` is a collection of object literals

Which would produce the result

```
[{"Similarity":0.6632728204403626,"Data":{"cars":87,"money":40}}]
```