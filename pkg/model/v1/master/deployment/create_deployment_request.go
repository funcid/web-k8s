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
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type BaseMetadata struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

type CreateDeploymentRequest struct {
	model.BaseRequest
	BaseMetadata
	Replicas    int32                         `json:"replicas"`
	Strategy    appsv1.DeploymentStrategyType `json:"strategy"`
	MatchLabels map[string]string             `json:"matchLabels"`
	Pod         PodCreateRequest              `json:"pod"`
}

type PodCreateRequest struct {
	BaseMetadata
	RestartPolicy corev1.RestartPolicy        `json:"restartPolicy"`
	Containers    []PodCreateRequestContainer `json:"containers"`
}

type PodCreateRequestContainer struct {
	Name    string                           `json:"name"`
	Image   string                           `json:"image"`
	Command []string                         `json:"command"`
	Args    []string                         `json:"args"`
	Ports   []PodCreateRequestContainersPort `json:"ports"`
}

type PodCreateRequestContainersPort struct {
	ContainerPort int32 `json:"containerPort"`
}
