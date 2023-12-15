/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package v1

import (
	. "github.com/funcid/web-k8s/internal/master/internal/app"
	"github.com/funcid/web-k8s/internal/master/internal/routes/v1/deployment"
	_ "github.com/funcid/web-k8s/internal/master/internal/routes/v1/docs"
	"github.com/funcid/web-k8s/internal/master/internal/routes/v1/node"
	"github.com/funcid/web-k8s/internal/master/internal/routes/v1/pods"
	"github.com/gofiber/fiber/v2"
)

// @title WebK8S API v1
// @version 1.0
// @BasePath /v1
func Register(app *fiber.App, ctx *App) {
	node.Register(app, ctx)
	pods.Register(app, ctx)
	deployment.Register(app, ctx)
}
