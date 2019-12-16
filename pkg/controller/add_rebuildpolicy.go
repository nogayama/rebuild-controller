package controller

import (
	"github.com/nogayama/rebuild-controller/pkg/controller/rebuildpolicy"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, rebuildpolicy.Add)
}
