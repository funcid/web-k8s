/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package util

import (
	. "github.com/funcid/web-k8s/internal/util"
	. "github.com/funcid/web-k8s/pkg/model/v1/master/deployment"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreatePod(payload *PodCreateRequest) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      payload.Name,
			Labels:    payload.Labels,
			Namespace: payload.Namespace,
		},
		Spec: v1.PodSpec{
			RestartPolicy: payload.RestartPolicy,
			Containers: MapSlice(payload.Containers, func(container PodCreateRequestContainer) v1.Container {
				return v1.Container{
					Name:    container.Name,
					Image:   container.Image,
					Command: container.Command,
					Args:    container.Args,
					Ports: MapSlice(container.Ports, func(port PodCreateRequestContainersPort) v1.ContainerPort {
						return v1.ContainerPort{ContainerPort: port.ContainerPort}
					}),
				}
			}),
		},
	}
}
