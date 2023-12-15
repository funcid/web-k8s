/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

//go:build !linux

package provider

import (
	"github.com/funcid/web-k8s/pkg/sysinfo/internal/types"
)

func FreeMemory() types.Memory {
	panic("Getting free memory is not supported on the current OS")
}

func TotalMemory() types.Memory {
	panic("Getting memory is not supported on the current OS")
}

func Uptime() types.Uptime {
	panic("Getting uptime is not supported on the current OS")
}
