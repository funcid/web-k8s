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
	. "github.com/funcid/web-k8s/internal/master/internal/app"
	. "github.com/funcid/web-k8s/pkg/model/v1/master/deployment"
	utilv1 "github.com/funcid/web-k8s/pkg/model/v1/util"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// @Summary     Delete deployment
// @Description Deletes k8s deployment
// @Tags        deployments
// @Accept      json
// @Produce     json
// @Param       data body DeleteDeploymentRequest true "the request"
// @Success     200 {object} DeleteDeploymentReply
// @Router      /v1/deployment [delete]
func deleteHandler(app *App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := DeleteDeploymentRequest{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if err := app.K8S.AppsV1().Deployments(payload.Namespace).Delete(
			c.UserContext(),
			payload.Name,
			metav1.DeleteOptions{},
		); err != nil {
			app.Log.Error(c, "Failed to delete deployment", zap.Error(err))
			return c.JSON(utilv1.ErrorReply("failed to delete"))
		} else {
			return c.JSON(utilv1.SuccessReply())
		}
	}
}
