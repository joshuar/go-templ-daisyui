// Code generated by "stringer -type=ObjectFit -linecomment -output fit_generated.go"; DO NOT EDIT.

package components

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ObjectNone-1]
	_ = x[ObjectContain-2]
	_ = x[ObjectCover-3]
	_ = x[ObjectFill-4]
	_ = x[ObjectScaleDown-5]
}

const _ObjectFit_name = "object-noneobject-containobject-coverobject-fillobject-scale-down"

var _ObjectFit_index = [...]uint8{0, 11, 25, 37, 48, 65}

func (i ObjectFit) String() string {
	i -= 1
	if i < 0 || i >= ObjectFit(len(_ObjectFit_index)-1) {
		return "ObjectFit(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ObjectFit_name[_ObjectFit_index[i]:_ObjectFit_index[i+1]]
}