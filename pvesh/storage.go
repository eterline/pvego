package pvesh

type PveSystemVersion struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}

type PveStorageInfo struct {
	Content  string `json:"content"`
	Digest   string `json:"digest"`
	Nodes    string `json:"nodes"`
	Storage  string `json:"storage"`
	ThinPool string `json:"thinpool"`
	Type     string `json:"type"`
	VGname   string `json:"vgname"`
}
