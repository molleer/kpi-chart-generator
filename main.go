package main

import (
	"log"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func main() {
	nameItems := []string{"One", "Two", "Three", "Four", "Five", "Six"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Title"})
	bar.AddXAxis(nameItems).
		AddYAxis("Hi", []int{1, 2, 3, 4, 5, 6}, charts.BarOpts{Stack: "stack"}).
		AddYAxis("Hello", []int{-6, -5, -4, -3, -2, -1}, charts.BarOpts{Stack: "stack"})
	f, _ := os.Create("bar.kpi.html")
	bar.Render(f)
	log.Println("Charts created")
}
