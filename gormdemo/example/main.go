package main

import (
	"gormdemo/example/sample"
)

func main() {

	m := sample.NewSample()
	// info := m.GetInfo(1)
	// fmt.Print(info)

	data := sample.Sample{
		Title: "cesshi",
	}
	m.AddInfo(data)
}
