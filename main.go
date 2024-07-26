package main

import "fmt"

var (
	Author    string
	AppName   string
	BuildTime string
	Version   string
)

func main() {
	fmt.Println(Author, AppName, BuildTime, Version)
}
