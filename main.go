package main

import "fmt"

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

}
