package main

import (
	"flag"
	"fmt"
)

const (
	RedColor    = "\033[31m"
	GreenColor  = "\033[32m"
	YellowColor = "\033[33m"
	ResetColor  = "\033[0m"
)

func main() {
	flUrl := flag.String("url", "", "The target url")
	flOutput := flag.String("d", "", "The target folder")
	flPattern := flag.String("p", "", "Pattern")
	flBeautify := flag.Bool("b", false, "Beautify scripts")

	flag.Parse()

	if *flUrl == "" {
		fmt.Println("Scanning javasript files located in", *flOutput, RedColor)
		GrepAllFiles(*flOutput, *flPattern, *flBeautify)

	} else {
		fmt.Println("output", *flOutput)
		err := CreateOutputDir(*flOutput)
		if err != nil {
			panic(err)
		}

		fmt.Println("Finding javascript for "+*flUrl, RedColor)
		scrpts := GetAllScripts(*flUrl)
		for i, s := range scrpts {
			fmt.Printf("%d: %s\n", i, s)

			Wget(s, *flOutput)

		}

		GrepAllFiles(*flOutput, *flPattern, *flBeautify)
	}

}
