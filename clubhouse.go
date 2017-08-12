// Package clubhouse contains functions for getting and parsing JSON
// from the API of Clubhouse project management software
package clubhouse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Label struct {
	Name string `json:"name"`
}

type ProjectIds struct {
	ID []int
}

type Stats struct {
	NumPoints           int `json:"num_points"`
	NumStoriesDone      int `json:"num_stories_done"`
	NumStoriesStarted   int `json:"num_stories_started"`
	NumStoriesUnstarted int `json:"num_stories_unstarted"`
}

type Epic struct {
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	State      string        `json:"state"`
	CreatedAt  string        `json:"created_at"`
	UpdatedAt  string        `json:"updated_at"`
	ProjectIds []*ProjectIds `json:"project_ids, omitempty"`
	Labels     []*Label      `json:"labels, omitempty"`
	Stats      *Stats        `json:"stats"`
}

func GetProjects(token string) ([]Project, error) {
	url := fmt.Sprintf("https://api.clubhouse.io/api/beta/projects?token=%s", token)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting projects:", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading projects:", err)
	}

	projects := []Project{}
	json.Unmarshal(body, &projects)

	return projects, nil
}

func GetEpics(token string) ([]Epic, error) {
	url := fmt.Sprintf("https://api.clubhouse.io/api/beta/epics?token=%s", token)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting epics:", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading epics data:", err)
	}

	epics := []Epic{}
	json.Unmarshal(body, &epics)

	return epics, nil
}
