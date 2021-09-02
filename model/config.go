package model

type GlobalConfig struct {
	Combinations *Combinations
}

type Combinations map[string]CombinationConfig // 组合名称：配置

type CombinationConfig map[string]float64 // 基金代码：占比

func (c Combinations) GetAllCodes() []string {
	resMap := make(map[string]int8)
	for _, cconfig := range c {
		for code, _ := range cconfig {
			if _, ok := resMap[code]; !ok {
				resMap[code] = 1
			}
		}
	}
	res := make([]string, 0)
	for code, _ := range resMap {
		res = append(res, code)
	}
	return res
}
