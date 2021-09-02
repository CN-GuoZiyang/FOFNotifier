package common

import (
	"FOFNotifier/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadCombinationConfig() (*model.Combinations, error) {
	combinations := &model.Combinations{}

	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Get home directory error! err=%v\n", err)
		return nil, err
	}
	path := filepath.Join(homePath, "fof.json")

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Open config file error! err=%v\n", err)
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Read config file error! err=%v\n", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, combinations)
	if err != nil {
		fmt.Printf("Unmarshal config error! err=%v\n", err)
		return nil, err
	}
	return combinations, nil
}
