package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func UnminifyJS(path string) {
	cmd := exec.Command("js-beautify", "-r", path)
	cmd.Run()

}

func Grep(file string, pattern string) {
	cmd := exec.Command("rg", "-i", pattern, file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() != 1 {
				log.Println(err)
			}
		}
	}
}

func GrepAllFiles(directory string, patttern string, beautify bool) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range files {
		file := fmt.Sprintf("%s/%s", directory, f.Name())
		if beautify {
			fmt.Printf("beautifying %s \n", file)
			UnminifyJS(file)
		}

		fmt.Printf("grepping %s%s \n", file, RedColor)
		Grep(file, patttern)
	}
}
