package pvesh

import (
	"net"
	"strings"
	"time"
)

type LxcContainer struct {
	CPU       int    `json:"cpu"`
	Cpus      int    `json:"cpus"`
	Disk      int64  `json:"disk"`
	Diskread  int64  `json:"diskread"`
	Diskwrite int64  `json:"diskwrite"`
	Maxdisk   int64  `json:"maxdisk"`
	Maxmem    int64  `json:"maxmem"`
	Maxswap   int64  `json:"maxswap"`
	Mem       int64  `json:"mem"`
	Name      string `json:"name"`
	Netin     int64  `json:"netin"`
	Netout    int64  `json:"netout"`
	Pid       int64  `json:"pid,omitempty"`
	Status    string `json:"status"`
	Swap      int64  `json:"swap"`
	Tags      string `json:"tags"`
	Type      string `json:"type"`
	Uptime    int64  `json:"uptime"`
	Vmid      VMID   `json:"vmid"`

	api *Pvesh `json:"-"`
}

// LxcList - get lxc containers list
func (sh *Pvesh) LxcList() ([]LxcContainer, error) {

	list := []LxcContainer{}

	if err := sh.Get("nodes", sh.Hostname, "lxc").
		Resolve(&list); err != nil {
		return nil, err
	}

	return list, nil
}

func LxcByVmid(lxcList []LxcContainer, vmid int) (*LxcContainer, bool) {
	for _, container := range lxcList {
		if container.Vmid.Value() == vmid {
			return &container, true
		}
	}
	return nil, false
}

func (ct *LxcContainer) IsRunning() bool {
	return ct.Status == "running"
}

func (ct *LxcContainer) UptimeDuration() time.Duration {
	return time.Duration(ct.Uptime)
}

func (ct *LxcContainer) TagList() []string {
	return strings.Split(ct.Tags, ";")
}

func (sh *LxcContainer) Start() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "lxc", sh.Vmid.String(), "status", "start",
	).Error
}

func (sh *LxcContainer) Stop() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "lxc", sh.Vmid.String(), "status", "stop",
	).Error
}

func (sh *LxcContainer) Shutdown() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "lxc", sh.Vmid.String(), "status", "shutdown",
	).Error
}

func (sh *LxcContainer) Reboot() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "lxc", sh.Vmid.String(), "status", "reboot",
	).Error
}

type LxcСonfig struct {
	Arch        string `json:"arch,omitempty"`
	Tags        string `json:"tags,omitempty"`
	Description string `json:"description,omitempty"`
	Features    string `json:"features,omitempty"`

	Hostname string `json:"hostname,omitempty"`

	Cores  int   `json:"cores,omitempty"`
	Memory int64 `json:"memory,omitempty"`
	Swap   int64 `json:"swap,omitempty"`

	Onboot       int `json:"onboot,omitempty"`
	Unprivileged int `json:"unprivileged,omitempty"`
}

func (sh *LxcContainer) GetConfig() (LxcСonfig, error) {

	conf := LxcСonfig{}

	res := sh.api.Get("nodes", sh.api.Hostname, "lxc", sh.Vmid.String(), "config")

	if err := res.Resolve(&conf); err != nil {
		return conf, err
	}

	return conf, nil
}

func (sh *LxcContainer) Clone(node string, newId int, full bool) error {
	args := make(CommandArguments)
	isFull := 0

	if full {
		isFull = 1
	}

	args.AddInt("newid", newId)
	args.AddInt("full", isFull)

	return sh.api.CreateWith(args, "nodes", sh.api.Hostname, "lxc", sh.Vmid.String(), "clone").Error
}

type LxcInterface struct {
	Hwaddr string `json:"hwaddr"`
	Inet   string `json:"inet,omitempty"`
	Inet6  string `json:"inet6,omitempty"`
	Name   string `json:"name"`
}

func (i LxcInterface) Mac() (net.HardwareAddr, bool) {
	hw, err := net.ParseMAC(i.Hwaddr)
	if err != nil {
		return nil, false
	}
	return hw, true
}

func (i LxcInterface) Addr() net.IP {
	return net.ParseIP(i.Inet)
}

func (i LxcInterface) Addr6() net.IP {
	return net.ParseIP(i.Inet6)
}

func (sh *LxcContainer) Interfaces() ([]LxcInterface, error) {
	list := []LxcInterface{}

	if err := sh.api.Get("nodes", sh.api.Hostname, "lxc", sh.Vmid.String(), "interfaces").
		Resolve(&list); err != nil {
		return nil, err
	}

	return list, nil
}
