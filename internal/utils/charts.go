package utils

import (
	"github.com/go-echarts/go-echarts/charts"
)

type DataSet struct {
	Name  string
	Data  interface{}
	Color string
}

func CreateBarChart(title string, columns []string, sets []DataSet) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: title})
	bar.AddXAxis(columns)

	for _, set := range sets {
		bar.AddYAxis(set.Name, set.Data, charts.BarOpts{Stack: "stack"}, charts.ColorOpts{set.Color})
	}

	return bar
}
