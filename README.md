# jsfinder

## Description
A web pentest program that finds <script> tags in a html file and downloads them to find specific keywords. Forexample you may want to check if your website is leaking passwords or tokens. Using this program you can automate the whole process,

## Installation and Usage 
1. Clone the repo
2. Run `go build` in the project directory 
3. ```./jsfinder -url <target url> -d <download directory> -b <beautifies javascript files> -p <pattern>```

## Dependancy
this program relies on a few programs
1. wget
2. rg (better grep)
3. js-beautify
