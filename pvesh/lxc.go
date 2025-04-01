package pvesh

import (
	"strings"
)

type LxcContainer struct {
	CPU       int    `json:"cpu"`
	Cpus      int    `json:"cpus"`
	Disk      int64  `json:"disk"`
	Diskread  int64  `json:"diskread"`
	Diskwrite int    `json:"diskwrite"`
	Maxdisk   int64  `json:"maxdisk"`
	Maxmem    int64  `json:"maxmem"`
	Maxswap   int    `json:"maxswap"`
	Mem       int    `json:"mem"`
	Name      string `json:"name"`
	Netin     int    `json:"netin"`
	Netout    int    `json:"netout"`
	Pid       int    `json:"pid,omitempty"`
	Status    string `json:"status"`
	Swap      int    `json:"swap"`
	Tags      string `json:"tags"`
	Type      string `json:"type"`
	Uptime    int    `json:"uptime"`
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
		if container.Vmid == VMID(vmid) {
			return &container, true
		}
	}
	return nil, false
}

func (ct *LxcContainer) IsRunning() bool {
	return ct.Status == "running"
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
