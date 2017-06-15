package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var arr = []string{
	"George", "Michael", "Robert", "Dennis", "Ian",
}

func main() {
	fmt.Println(printBoxes(arr, maxArrayLen(arr)))
}

func printArr(in []string) string {
	max := maxArrayLen(in)
	var ret string
	buf := bytes.NewBufferString(ret)
	str := strconv.Itoa(max)
	_, err := buf.WriteString(str)
	if err != nil {
		log.Println("Unable to write string: ", err)
	}
	return buf.String()
}

func maxArrayLen(in []string) (max int) {
	max = 0
	for _, v := range in {
		if len(v) > max {
			max = len(v)
		}
	}
	return max
}

func printBoxes(s []string, i int) string {
	var strbuff string
	ret := bytes.NewBufferString(strbuff)
	printBox(ret, i)
	for _, line := range s {
		pad := strings.Repeat(" ", i-len(line))
		ret.WriteString("|" + line + pad)
		ret.WriteString("|\n")
		printBox(ret, i)
	}
	return ret.String()
}

func printBox(buff *bytes.Buffer, i int) string {
	buff.WriteString("+")
	for l := 0; l < i; l++ {
		buff.WriteString("-")
	}
	buff.WriteString("+")
	buff.WriteString("\n")
	return buff.String()
}
