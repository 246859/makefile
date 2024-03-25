package main

import "fmt"

var (
	AppName   string
	BuildTime string
	Version   string
)

func main() {
	fmt.Println(AppName, BuildTime, Version)
}
