// Code generated by "stringer -type=PatsType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Easy-0]
	_ = x[Hard-1]
	_ = x[Impossible-2]
	_ = x[Lines2-3]
	_ = x[PatsTypeN-4]
}

const _PatsType_name = "EasyHardImpossibleLines2PatsTypeN"

var _PatsType_index = [...]uint8{0, 4, 8, 18, 24, 33}

func (i PatsType) String() string {
	if i < 0 || i >= PatsType(len(_PatsType_index)-1) {
		return "PatsType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PatsType_name[_PatsType_index[i]:_PatsType_index[i+1]]
}