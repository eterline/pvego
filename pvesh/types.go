package pvesh

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type VMID int

func (id VMID) String() string {
	return fmt.Sprintf("%d", id)
}

func (id VMID) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%d", id))
}

func (id *VMID) UnmarshalJSON(data []byte) error {

	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		*id = VMID(num)
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	parsedNum, err := strconv.Atoi(str)
	if err != nil {
		return err
	}

	*id = VMID(parsedNum)
	return nil
}

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
