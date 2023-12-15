/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

//go:build linux

package util

import (
	"syscall"
)

func GetSysinfo() syscall.Sysinfo_t {
	var info syscall.Sysinfo_t
	if err := syscall.Sysinfo(&info); err != nil {
		panic(err)
	}

	return info
}
