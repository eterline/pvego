package pvesh

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// NodeStatus status info of proxmox host
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

// NodeStatus of proxmox node
func (sh *Pvesh) NodeStatus() (NodeStatus, error) {

	stats := NodeStatus{}

	if err := sh.Get("nodes", sh.Hostname, "status").
		Resolve(&stats); err != nil {
		return stats, err
	}

	return stats, nil
}

// =====================================

type NodeDnsInfo struct {
	Search string `json:"search,omitempty"`
	Dns1   string `json:"dns1,omitempty"`
	Dns2   string `json:"dns2,omitempty"`
	Dns3   string `json:"dns3,omitempty"`
}

// DnsAddrs list of dns addresses of host
func (dns NodeDnsInfo) DnsAddrs() []net.IP {
	return []net.IP{
		net.ParseIP(dns.Dns1),
		net.ParseIP(dns.Dns2),
		net.ParseIP(dns.Dns3),
	}
}

// DnsInfo read DNS settings
func (sh *Pvesh) Dns() (NodeDnsInfo, error) {
	infoList := NodeDnsInfo{}
	err := sh.Get("nodes", sh.Hostname, "dns").Resolve(&infoList)
	return infoList, err
}

// =====================================

type NodeHostsInfo struct {
	Data   string `json:"data"`
	Digest string `json:"digest"`
}

// FormatData hosts file ip map with names list
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

// =====================================

type NetstatInfo struct {
	Dev  string `json:"dev"`
	In   string `json:"in"`
	Out  string `json:"out"`
	Vmid VMID   `json:"vmid"`
}

// BytesIn of netstat info
func (stat NetstatInfo) BytesIn() int64 {
	i, err := strconv.ParseInt(stat.In, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// BytesOut of netstat info
func (stat NetstatInfo) BytesOut() int64 {
	i, err := strconv.ParseInt(stat.Out, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// Netstat returns network stats for vms and containers
func (sh *Pvesh) Netstat() ([]NetstatInfo, error) {
	infoList := []NetstatInfo{}
	err := sh.Get("nodes", sh.Hostname, "netstat").Resolve(&infoList)
	return infoList, err
}

// =====================================

// TimeInfo of proxmox host
type TimeInfo struct {
	Local    int64  `json:"localtime"`
	Time     int64  `json:"time"`
	TimeZone string `json:"timezone"`
}

// LocalIs returns local time of proxmox host
func (t TimeInfo) LocalIs() time.Time {
	return time.Unix(t.Local, 0)
}

// TimeIs returns system time of proxmox host
func (t TimeInfo) TimeIs() time.Time {
	return time.Unix(t.Time, 0)
}

// Time returns time info of proxmox host
func (sh *Pvesh) Time() (TimeInfo, error) {
	infoList := TimeInfo{}
	if err := sh.Get("nodes", sh.Hostname, "time").
		Resolve(&infoList); err != nil {
		return infoList, err
	}
	return infoList, nil
}

// =====================================

type NetworkConfig struct {
	Exists   ProxmoxBoolean `json:"exists"`
	Active   ProxmoxBoolean `json:"active"`
	Method   string         `json:"method"`
	Method6  string         `json:"method6"`
	Priority int            `json:"priority"`
	Type     string         `json:"type"`

	Families []string `json:"families"`
	Iface    string   `json:"iface"`

	Gateway string `json:"gateway"`
	Address string `json:"address"`
	CIDR    string `json:"cidr"`
	Netmask string `json:"netmask"`

	Autostart   ProxmoxBoolean `json:"autostart,omitempty"`
	BridgeFD    string         `json:"bridge_fd,omitempty"`
	BridgePorts string         `json:"bridge_ports,omitempty"`
	BridgeSTP   string         `json:"bridge_stp,omitempty"`
	Comments    string         `json:"comments,omitempty"`
}

func (cfg NetworkConfig) AddrGateway() net.IP {
	return net.ParseIP(cfg.Gateway)
}

func (cfg NetworkConfig) Addr() (net.IP, *net.IPNet, error) {
	return net.ParseCIDR(fmt.Sprintf("%s/%s", cfg.Address, cfg.Netmask))
}

func (cfg NetworkConfig) AddrIP() net.IP {
	return net.ParseIP(cfg.Address)
}

// Network returns network interfaces info of proxmox host
func (sh *Pvesh) Network() ([]NetworkConfig, error) {
	infoList := []NetworkConfig{}
	if err := sh.Get("nodes", sh.Hostname, "network").
		Resolve(&infoList); err != nil {
		return infoList, err
	}
	return infoList, nil
}
