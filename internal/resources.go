package internal

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

//todo:
type Resources struct {
	root string
}

func (r *Resources) String() string {
	return "Resources{root=" + r.root + "}"
}

func (r *Resources) GetRoot() string {
	return r.root
}

//todo:
func InitResources(dir string) *Resources {
	fmt.Println("Initializing local configuration folder at " + dir)
	return &Resources{dir}
}

//todo:
func (r *Resources) WritePolicy(name string, content string) {
	p := filepath.Join(r.root, "sys/policy", AppendFileExtension(name, "hcl"))
	fmt.Println("Writing policy " + name + " to " + p)
	StringToFile(content, p)
}

//todo:
func (r *Resources) ForEachPolicy(callback func(name string, content string)) {
	dir := filepath.Join(r.root, "sys/policy")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err) //todo:
	}
	for _, file := range files {
		content, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			panic(err) //todo:
		}

		name := strings.TrimSuffix(file.Name(), path.Ext(file.Name()))

		callback(name, string(content))
	}
}
