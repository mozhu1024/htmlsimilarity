package similarity

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserHTML(t *testing.T) {
	b, err := ioutil.ReadFile("./tests/404.html") // just pass the file name
	if err != nil {
		t.Fatal(err)
	}
	content := string(b)
	parserHTML(content)
}

func TestGetSimilar(t *testing.T) {
	a, err := ioutil.ReadFile("./tests/404.html") // just pass the file name
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadFile("./tests/1.html") // just pass the file name
	if err != nil {
		t.Fatal(err)
	}
	c, err := ioutil.ReadFile("./tests/2.html") // just pass the file name
	if err != nil {
		t.Fatal(err)
	}
	isSimilar, value := GetSimilar(string(a), string(b))
	fmt.Println("0.html & 1.html - isSimilar : ", isSimilar, "value : ", value)

	isSimilar, value = GetSimilar(string(b), string(c))
	fmt.Println("1.html & 2.html - isSimilar : ", isSimilar, "value : ", value)
}
