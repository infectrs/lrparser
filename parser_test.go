package lrparser

import (
	"fmt"
	"testing"
)

func TestLrParser(t *testing.T) {
	input := `
	<h1>Infect</h1>
	<p id="information">
	Hello, this was made by Infect
	</p>
	`

	leftDelimiter := `<p id="information">`
	rightDelimiter := `</p>`

	parsedValue, err := ParseLr(input, leftDelimiter, rightDelimiter)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(parsedValue)
}
