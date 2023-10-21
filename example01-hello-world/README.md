# Hello World

## gofmt ✓

Try to fix `main.go` file using `gofmt` command and show diff data from `gofmt -d example01` command.

```diff
        fmt.Printf(HelloWorld("appleboy"))
        fmt.Println("一天就學會 Go 語言")

-       if (a >= 1) { fmt.Println("a >= 1") }
+       if a >= 1 {
+               fmt.Println("a >= 1")
+       }
```

Fix and save automatically using `-w` flag: `gofmt -w example01`

## golint

Golint is deprecated and should be replaced by [Staticcheck](https://staticcheck.io/).

```
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck -checks all
```

The explanations are available at 

```
staticcheck -explain SA1006
```

Go vet should also be used: `go vet .`

```
main.go:7:2: printf-style function with dynamic format string and no further arguments should use print-style function instead (SA1006)
main.go:15:17: should not use underscores in Go names; func parameter user_name should be userName (ST1003)
```
