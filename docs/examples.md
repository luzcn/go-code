## Go Examples

### create a slice of slice 
```go
// similar to java List<List<Integer>> x = new ArrayList<>()
x := make([][]int, 0)

x = append(x, []int{1, 24, 4})
x = append(x, []int{1, 2})
```

### `for` loop
```go
// while loop
for i < 10 {
    i++
}

// for with condition
for i:=1; i < 10; i++ {
    fmt.Println(i)
}

// for with range
nums := []int{1,2,3,4,12}
for i,n := range nums {
    fmt.Println(n)
}
```

### use hashmap
```go
// declare a hashmap object
var mp map[string]int

// define a hashmap
mp := make(map[string]int)

// check if an element in the map
v,found := mp["key"]

// directly use in if condition
if _,found := mp["key"]; found {
    // do something
}

```

#### string to int, int to string
```go
// string to integer
s, _:= strconv.Atoi("123")
print(s)


// integer to string
s := strconv.Itoa(123)
print(s)

// float64 to string
s := strconv.FormatFloat(123.12301, 'f', 3, 64)
print(s)

// string to float64
f, err := strconv.ParseFloat("3.1415", 64)
print(f)

// convert binary to int64
i, err := strconv.ParseInt("-42", 10, 64)
print(i)
```

#### Returns multiple values
```go
package example
import (
    "math"
    "strconv"
)

func CircleInfo(radius float64) (string, string) {
	area := math.Pi * radius * radius
	circumference := math.Pi * radius * 2

	// convert to string type
	areaStr := strconv.FormatFloat(area, 'f', 1, 64)
	circumferenceStr := strconv.FormatFloat(circumference, 'f', 1, 64)

	return areaStr, circumferenceStr
}

func main() {
    CircleInfo(12.3)
}
```

### Printf example
```go
package example
import "fmt"

func printf_exercise() {

	//%f = print a float value
	//%v = print the value in a default format
	//%T = print the type of the variable
	//%t = print a Boolean value as true or false
	var f float64 = 38730204.3832
	fmt.Printf("%v\n", f)
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)

	var i int = 13
	fmt.Printf("|%5d|\n", i)
	fmt.Printf("|%-5d|\n", i)

	var b bool = false
	fmt.Printf("%t\n", b)
	fmt.Printf("%T %T %T\n", f, i, b)

	// read from stdin
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(num)
}

```