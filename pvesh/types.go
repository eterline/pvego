package pvesh

import "strconv"

type AvgLoad [3]string

type AvgLoadData struct {
	Load1  float32 `json:"load-1"`
	Load5  float32 `json:"load-5"`
	Load15 float32 `json:"load-15"`
}

func (avg AvgLoad) Struct() AvgLoadData {

	value1, _ := strconv.ParseFloat(avg[0], 32)
	value5, _ := strconv.ParseFloat(avg[1], 32)
	value15, _ := strconv.ParseFloat(avg[2], 32)

	return AvgLoadData{
		Load1:  float32(value1),
		Load5:  float32(value5),
		Load15: float32(value15),
	}
}
