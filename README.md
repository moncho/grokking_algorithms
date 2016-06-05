## Grokking Algorithms in Go

Go implementation of algorithms explained in the [Grokking Algorithms](https://www.manning.com/books/grokking-algorithms) book.

I have tried (to the best of my abilities) to be faithful to the implementation of the algorithms as they are described in the book.

### Running tests

```
go test ./...
```

### Running benchmarks

```
go test -bench . ./...
```

### Debugging

[Godebug](https://github.com/mailgun/godebug) is used for debugging.

Install it, breakpoints are in the source code, but more can be inserted adding the following:

```
_ = "breakpoint"
```
And run the debugger:
```
godebug test -instrument github.com/moncho/grokking ./...
```
