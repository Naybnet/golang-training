# Golang package

Go programs are made of packages. 

Compile and run with

```
go run main.go
```

## Testing

The package testing is used for unit tests.
Convention for test filenames `{filename}_test.go` and the test functions are called `Test{Fn}`.

Tests are executed with `go test controller/*`.

An interesting talk on testing : [Mitchell Hashimoto - Advanced Testing with Go](https://youtu.be/8hQG7QlcLBk).