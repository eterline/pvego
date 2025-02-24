package pvesh

import (
	"context"
	"os"
	"os/exec"
)

const (
	MainCommand = "pvesh"

	FormatJSON = "json" // command postfix arg json format
	FormatYAML = "yaml" // command postfix arg yaml format
)

// ========== Class declare ==========

// Pvesh - respresnts ProxmoVE 'pvesh' object for CLI API implementation.
// Shell interface for the Proxmox VE API.
type Pvesh struct {
	Hostname string // Hostname of proxmox host
	command  string // pvesh root dir

	ctx context.Context
}

// ========== Class constructor ==========

func New() (*Pvesh, error) {
	return NewWithContext(
		context.Background(),
	)
}

func NewWithContext(ctx context.Context) (*Pvesh, error) {

	shPath, err := exec.LookPath(MainCommand)
	if err != nil {
		return nil, err
	}

	hName, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Pvesh{
		Hostname: hName,
		command:  shPath,

		ctx: ctx,
	}, nil
}

// =======================================

func (sh *Pvesh) Version() (PveSystemVersion, error) {

	data := PveSystemVersion{}

	if err := sh.Get("version").
		Resolve(&data); err != nil {
		return data, err
	}

	return data, nil
}

// =======================================

func (sh *Pvesh) Storage() ([]PveStorageInfo, error) {

	infoList := []PveStorageInfo{}

	if err := sh.Get("storage").
		Resolve(&infoList); err != nil {
		return nil, err
	}

	return infoList, nil
}

// ==============================================================================

func (sh *Pvesh) Lxc() ([]LxcContainer, error) {

	list := []LxcContainer{}

	if err := sh.Get("nodes", sh.Hostname, "lxc").
		Resolve(&list); err != nil {
		return nil, err
	}

	for i, data := range list {
		data.api = sh
		list[i] = data
	}

	return list, nil
}

// ======================================

func (sh *Pvesh) Qemu() ([]QemuVirtualMachine, error) {

	list := []QemuVirtualMachine{}

	if err := sh.Get("nodes", sh.Hostname, "qemu").
		Resolve(&list); err != nil {
		return nil, err
	}

	for i, data := range list {
		data.api = sh
		list[i] = data
	}

	return list, nil
}
