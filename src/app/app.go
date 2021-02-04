/*Package app is a package for this project.

Copyright (c) 2021 Jeremy Carter <jeremy@jeremycarter.ca>

All uses of this project in part or in whole are governed
by the terms of the license contained in the file titled
"LICENSE" that's distributed along with the project, which
can be found in the top-level directory of this project.

If you don't agree to follow those terms or you won't
follow them, you are not allowed to use this project or
anything that's made with parts of it at all. The project
is also	depending on some third-party technologies, and
some of those are governed by their own separate licenses,
so furthermore, whenever legally possible, all license
terms from all of the different technologies apply, with
this project's license terms taking first priority.
*/
package app

import (
	"encoding/json"
	"fmt"

	"gitlab.com/defcronyke/libhob/src/id"
	"gitlab.com/defcronyke/libhob/src/kind"
	"gitlab.com/defcronyke/libhob/src/result"
	"gitlab.com/defcronyke/libhob/src/root"
)

type HobAppIntr interface {
	Main() (result.HobResultSuccess, result.HobResultError, result.HobResultErrorCode)
}

type HobApp struct {
	GlobalID   id.HobID     `json:"global_id"`
	ParentKind kind.HobKind `json:"parent_kind"`
	Kind       kind.HobKind `json:"kind"`
	Name       string       `json:"name"`
	Tag        string       `json:"tag"`
}

func NewHobApp(name string) HobApp {
	root := root.NewHobRoot(name)

	res := HobApp{
		GlobalID:   root.GlobalID,
		ParentKind: root.ParentKind,
		Kind:       root.Kind,
		Name:       root.Name,
		Tag:        root.Tag,
	}

	return res
}

func (h *HobApp) Main() (result.HobResultSuccess, result.HobResultError, result.HobResultErrorCode) {
	err := fmt.Errorf("You need to implement `HobApp::main()`, otherwise " +
		"you will be seeing this error message right now, and nothing " +
		"useful is going to happen.")

	code := 1

	return nil, err, &code
}

func (h *HobApp) String() string {
	bytes, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", h)
	}

	return fmt.Sprintf("%s", bytes)
}
