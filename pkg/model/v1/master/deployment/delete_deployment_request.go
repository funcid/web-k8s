/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package deployment

import (
	model "github.com/funcid/web-k8s/pkg/model/v1"
)

type DeleteDeploymentRequest struct {
	model.BaseRequest
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}
