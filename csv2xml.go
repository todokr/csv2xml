package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReportError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func SplitAsColumn(line string) []string {
	return Map(strings.Split(line, ","), func(v string) string {
		return strings.Trim(v, "\"")
	})
}

func main() {
	var (
		file             = flag.String("f", "", "CSV file to convert")
		rootElemName     = flag.String("r", "root", "root element name of XML")
		eachLineElemName = flag.String("e", "item", "node name of each line")
	)
	flag.Parse()

	if *file == "---" {
		log.Fatal("declear -f option to open CSV")
	}

	csv, err := os.Open(*file)
	ReportError(err)
	defer csv.Close()

	reader := bufio.NewReader(csv)

	headerLine, _, err := reader.ReadLine()
	ReportError(err)

	headers := SplitAsColumn(string(headerLine))

	fmt.Println(`<?xml version="1.0" encoding="UTF-8" ?>`)
	fmt.Println("<" + *rootElemName + ">")

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			ReportError(err)
		}

		columns := SplitAsColumn(string(line))

		fmt.Println("  <" + *eachLineElemName + ">")
		for i, v := range columns {
			fmt.Print("    <" + headers[i] + ">")
			fmt.Print("<![CDATA[" + v + "]]>")
			fmt.Println("</" + headers[i] + ">")
		}
		fmt.Println("  </" + *eachLineElemName + ">")
	}

	fmt.Println("</" + *rootElemName + ">")
}
