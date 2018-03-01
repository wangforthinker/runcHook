package hook

import (
	"encoding/json"
	"os"
)

func GetHookState() (*HookState, error) {
	s := &HookState{}
	err := json.NewDecoder(os.Stdin).Decode(s)
	if err != nil {
		return nil,err
	}

	return s,nil
}
