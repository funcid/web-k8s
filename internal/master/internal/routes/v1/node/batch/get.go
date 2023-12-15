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
	. "github.com/funcid/web-k8s/internal/master/internal/app"
	. "github.com/funcid/web-k8s/pkg/model/v1/master/node/batch"
	utilv1 "github.com/funcid/web-k8s/pkg/model/v1/util"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Get nodes
// @Description Get a list of all nodes
// @Tags        nodes
// @Accept      json
// @Produce     json
// @Param       data body GetBatchRequest true "the request"
// @Success     200 {object} GetBatchReply
// @Router      /v1/node/batch [get]
func getHandler(app *App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Pagination, etc... ? (really not needed for now)

		return c.JSON(GetBatchReply{
			BaseReply: utilv1.SuccessReply(),
			Nodes:     app.Nodes,
		})
	}
}
