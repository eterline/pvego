package pvesh

import (
	"strings"
	"time"
)

type QemuVirtualMachine struct {
	CPU       int    `json:"cpu"`
	Cpus      int    `json:"cpus"`
	Disk      int    `json:"disk"`
	Diskread  int64  `json:"diskread"`
	Diskwrite int64  `json:"diskwrite"`
	Maxdisk   int64  `json:"maxdisk"`
	Maxmem    int64  `json:"maxmem"`
	Mem       int64  `json:"mem"`
	Name      string `json:"name"`
	Netin     int64  `json:"netin"`
	Netout    int64  `json:"netout"`
	Pid       int64  `json:"pid"`
	Status    string `json:"status"`
	Tags      string `json:"tags"`
	Uptime    int64  `json:"uptime"`
	Vmid      VMID   `json:"vmid"`

	api *Pvesh `json:"-"`
}

// QemuList get qemu vms list
func (sh *Pvesh) QemuList() ([]QemuVirtualMachine, error) {

	list := []QemuVirtualMachine{}

	if err := sh.Get("nodes", sh.Hostname, "qemu").
		Resolve(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// QemuByVmid get qemu vms list
func QemuByVmid(qemuList []QemuVirtualMachine, vmid int) (*QemuVirtualMachine, bool) {
	for _, vm := range qemuList {
		if vm.Vmid.Value() == vmid {
			return &vm, true
		}
	}
	return nil, false
}

func (ct *QemuVirtualMachine) IsRunning() bool {
	return ct.Status == "running"
}

func (ct *QemuVirtualMachine) UptimeDuration() time.Duration {
	return time.Duration(ct.Uptime)
}

func (ct *QemuVirtualMachine) TagList() []string {
	return strings.Split(ct.Tags, ";")
}

func (sh *QemuVirtualMachine) Start() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "qemu", sh.Vmid.String(), "status", "start",
	).Error
}

func (sh *QemuVirtualMachine) Stop() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "qemu", sh.Vmid.String(), "status", "stop",
	).Error
}

func (sh *QemuVirtualMachine) Shutdown() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "qemu", sh.Vmid.String(), "status", "shutdown",
	).Error
}

func (sh *QemuVirtualMachine) Reboot() error {
	return sh.api.Create(
		"nodes", sh.api.Hostname, "qemu", sh.Vmid.String(), "status", "reboot",
	).Error
}
