package handler

import (
	"FOFNotifier/common"
	"FOFNotifier/crawler"
	"FOFNotifier/model"
	"fmt"
	"strconv"
	"sync"
)

func MainRoutine() {
	config := common.GetGlobalConfig()

	combinations := config.Combinations
	codes := combinations.GetAllCodes()
	var wg sync.WaitGroup
	var resMap sync.Map
	for _, code := range codes {
		wg.Add(1)
		go crawler.CrawlerByCode(code, &resMap, &wg)
	}
	wg.Wait()

	detailResult := make(map[string]*model.Report)
	for comName, comConfig := range *combinations {
		comReport := &model.Report{
			Name:     comName,
			Rate:     0,
			Children: []*model.Report{},
		}
		detailResult[comName] = comReport
		for code, share := range comConfig {
			rawFundInfo, ok := resMap.Load(code)
			if !ok {
				continue
			}
			fundInfo := rawFundInfo.(*model.CrawlerRes)
			rate, _ := strconv.ParseFloat(fundInfo.Rate, 64)
			contribution := rate * share / 100.0
			comReport.Rate = comReport.Rate + contribution

			comReport.Children = append(comReport.Children, &model.Report{
				Name:         fundInfo.Name,
				Code:         code,
				Rate:         rate,
				Share:        share / 100.0,
				Contribution: contribution,
				Children:     nil,
			})
		}
	}

	go common.WriteDetail(detailResult)

	fmt.Printf("\n\n\n数据刷新时间：%v\n", crawler.Time)
	for comName, report := range detailResult {
		if report.Rate > 0 {
			fmt.Printf("- %s:\t+%.2f %%\n", comName, report.Rate)
		} else {
			fmt.Printf("- %s:\t%.2f %%\n", comName, report.Rate)
		}
	}
}
