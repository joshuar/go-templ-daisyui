// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

package size

const (
	MinFixedSize FixedSize = iota
	Size0
	SizePx
	Size0_5
	Size1
	Size1_5
	Size2
	Size2_5
	Size3
	Size3_5
	Size4
	Size5
	Size6
	Size7
	Size8
	Size9
	Size10
	Size11
	Size12
	Size14
	Size16
	Size20
	Size24
	Size28
	Size32
	Size36
	Size40
	Size44
	Size48
	Size52
	Size56
	Size64
	Size72
	Size80
	Size96
	SizeAuto
	SizeOneHalf
	SizeOneThird
	SizeTwoThirds
	SizeOneQuarter
	SizeTwoQuarters
	SizeThreeQuarters
	SizeOneFifth
	SizeTwoFifths
	SizeThreeFifths
	SizeFourFifths
	SizeOneSixth
	SizeTwoSixths
	SizeThreeSixths
	SizeFourSixths
	SizeFiveSixths
	SizeMin
	SizeMax
	MaxFixedSize
)

// FixedSize represents a fixed, percentage or one of auto/min/max sizes.
//
// https://tailwindcss.com/docs/size
type FixedSize int

// Valid returns true if the size is a valid response size.
func (size FixedSize) Valid() bool {
	return size > MinFixedSize && size < MaxFixedSize
}
