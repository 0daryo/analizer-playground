package main

import (
	"fmt"
	"io/ioutil"
)

func t() {
	fmt.Fprintln(ioutil.Discard, "hoge")
}
