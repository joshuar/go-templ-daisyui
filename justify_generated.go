// Code generated by "stringer -type=JustifyContent -linecomment -output justify_generated.go"; DO NOT EDIT.

package components

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[JustifyNormal-0]
	_ = x[JustifyStart-1]
	_ = x[JustifyEnd-2]
	_ = x[JustifyCenter-3]
	_ = x[JustifyBetween-4]
	_ = x[JustifyAround-5]
	_ = x[JustifyEvenly-6]
	_ = x[JustifyStretch-7]
}

const _JustifyContent_name = "justify-startjustify-endjustify-centerjustify-betweenjustify-aroundjustify-evenlyjustify-stretch"

var _JustifyContent_index = [...]uint8{0, 0, 13, 24, 38, 53, 67, 81, 96}

func (i JustifyContent) String() string {
	if i < 0 || i >= JustifyContent(len(_JustifyContent_index)-1) {
		return "JustifyContent(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JustifyContent_name[_JustifyContent_index[i]:_JustifyContent_index[i+1]]
}
