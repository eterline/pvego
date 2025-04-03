package pvesh

type Storage struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	Storage string `json:"storage"`

	Active  ProxmoxBoolean `json:"active,omitempty"`
	Enabled ProxmoxBoolean `json:"enabled,omitempty"`
	Shared  ProxmoxBoolean `json:"shared,omitempty"`

	Avail        int64   `json:"avail,omitempty"`
	Total        int64   `json:"total,omitempty"`
	Used         int64   `json:"used,omitempty"`
	UsedFraction float64 `json:"used_fraction,omitempty"`

	api *Pvesh `json:"-"`
}

// StorageList get proxmox storage list
func (sh *Pvesh) StorageList() ([]Storage, error) {
	list := []Storage{}

	if err := sh.Get("nodes", sh.Hostname, "storage").
		Resolve(&list); err != nil {
		return nil, err
	}

	return list, nil
}

// StorageList get proxmox storage list
func StorageByName(list []Storage, name string) (*Storage, bool) {
	for _, st := range list {
		if st.Storage == name {
			return &st, true
		}
	}
	return nil, false
}
