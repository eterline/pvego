package pvesh

import (
	"encoding/json"
	"os/exec"
	"strings"
)

type PveshCallResponse struct {
	ResponseBytes []byte
	Error         error
}

func (r PveshCallResponse) Resolve(v any) error {

	if r.Error != nil {
		return r.Error
	}

	return json.Unmarshal(r.ResponseBytes, v)
}

func NewPveshCallResponse(
	b []byte,
	e error,
) PveshCallResponse {
	return PveshCallResponse{
		ResponseBytes: b,
		Error:         e,
	}
}

// ========== Class methods ==========

func (sh *Pvesh) fetch(name string, arg ...string) PveshCallResponse {

	arg = append(arg, "--output-format", FormatJSON)

	exc := exec.CommandContext(sh.ctx, name, arg...)

	out, err := exc.CombinedOutput()
	if err != nil {
		return NewPveshCallResponse(nil, err)
	}

	return NewPveshCallResponse(out, nil)
}

// Get - Call API GET on <path>.
func (sh *Pvesh) Get(path ...string) PveshCallResponse {
	return sh.fetch(sh.command, "get", joinPath(path...))
}

// Get - Call API POST on <path>.
func (sh *Pvesh) Create(path ...string) PveshCallResponse {
	return sh.fetch(sh.command, "create", joinPath(path...))
}

// Get - Call API DELETE on <path>.
func (sh *Pvesh) Delete(path ...string) PveshCallResponse {
	return sh.fetch(sh.command, "delete", joinPath(path...))
}

// Get - Call API PUT on <path>.
func (sh *Pvesh) Set(path ...string) PveshCallResponse {
	return sh.fetch(sh.command, "set", joinPath(path...))
}

func joinPath(path ...string) string {
	return "/" + strings.Join(path, "/")
}
