// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavutil/mathematics.h>
//#include <stdlib.h>
import "C"
import (
)

type (
	AVRounding C.enum_AVRounding
)

const (
	AV_ROUND_ZERO = AVRounding(C.AV_PICTURE_TYPE_NONE)
	AV_ROUND_INF = AVRounding(C.AV_ROUND_INF)
	AV_ROUND_DOWN = AVRounding(C.AV_ROUND_DOWN)
	AV_ROUND_UP = AVRounding(C.AV_ROUND_UP)
	AV_ROUND_NEAR_INF = AVRounding(C.AV_ROUND_NEAR_INF)
	AV_ROUND_PASS_MINMAX = AVRounding(C.AV_ROUND_PASS_MINMAX)
)

//Return the .
func AvRescaleRnd(a int64, b int64, c int64, rnd AVRounding) int64 {
	return int64(C.av_rescale_rnd((C.int64_t)(a), (C.int64_t)(b), (C.int64_t)(c), (C.enum_AVRounding)(rnd)))
}
