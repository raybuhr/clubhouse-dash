package main

import (
	"fmt"
	"os"

	"github.com/raybuhr/clubhouse"
	// "github.com/MCBrandenburg/clubhouse-go"
)

func main() {
	chToken := os.Getenv("CLUBHOUSE_API_TOKEN")
	projects, _ := clubhouse.GetProjects(chToken)
	epics, _ := clubhouse.GetEpics(chToken)

	fmt.Println("Projects:")
	fmt.Println(projects)
	fmt.Println("Epics:")
	fmt.Printf("%#v", epics)
}
