package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/net/html"
)

type Attributes map[string]string

var body []byte
var currentUrl string

func setHtmlRawBytes(urlString string) {
	resp, err := http.Get(urlString)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	// read body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	body = b
}

// only gets Header nodes
func GetHeaderNode(url string) *html.Node {
	if url != currentUrl {
		setHtmlRawBytes(url)
	}
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatalln("error parsing body")
	}

	head := doc.FirstChild
	for head != nil {
		if head.Type == html.ElementNode && head.Data == "head" {
			return head
		}
		head = head.NextSibling
	}
	return nil
}

// returns all nodes in the body
func GetRootNode(url string) *html.Node {
	if url != currentUrl {
		setHtmlRawBytes(url)
	}
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Fatalln("error parsing body")
	}
	return doc
}

// finds all Nodes based on their tag name ex <script> , <img>
func GetNodesByTagName(n *html.Node, tag string) []html.Node {
	var results []html.Node
	getNodesByTagName(n, tag, &results)
	return results
}

func getNodesByTagName(n *html.Node, tag string, results *[]html.Node) {
	if n.Type == html.ElementNode && n.Data == tag {
		*results = append(*results, *n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getNodesByTagName(c, tag, results)
	}
}

// returns all attributes of a given node
func GetTagAttributes(node *html.Node) Attributes {
	attrs := make(Attributes)
	for _, attr := range node.Attr {
		attrs[attr.Key] = attr.Val
	}
	return attrs
}

// return only the specified attribute
func GetAttrsByName(node *html.Node, attr string) string {
	attrs := GetTagAttributes(node)
	return attrs[attr]
}

// downloads file

func Wget(fileURL string, path string) {
	fmt.Println(fileURL, path)

	cmd := exec.Command("wget", fileURL, "-P", path, "-c")
	if err := cmd.Run(); err != nil {
		log.Println("error running wget: %v", err)
	}
}

func CreateOutputDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func ReadContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("failed to download %s: %s", url, resp.Status))
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
