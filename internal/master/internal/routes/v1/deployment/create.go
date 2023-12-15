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
	"github.com/funcid/web-k8s/internal/master/internal/routes/v1/deployment/internal/util"
	. "github.com/funcid/web-k8s/pkg/model/v1/master/deployment"
	utilv1 "github.com/funcid/web-k8s/pkg/model/v1/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// @Summary     Create deployment
// @Description Creates k8s deployment
// @Tags        deployments
// @Accept      json
// @Produce     json
// @Param       data body CreateDeploymentRequest true "the request"
// @Success     200 {object} CreateDeploymentReply
// @Router      /v1/deployment [post]
func createHandler(app *App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := CreateDeploymentRequest{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		pod := util.CreatePod(&payload.Pod)

		deployment := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      payload.Name,
				Labels:    payload.Labels,
				Namespace: payload.Namespace,
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &payload.Replicas,
				Selector: &metav1.LabelSelector{
					MatchLabels: payload.MatchLabels,
				},
				Template: v1.PodTemplateSpec{
					ObjectMeta: pod.ObjectMeta,
					Spec:       pod.Spec,
				},
				Strategy: appsv1.DeploymentStrategy{Type: payload.Strategy},
			},
		}
		if created, err := app.K8S.AppsV1().Deployments(payload.Namespace).Create(
			c.UserContext(),
			deployment,
			metav1.CreateOptions{},
		); err != nil {
			app.Log.Error(c, "Failed to create new deployment", zap.Error(err))
			return c.JSON(utilv1.ErrorReply("failed to create"))
		} else {
			return c.JSON(CreateDeploymentReply{
				BaseReply: utilv1.SuccessReply(),
				UUID:      uuid.MustParse(string(created.UID)),
			})
		}
	}
}
