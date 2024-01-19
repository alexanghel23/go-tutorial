## Function parameters
- Parameters are listed in paranthesis after function name
- Arguments is the data supplied in the function call

## Return values
- Functions can return a value as a result
- Type of return value after parameters in declaration
- Function call used on right-hand side of an assignment

```go
func foo(x int) int {
    return x + 1
}

y := foo(1)
```

## Multiple return values

```go
func foo2(x int) (int, int) {
    return x, x+1
}

a,b := foo2(3)
```