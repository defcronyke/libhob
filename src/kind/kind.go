/*Package kind is a package for this project.

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
package kind

import (
	"reflect"
	"strings"
)

type HobKind = string

type HobKindIntr interface{}

type HobRootKind struct {
	ParentKind HobKind
	Kind       HobKind
}

func NewHobRootKind() HobRootKind {
	self := HobRootKind{}
	parentStr := ""

	selfStr := reflect.TypeOf(self).String()
	selfStrParts := strings.Split(selfStr, ".")
	if len(selfStrParts) == 2 {
		selfStr = selfStrParts[1]
	}

	return HobRootKind{
		ParentKind: parentStr,
		Kind:       selfStr,
	}
}

type HobRootKind0 struct {
	ParentKind HobKind
	Kind       HobKind
}

func NewHobRootKind0() HobRootKind0 {
	parent := HobRootKind{}
	self := HobRootKind0{}

	parentStr := reflect.TypeOf(parent).String()
	parentStrParts := strings.Split(parentStr, ".")
	if len(parentStrParts) == 2 {
		parentStr = parentStrParts[1]
	}

	selfStr := reflect.TypeOf(self).String()
	selfStrParts := strings.Split(selfStr, ".")
	if len(selfStrParts) == 2 {
		selfStr = selfStrParts[1]
	}

	return HobRootKind0{
		ParentKind: parentStr,
		Kind:       selfStr,
	}
}
