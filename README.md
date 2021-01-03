# hannamil/mapper
---
A simple go `struct` mapper.
### 1. Install
```bash
go get -u github.com/hannamil/mapper
```

### 2. Getting Started

```go
package main

import (
	"fmt"
	"time"

	"github.com/hannamil/mapper"
)

type (
	Class struct {
		Name  string
		Level string
	}
	
	Person struct {
		Name  string
		Age   int
		Birth time.Time
		Class Class
	}

	Student struct {
		Name  string
		Age   int
		Class interface{}
	}
)

func main() {
	person := &Person{
		Name: "Chris han",
		Age:  40,
		Class: Class{
			Name: "algorithm",
			Level: "basic",
        },
	}
	student, err := mapper.Mapper(person, &Student{})
	// or
	// student, err := mapper.Mapper(person, Student{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("student: %v\n", student)
}
```
### 3. output
```bash
student: {Chris han 40 {algorithm basic}}
```