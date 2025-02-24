package pvesh

import "strconv"

type NodeTime struct {
	Localtime int    `json:"localtime"`
	Time      int    `json:"time"`
	Timezone  string `json:"timezone"`
}

func (sh *Pvesh) NodeTime() (NodeTime, error) {

	list := NodeTime{}

	if err := sh.Get("nodes", sh.Hostname, "time").
		Resolve(&list); err != nil {
		return list, err
	}

	return list, nil
}

type NodeNetstat struct {
	Dev  string `json:"dev"`
	In   string `json:"in"`
	Out  string `json:"out"`
	Vmid string `json:"vmid"`
}

func (s *NodeNetstat) VMID() int {
	v, _ := strconv.Atoi(s.Vmid)
	return v
}

func (sh *Pvesh) NodeNetstat() ([]NodeNetstat, error) {

	list := []NodeNetstat{}

	if err := sh.Get("nodes", sh.Hostname, "netstat").
		Resolve(&list); err != nil {
		return list, err
	}

	return list, nil
}

type NodeStatus struct {
	BootInfo struct {
		Mode       string `json:"mode"`
		Secureboot int    `json:"secureboot"`
	} `json:"boot-info"`
	CPU     int `json:"cpu"`
	Cpuinfo struct {
		Cores   int    `json:"cores"`
		Cpus    int    `json:"cpus"`
		Flags   string `json:"flags"`
		Hvm     string `json:"hvm"`
		Mhz     string `json:"mhz"`
		Model   string `json:"model"`
		Sockets int64  `json:"sockets"`
		UserHz  int64  `json:"user_hz"`
	} `json:"cpuinfo"`
	CurrentKernel struct {
		Machine string `json:"machine"`
		Release string `json:"release"`
		Sysname string `json:"sysname"`
		Version string `json:"version"`
	} `json:"current-kernel"`
	Idle int64 `json:"idle"`
	Ksm  struct {
		Shared int64 `json:"shared"`
	} `json:"ksm"`
	Kversion string  `json:"kversion"`
	Loadavg  AvgLoad `json:"loadavg"`
	Memory   struct {
		Free  int64 `json:"free"`
		Total int64 `json:"total"`
		Used  int64 `json:"used"`
	} `json:"memory"`
	Pveversion string `json:"pveversion"`
	Rootfs     struct {
		Avail int64 `json:"avail"`
		Free  int64 `json:"free"`
		Total int64 `json:"total"`
		Used  int64 `json:"used"`
	} `json:"rootfs"`
	Swap struct {
		Free  int64 `json:"free"`
		Total int64 `json:"total"`
		Used  int64 `json:"used"`
	} `json:"swap"`
	Uptime int64 `json:"uptime"`
	Wait   int64 `json:"wait"`
}

func (sh *Pvesh) NodeStatus() (NodeStatus, error) {

	stats := NodeStatus{}

	if err := sh.Get("nodes", sh.Hostname, "status").
		Resolve(&stats); err != nil {
		return stats, err
	}

	return stats, nil
}
