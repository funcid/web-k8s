/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

//go:build linux

package errors

import (
	"fmt"

	"github.com/funcid/web-k8s/pkg/sysinfo/internal/platform/linux/gpu/nvidia/internal/common"
)

const (
	NvmlBinaryIsNotSet = "the nvml binary is not set in the environment variable " + common.NvmlBinaryEnvVar
)

var (
	ErrorNvmlBinaryIsNotSet = fmt.Errorf(NvmlBinaryIsNotSet)
)
