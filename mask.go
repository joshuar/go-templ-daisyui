// Copyright 2024 Joshua Rich <joshua.rich@gmail.com>.
// SPDX-License-Identifier: 	AGPL-3.0-or-later

//go:generate stringer -type=Mask -linecomment -output mask_generated.go
package components

const (
	MaskNone              Mask = iota //
	MaskSquircle                      // mask-squircle
	MaskHeart                         // mask-heart
	MaskHexagon                       // mask-hexagon
	MaskHexagonAlt                    // mask-hexagon-2
	MaskDecagon                       // mask-decagon
	MaskPentagon                      // mask-pentagon
	MaskDiamond                       // mask-diamond
	MaskSquare                        // mask-square
	MaskCircle                        // mask-circle
	MaskParallelogram                 // mask-parallelogram
	MaskParallelogramAlt1             // mask-parallelogram-2
	MaskParallelogramAlt2             // mask-parallelogram-3
	MaskParallelogramAlt3             // mask-parallelogram-4
	MaskStar                          // mask-star
	MaskStarAlt                       // mask-star-2
	MaskTriangle                      // mask-triangle
	MaskTriangleAlt                   // mask-triangle-2
	MaskTriangleAlt2                  // mask-triangle-3
	MaskTriangleAlt3                  // mask-triangle-4
	MaskHalfLeft                      // mask-half-1
	MaskHalfRight                     // mask-half-2
)

type Mask int
