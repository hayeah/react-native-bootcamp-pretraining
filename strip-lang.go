// Keep only what's between <cn>...</cn>
// go run strip-lang.go index{,.zh_CN}.md
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var (
	startTag   = "<cn>"
	endTag     = "</cn>"
	startTagRE = regexp.MustCompile(startTag)
	endTagRE   = regexp.MustCompile(endTag)
)

func main() {
	inputFileName := os.Args[1]

	input, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var output io.Writer
	if len(os.Args) >= 3 {
		outputFileName := os.Args[2]
		output, err = os.Create(outputFileName)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		output = os.Stdout
	}

	content, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print(string(content))

	for {
		startLoc := startTagRE.FindIndex(content)
		if startLoc == nil {
			break
		}

		sectionStart := startLoc[1]
		content = content[sectionStart:]

		endLoc := endTagRE.FindIndex(content)
		if endLoc == nil {
			log.Fatalln("no end tag found:", endTag)
		}

		sectionEnd := endLoc[0]
		// log.Println("sectionStart", startLoc)
		// log.Println("sectionEnd", endLoc)

		chunk := content[0:sectionEnd]
		output.Write(chunk)

		content = content[endLoc[1]:]
	}

}
