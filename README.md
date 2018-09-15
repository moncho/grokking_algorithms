# Grokking Algorithms in Go

Go implementation of algorithms explained in the [Grokking Algorithms](https://www.manning.com/books/grokking-algorithms) book.

I have tried (to the best of my abilities) to be faithful to the implementation of the algorithms as they are described in the book.

## Running tests

```
go test ./...
```

## Running benchmarks

```bash
go test -benchmem -bench=. ./...
```