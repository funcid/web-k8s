/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package master

import (
	. "github.com/funcid/web-k8s/internal/master/internal/app"
	routesv1 "github.com/funcid/web-k8s/internal/master/internal/routes/v1"
	. "github.com/funcid/web-k8s/internal/util"
	zapfiber "github.com/funcid/web-k8s/internal/util/fiber/zap"
	zaputil "github.com/funcid/web-k8s/internal/util/zap"
	"github.com/gofiber/fiber/v2"
	k8s "k8s.io/client-go/kubernetes"
	k8srest "k8s.io/client-go/rest"
)

func Register(
	app *fiber.App,
	logger zaputil.Logger,
	kubeconfig *k8srest.Config,
) {
	ctx := &App{
		Log:   zapfiber.CreateLogger(logger),
		K8S:   Must(k8s.NewForConfig(kubeconfig)),
		Nodes: make(NodesType),
	}

	routesv1.Register(app, ctx)
}
