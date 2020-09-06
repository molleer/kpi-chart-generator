package main

import (
	"log"
	"os"
	"time"

	"github.com/molleer/kpi-chart-generator/internal/utils"
	//"github.com/go-echarts/go-echarts/charts"
)

func main() {
	contribs, err := utils.GetLogs()
	if err != nil {
		log.Print("Command failed")
		log.Println(err)
		return
	}

	nameItems := []string{}

	deletions := []int64{}
	insertions := []int64{}

	for _, contr := range contribs {
		nameItems = append(nameItems, contr.Email)
		deletions = append(deletions, -contr.Deletions)
		insertions = append(insertions, contr.Insertions)
	}

	setOne := utils.DataSet{"Insertions", insertions, "red"}
	setTwo := utils.DataSet{"Deletions", deletions, "green"}

	bar := utils.CreateBarChart("Contributions", nameItems, []utils.DataSet{setOne, setTwo})

	f, _ := os.Create("bar.kpi.html")
	bar.Render(f)
	log.Println("Charts created")

	if os.Getenv("STAY_OPEN") == "True" {
		time.Sleep(time.Hour)
	}
}
