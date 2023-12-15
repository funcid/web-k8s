/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package node

import (
	. "github.com/funcid/web-k8s/internal/master/internal/app"
	. "github.com/funcid/web-k8s/pkg/model/v1/master/node"
	utilv1 "github.com/funcid/web-k8s/pkg/model/v1/util"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// @Summary     Update node
// @Description Update (or create) a node
// @Tags        nodes
// @Accept      json
// @Produce     json
// @Param       data body UpdateNodeRequest true "the request"
// @Success     200 {object} UpdateNodeReply
// @Router      /v1/node [post]
func updateHandler(app *App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := UpdateNodeRequest{}

		if err := c.BodyParser(&payload); err != nil {
			app.Log.Error(c, "Failed to parse body", zap.Error(err))
			return c.JSON(utilv1.ErrorReply("invalid request"))
		}

		app.Nodes[payload.Name] = &payload.Node

		return c.JSON(utilv1.SuccessReply())
	}
}
