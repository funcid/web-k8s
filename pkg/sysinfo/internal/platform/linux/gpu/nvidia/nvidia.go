/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

//go:build linux

package nvidia

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	. "github.com/funcid/web-k8s/internal/util"
	osutil "github.com/funcid/web-k8s/internal/util/os"
	"github.com/funcid/web-k8s/pkg/sysinfo/internal/platform/linux/gpu/nvidia/internal/errors"
	"github.com/funcid/web-k8s/pkg/sysinfo/internal/types"
)

var (
	nvml = osutil.GetEnv("WEBK8S_SYSINFO_NVML", "")
)

func Gpus() ([]types.GPU, error) {
	if nvml == "" {
		return nil, errors.ErrorNvmlBinaryIsNotSet
	}

	cmd := exec.Command(nvml)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	if cmd.ProcessState.ExitCode() == 0 {
		gpus := strings.Split(stdout.String(), "\n")
		return MapSlice(gpus, func(line string) types.GPU {
			split := strings.Split(line, "|")
			return types.GPU{
				Model: split[1],
				Cores: int32(Must(strconv.ParseInt(split[0], 10, 32))),
			}
		}), nil
	} else {
		return nil, fmt.Errorf("%s", stderr.String())
	}
}
