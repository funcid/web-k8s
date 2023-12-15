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

func CPU() types.CPU {
	panic("Retrieving processor information is not supported in the current OS")
}
