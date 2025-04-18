package pvesh

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type PveshCallResponse struct {
	ResponseBytes []byte
	Error         error
}

func (r PveshCallResponse) Resolve(v any) error {

	if r.Error != nil {
		return fmt.Errorf("pvesh response error: %w", r.Error)
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

	fmt.Println(name, arg)

	exc := exec.CommandContext(sh.ctx, name, arg...)
	out, err := exc.CombinedOutput()

	return NewPveshCallResponse(out, err)
}

// Get Call API GET on <path>.
func (sh *Pvesh) Get(path ...string) PveshCallResponse {
	return sh.fetch(sh.root, "get", joinPath(path...))
}

// Create Call API POST on <path>.
func (sh *Pvesh) Create(path ...string) PveshCallResponse {
	return sh.fetch(sh.root, "create", joinPath(path...))
}

// Delete Call API DELETE on <path>.
func (sh *Pvesh) Delete(path ...string) PveshCallResponse {
	return sh.fetch(sh.root, "delete", joinPath(path...))
}

// Set Call API PUT on <path>.
func (sh *Pvesh) Set(path ...string) PveshCallResponse {
	return sh.fetch(sh.root, "set", joinPath(path...))
}

type CommandArguments map[string]string

func (args CommandArguments) AddInt(key string, value int) {
	args[key] = strconv.Itoa(value)
}

func (args CommandArguments) AddString(key, value string) {
	args[key] = value
}

func (args CommandArguments) Format() string {
	values := []string{}

	for key, value := range args {
		values = append(values, fmt.Sprintf("--%s", key), value)
	}

	return strings.Join(values, " ")
}

// Create Call API POST on <path>.
func (sh *Pvesh) CreateWith(args CommandArguments, path ...string) PveshCallResponse {
	if args != nil {
		return sh.fetch(sh.root, "create", joinPath(path...), args.Format())
	}
	return sh.fetch(sh.root, "create", joinPath(path...))
}

func joinPath(path ...string) string {
	return "/" + strings.Join(path, "/")
}
