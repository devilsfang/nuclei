package fs

import (
	lib_fs "github.com/devilsfang/nuclei/v3/pkg/js/libs/fs"

	"github.com/devilsfang/nuclei/v3/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/fs")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"ListDir":          lib_fs.ListDir,
			"ReadFile":         lib_fs.ReadFile,
			"ReadFileAsString": lib_fs.ReadFileAsString,
			"ReadFilesFromDir": lib_fs.ReadFilesFromDir,

			// Var and consts

			// Objects / Classes

		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
