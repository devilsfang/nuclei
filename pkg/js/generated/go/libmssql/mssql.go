package mssql

import (
	lib_mssql "github.com/devilsfang/nuclei/v3/pkg/js/libs/mssql"

	"github.com/devilsfang/nuclei/v3/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/mssql")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions

			// Var and consts

			// Objects / Classes
			"MSSQLClient": gojs.GetClassConstructor[lib_mssql.MSSQLClient](&lib_mssql.MSSQLClient{}),
		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
