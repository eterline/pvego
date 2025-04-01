package pvesh

import (
	"strings"
)

type QemuVirtualMachine struct {
	CPU       int    `json:"cpu"`
	Cpus      int    `json:"cpus"`
	Disk      int    `json:"disk"`
	Diskread  int    `json:"diskread"`
	Diskwrite int    `json:"diskwrite"`
	Maxdisk   int64  `json:"maxdisk"`
	Maxmem    int64  `json:"maxmem"`
	Mem       int    `json:"mem"`
	Name      string `json:"name"`
	Netin     int    `json:"netin"`
	Netout    int    `json:"netout"`
	Pid       int    `json:"pid"`
	Status    string `json:"status"`
	Tags      string `json:"tags"`
	Uptime    int    `json:"uptime"`
	Vmid      VMID   `json:"vmid"`

	api *Pvesh `json:"-"`
}

// QemuList - get qemu vms list
func (sh *Pvesh) QemuList() ([]QemuVirtualMachine, error) {

	list := []QemuVirtualMachine{}

	if err := sh.Get("nodes", sh.Hostname, "qemu").
		Resolve(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// QemuByVmid - get qemu vms list
func QemuByVmid(qemuList []QemuVirtualMachine, vmid int) (*QemuVirtualMachine, bool) {
	for _, vm := range qemuList {
		if vm.Vmid == VMID(vmid) {
			return &vm, true
		}
	}
	return nil, false
}

func (ct *QemuVirtualMachine) IsRunning() bool {
	return ct.Status == "running"
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
