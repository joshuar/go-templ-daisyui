// Code generated by "stringer -type=CardLayout -linecomment -output card_generated.go"; DO NOT EDIT.

package components

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CardLayoutNormal-0]
	_ = x[CardLayoutCompact-1]
	_ = x[CardLayoutSide-2]
	_ = x[CardLayoytSideResponsive-3]
}

const _CardLayout_name = "card-normalcard-compactcard-sidelg:card-side"

var _CardLayout_index = [...]uint8{0, 11, 23, 32, 44}

func (i CardLayout) String() string {
	if i < 0 || i >= CardLayout(len(_CardLayout_index)-1) {
		return "CardLayout(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CardLayout_name[_CardLayout_index[i]:_CardLayout_index[i+1]]
}
