package pvesh

import (
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
	Vmid VMID   `json:"vmid"`
}

func (s *NodeNetstat) VMID() int {
	return int(s.Vmid)
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

type NodeDnsAddr string

func (dns NodeDnsAddr) IP() (net.IP, bool) {
	if dns != "" {
		return net.ParseIP(string(dns)), true
	}
	return nil, false
}

type NodeDnsInfo struct {
	Search string      `json:"search,omitempty"`
	Dns1   NodeDnsAddr `json:"dns1,omitempty"`
	Dns2   NodeDnsAddr `json:"dns2,omitempty"`
	Dns3   NodeDnsAddr `json:"dns3,omitempty"`
}

// DnsInfo - read DNS settings
func (sh *Pvesh) Dns() (NodeDnsInfo, error) {
	infoList := NodeDnsInfo{}
	err := sh.Get("nodes", sh.Hostname, "dns").Resolve(&infoList)
	return infoList, err
}

type NodeHostsInfo struct {
	Data   string `json:"data"`
	Digest string `json:"digest"`
}

func (hosts NodeHostsInfo) FormatData() map[*net.IP][]string {
	mappedHosts := make(map[*net.IP][]string)

	lines := strings.Split(hosts.Data, "\n")
	reSpaceClean := regexp.MustCompile(`\s+`)

	for _, line := range lines {

		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.Contains(line, "#") {
			continue
		}

		line = reSpaceClean.ReplaceAllString(line, " ")
		splittedLine := strings.Split(line, " ")
		lineIp := net.ParseIP(splittedLine[0])

		mappedHosts[&lineIp] = splittedLine[1:]
	}

	return mappedHosts
}

// Hosts - read hosts file
func (sh *Pvesh) Hosts() (NodeHostsInfo, error) {
	infoList := NodeHostsInfo{}
	err := sh.Get("nodes", sh.Hostname, "hosts").Resolve(&infoList)
	return infoList, err
}

type NetstatInfo struct {
	Dev  string `json:"dev"`
	In   string `json:"in"`
	Out  string `json:"out"`
	Vmid VMID   `json:"vmid"`
}

func (stat NetstatInfo) BytesIn() int64 {
	i, err := strconv.ParseInt(stat.In, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func (stat NetstatInfo) BytesOut() int64 {
	i, err := strconv.ParseInt(stat.Out, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// Netstat - returns network stats for vms and containers
func (sh *Pvesh) Netstat() ([]NetstatInfo, error) {
	infoList := []NetstatInfo{}
	err := sh.Get("nodes", sh.Hostname, "netstat").Resolve(&infoList)
	return infoList, err
}

type TimeInfo struct {
	Local    int64  `json:"localtime"`
	Time     int64  `json:"time"`
	TimeZone string `json:"timezone"`
}

func (t TimeInfo) LocalIs() time.Time {
	return time.Unix(t.Local, 0)
}

func (t TimeInfo) TimeIs() time.Time {
	return time.Unix(t.Time, 0)
}

// Time - returns time of host
func (sh *Pvesh) Time() (TimeInfo, error) {
	infoList := TimeInfo{}
	err := sh.Get("nodes", sh.Hostname, "time").Resolve(&infoList)
	return infoList, err
}
