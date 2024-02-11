
# LRParser

Parse values easily using LR (left-right delimiter) with LRParser.




## Installation

`go get github.com/infectrs/lrparser`
## Usage

```go
package main

import (
    "fmt"
    "github.com/infectrs/lrparser"
)

func main() {
	input := `
	<h1>Infect</h1>
	<p id="information">
	Hello, this was made by Infect
	</p>
	`

	leftDelimiter := `<p id="information">`
	rightDelimiter := `</p>`

	parsedValue, err := lrparser.ParseLr(input, leftDelimiter, rightDelimiter)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedValue)
}
```