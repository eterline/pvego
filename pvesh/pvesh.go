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
	root     string // pvesh root command dir

	ctx context.Context
}

// ========== Class constructor ==========

func New() (*Pvesh, error) {
	return NewWithContext(
		context.Background(),
	)
}

// ========== Class methods ==========

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
		root:     shPath,

		ctx: ctx,
	}, nil
}

// =======================================

// Version - get proxmox ve version info
func (sh *Pvesh) Version() (PveSystemVersion, error) {

	data := PveSystemVersion{}

	if err := sh.Get("version").
		Resolve(&data); err != nil {
		return data, err
	}

	return data, nil
}

// =======================================
