# Multiple Returns 

Go functions supports multiple return values. This is usefull in getting additional values such as error data. 

```go

result, err := fetchData()

///...

func divrem(i, j int) (int, int) {

	quotient := i / j
	modulus := i % j

	return quotient, modulus

}


```

## Named Returns

Go also supports named returns. This allows easier management of multiple return variables by declaring the name to be returned from the beginning. 


```go
func circle(radius float32) (diameter, circumf float32) {

	diameter = radius * 2
	circumf = 2 * math.Pi * radius

	return
}

```
The `return` keyword is required on the end of code execution

