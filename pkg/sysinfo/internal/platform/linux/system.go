/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

//go:build linux

package linux

import (
	"github.com/funcid/web-k8s/pkg/sysinfo/internal/platform/linux/internal/util"
	"github.com/funcid/web-k8s/pkg/sysinfo/internal/types"
)

func FreeMemory() types.Memory {
	return util.GetSysinfo().Freeram
}

func TotalMemory() types.Memory {
	return util.GetSysinfo().Totalram
}

func Uptime() types.Uptime {
	return util.GetSysinfo().Uptime
}
