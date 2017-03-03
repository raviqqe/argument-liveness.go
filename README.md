# argument-liveness.go

Understanding argument liveness in Go 1.8


## Running examples

Each example directory has `main.go` which iterates over an infinite list.

```
go run $dir/main.go
```


## Results

Does the program cause memory leak?

dir | go1.7 | go1.8
---------|-------|------
struct | No | No
struct\_pointer | No | No


## License

[The Unlicense](https://unlicense.org)
