/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package app

import (
	zapfiber "github.com/funcid/web-k8s/internal/util/fiber/zap"
	complexv1 "github.com/funcid/web-k8s/pkg/model/v1/resources/complex"
	k8s "k8s.io/client-go/kubernetes"
)

type (
	NodesType = map[string]*complexv1.Node
)

type App struct {
	Log   zapfiber.Logger
	K8S   *k8s.Clientset
	Nodes NodesType
}
