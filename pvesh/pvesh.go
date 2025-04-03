package pvesh

import (
	"context"
	"fmt"
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

// New creates pvesh instance
func New() (*Pvesh, error) {
	return NewWithContext(
		context.Background(),
	)
}

// New creates pvesh instance
func NewWithContext(ctx context.Context) (*Pvesh, error) {

	bin, err := exec.LookPath(MainCommand)
	if err != nil {
		return nil, fmt.Errorf("couldn't find pvesh bin: %w", err)
	}

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &Pvesh{
		Hostname: host,
		root:     bin,
		ctx:      ctx,
	}, nil
}

// ========== Class methods ==========

// Version get proxmox ve version info
func (sh *Pvesh) Version() (PveSystemVersion, error) {
	data := PveSystemVersion{}

	if err := sh.Get("version").
		Resolve(&data); err != nil {
		return data, err
	}

	return data, nil
}
