/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package sysinfo

import (
	"github.com/funcid/web-k8s/pkg/sysinfo/internal/provider"
)

func FreeMemory() MemoryType {
	return provider.FreeMemory()
}

func TotalMemory() MemoryType {
	return provider.TotalMemory()
}

func Uptime() UptimeType {
	return provider.Uptime()
}
