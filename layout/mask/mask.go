// Copyright 2025 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

// Package mask represents a DaisyUI Mask.
//
// https://daisyui.com/components/mask/
package mask

const (
	MaskMin               Mask = iota
	MaskSquircle               // mask-squircle
	MaskHeart                  // mask-heart
	MaskHexagon                // mask-hexagon
	MaskHexagonAlt             // mask-hexagon-2
	MaskDecagon                // mask-decagon
	MaskPentagon               // mask-pentagon
	MaskDiamond                // mask-diamond
	MaskSquare                 // mask-square
	MaskCircle                 // mask-circle
	MaskParallelogram          // mask-parallelogram
	MaskParallelogramAlt1      // mask-parallelogram-2
	MaskParallelogramAlt2      // mask-parallelogram-3
	MaskParallelogramAlt3      // mask-parallelogram-4
	MaskStar                   // mask-star
	MaskStarAlt                // mask-star-2
	MaskTriangle               // mask-triangle
	MaskTriangleAlt1           // mask-triangle-2
	MaskTriangleAlt2           // mask-triangle-3
	MaskTriangleAlt3           // mask-triangle-4
	MaskHalfLeft               // mask-half-1
	MaskHalfRight              // mask-half-2
	MaskMax
)

type Mask int

func (m Mask) Valid() bool {
	return m > MaskMin && m < MaskMax
}
