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
	"os"
	"regexp"

	"github.com/funcid/web-k8s/pkg/sysinfo/internal/types"
)

var (
	modelNameRegex = regexp.MustCompile("(?m)^model name[ \t]+: (.+)$")
)

func CPU() types.CPU {
	cpuinfoBytes, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		panic("could not read /proc/cpuinfo")
	}

	cpuinfo := string(cpuinfoBytes)

	model := modelNameRegex.FindStringSubmatch(cpuinfo)
	if len(model) < 2 {
		panic("cpu not found (/proc/cpuinfo)")
	}
	cores := len(modelNameRegex.FindAllStringIndex(cpuinfo, -1))

	return types.CPU{Model: model[1] /* first group */, Cores: uint32(cores)}
}
