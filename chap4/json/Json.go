package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var movies = []Movie{
		{
			Title: "复仇者联盟",
			Year: 2017,
			Color: 100,
			Actors: []string{
				"美国队长", "钢铁侠",
			},
		},
	}

	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Printf("json marshal failed: %s", err)
	}

	fmt.Printf("data: %s", data)
}

type Movie struct {
	Title string
	Year int `json:"released"`
	Color int `json:"color,omitempty"`
	Actors []string
}
