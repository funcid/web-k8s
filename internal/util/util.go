/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package util

func Must[V any](v V, err error) V {
	if err != nil {
		panic(err)
	}
	return v
}

func Apply[T any](val T, fn func(*T)) T {
	fn(&val)
	return val
}

func MapSlice[V any, R any](in []V, fn func(V) R) []R {
	out := make([]R, len(in))
	for idx := range in { // Note: Don't change to idx, val - it will be slower @ 06.12.2023
		out[idx] = fn(in[idx])
	}
	return out
}
