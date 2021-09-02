package crawler

import (
	"FOFNotifier/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	fundInfoUrl = "https://fundgz.1234567.com.cn/js/%v.js"
)

var Time string

func CrawlerByCode(code string, resMap *sync.Map, wg *sync.WaitGroup) {
	url := fmt.Sprintf(fundInfoUrl, code)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Http get code info error, err=%v, url=%v\n", err, url)
		wg.Done()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read http body error, err=%v\n", err)
		wg.Done()
		return
	}

	bodyStr := string(body)
	bodyStr = bodyStr[8 : len(bodyStr)-2]

	if bodyStr != "" {
		var r model.CrawlerRes
		err := json.Unmarshal([]byte(bodyStr), &r)
		if err != nil {
			fmt.Printf("Unmarshal error, err=%v, json=%v\n", err, bodyStr)
			wg.Done()
			return
		}
		Time = r.Time
		resMap.Store(code, &r)
	}

	wg.Done()
}
