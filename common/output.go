package common

import (
	"FOFNotifier/crawler"
	"FOFNotifier/model"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteDetail(detail map[string]*model.Report) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Get home directory error! err=%v\n", err)
		return
	}
	path := filepath.Join(homePath, "fof_detail.txt")

	content := fmt.Sprintf("数据刷新时间：%v\n", crawler.Time)
	for comName, report := range detail {
		content = content + fmt.Sprintf("- %v: %.2f%%\n", comName, report.Rate)
		for _, fund := range report.Children {
			content = content + fmt.Sprintf("\t- %s:\t估值变动: %.2f%%\t份额: %.2f%%\t贡献: %.2f%%\n",
				fund.Name,
				fund.Rate,
				fund.Share/100.0,
				fund.Contribution,
			)
		}
	}

	err = ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Write detail to file error, err=%v\n", err)
	}
}
