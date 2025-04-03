package pvesh

type AptUpdateInfo struct {
	Title       string `json:"Title,omitempty"`
	Arch        string `json:"Arch,omitempty"`
	Description string `json:"Description,omitempty"`
	Origin      string `json:"Origin,omitempty"`

	OldVersion string `json:"OldVersion,omitempty"`
	Version    string `json:"Version,omitempty"`

	Package  string `json:"Package,omitempty"`
	Priority string `json:"Priority,omitempty"`
	Section  string `json:"Section,omitempty"`
}

// AptUpdate returns apt update repository list
func (sh *Pvesh) AptUpdate() ([]AptUpdateInfo, error) {
	list := []AptUpdateInfo{}

	if err := sh.Get("nodes", sh.Hostname, "apt", "update").
		Resolve(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// AptGetUpdate resynchronize the package index files from their sources (apt-get update)
func (sh *Pvesh) AptGetUpdate() error {
	return sh.Create("nodes", sh.Hostname, "apt", "update").Error
}
