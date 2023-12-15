/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package complex

import (
	"github.com/funcid/web-k8s/pkg/model/v1/resources"
)

type Node struct {
	CPU    resources.CPU    `json:"cpu"`
	Memory resources.Memory `json:"memory"`

	GPUs []resources.GPU `json:"gpus"`

	Uptime resources.Uptime `json:"uptime"`

	DiskFree resources.DiskSpace `json:"disk_free"`
}
