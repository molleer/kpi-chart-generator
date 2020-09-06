package main

import (
	"log"
	"os"
	"time"

	"github.com/molleer/kpi-chart-generator/internal/utils"
	//"github.com/go-echarts/go-echarts/charts"
)

func main() {
	nameItems := []string{"One", "Two", "Three", "Four", "Five", "Six"}
	setOne := utils.DataSet{"Hi", []int{1, 2, 3, 4, 5, 6}, "red"}
	setTwo := utils.DataSet{"Hello", []int{-1, -2, -3, -4, -5, -6}, "green"}
	bar := utils.CreateBarChart("Title", nameItems, []utils.DataSet{setOne, setTwo})

	f, _ := os.Create("bar.kpi.html")
	bar.Render(f)
	log.Println("Charts created")

	if os.Getenv("STAY_OPEN") == "True" {
		time.Sleep(time.Hour)
	}
}
