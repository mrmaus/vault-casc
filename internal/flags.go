package internal

import "path/filepath"

//todo:
type Flags struct {
	VaultAddress string
	VaultToken   string
	Dir          string
}

//todo:
var flags Flags

//todo:
func GetFlags() *Flags {
	return &flags
}

//todo:
func (f *Flags) GetDir() string {
	abs, err := filepath.Abs(f.Dir)
	if err != nil {
		panic(err)
	}
	return abs
}
