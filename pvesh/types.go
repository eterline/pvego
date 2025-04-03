package pvesh

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// VMID - Virtual Machine ID in Proxmox is a unique id assigned to each virtual machine in a Proxmox cluster
type VMID int

// ParseVMID - parses vmid from string
func ParseVMID(value string) (vmid VMID, err error) {
	value = strings.TrimSpace(value)

	id, err := strconv.Atoi(value)
	if err != nil {
		return vmid, fmt.Errorf("couldn't parse vmid: %w", err)
	}

	if err := VMID(id).validate(); err != nil {
		return vmid, err
	}

	return VMID(id), nil
}

// validate - validates vmid value range
func (id VMID) validate() error {
	if id < 100 || id > 999999999 {
		return errors.New("uncorrect vmid value: must be above 100 and below 999999999")
	}
	return nil
}

// Format - format value into string: "ID: %d"
func (id VMID) Format(format string) string {
	return fmt.Sprintf(format, id)
}

// Value - returns integer type value
func (id VMID) Value() int {
	return int(id)
}

// Value - returns integer 64 type value
func (id VMID) Value64() int64 {
	return int64(id)
}

// String - returns value id as a string
func (id VMID) String() string {
	return fmt.Sprintf("%d", id)
}

func (id VMID) MarshalJSON() ([]byte, error) {
	if err := id.validate(); err != nil {
		return nil, err
	}

	return json.Marshal(int(id))
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

	if err := id.validate(); err != nil {
		return err
	}

	return nil
}

// ProxmoxBoolean logic value state for json parsing.
// 1 or 0 => true or false
type ProxmoxBoolean bool

func (boolean ProxmoxBoolean) Value() bool {
	return bool(boolean)
}

func (boolean ProxmoxBoolean) MarshalJSON() ([]byte, error) {
	if boolean {
		return []byte{'1'}, nil
	}
	return []byte{'0'}, nil
}

func (boolean *ProxmoxBoolean) UnmarshalJSON(data []byte) error {
	*boolean = data[0] == '1'
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

type PveSystemVersion struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}
