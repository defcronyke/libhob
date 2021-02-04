/*Package root is a package for this project.

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
package root

import (
	"encoding/json"
	"fmt"

	"gitlab.com/defcronyke/libhob/src/constant"
	"gitlab.com/defcronyke/libhob/src/id"
	"gitlab.com/defcronyke/libhob/src/kind"
	"gitlab.com/defcronyke/libhob/src/result"
)

type HobRoot struct {
	GlobalID   id.HobID     `json:"global_id"`
	ParentKind kind.HobKind `json:"parent_kind"`
	Kind       kind.HobKind `json:"kind"`
	Name       string       `json:"name"`
	Tag        string       `json:"tag"`
}

func NewHobRoot(name string) HobRoot {
	newKind := kind.NewHobRootKind0()

	name2 := name

	if name2 == "" {
		name2 = constant.HOB_NAME_DEFAULT
	}

	tag := fmt.Sprintf("[%v::%v::%v::%v]", id.HobRoot, newKind.ParentKind, newKind.Kind, name2)

	res := HobRoot{
		GlobalID:   id.HobRoot,
		ParentKind: newKind.ParentKind,
		Kind:       newKind.Kind,
		Name:       name2,
		Tag:        tag,
	}

	fmt.Printf("New object created: %v\n", res.Tag)

	return res
}

func (h *HobRoot) Main() (result.HobResultSuccess, result.HobResultError, result.HobResultErrorCode) {
	res := fmt.Sprintf("%v says goodbye.", h.Name)

	return &res, nil, nil
}

func (h *HobRoot) String() string {
	bytes, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return fmt.Sprintf("%#v", h)
	}

	return fmt.Sprintf("%s", bytes)
}
