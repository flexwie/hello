package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	fmt.Println(string(out))

	var outputString string = string(out)

	if strings.TrimSpace(outputString) != "Hello World" {
		t.Errorf("Expected %s, got %s", "Hello World", out)
	}
}
