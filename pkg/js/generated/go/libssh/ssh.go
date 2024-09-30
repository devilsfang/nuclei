package ssh

import (
	lib_ssh "github.com/devilsfang/nuclei/v3/pkg/js/libs/ssh"

	"github.com/devilsfang/nuclei/v3/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/ssh")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Objects / Classes
			"SSHClient": gojs.GetClassConstructor[lib_ssh.SSHClient](&lib_ssh.SSHClient{}),
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
