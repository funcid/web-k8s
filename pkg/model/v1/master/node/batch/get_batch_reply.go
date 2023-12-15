/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package batch

import (
	model "github.com/funcid/web-k8s/pkg/model/v1"
	"github.com/funcid/web-k8s/pkg/model/v1/resources/complex"
)

type GetBatchReply struct {
	model.BaseReply
	Nodes map[string]*complex.Node `json:"nodes"`
}
