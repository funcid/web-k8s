/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package pods

import (
	. "github.com/funcid/web-k8s/internal/master/internal/app"
	"github.com/funcid/web-k8s/internal/util"
	. "github.com/funcid/web-k8s/pkg/model/v1/master/pods"
	complexv1 "github.com/funcid/web-k8s/pkg/model/v1/resources/complex"
	utilv1 "github.com/funcid/web-k8s/pkg/model/v1/util"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// @Summary     Get pods
// @Description Get list of pods
// @Tags        pods
// @Accept      json
// @Produce     json
// @Param       data body GetPodsRequest true "the request"
// @Success     200 {object} GetPodsReply
// @Router      /v1/pods [get]
func getHandler(app *App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if podList, err := app.K8S.CoreV1().Pods("").List(
			c.UserContext(),
			metav1.ListOptions{},
		); err == nil {
			return c.JSON(GetPodsReply{
				BaseReply: utilv1.SuccessReply(),
				Pods: util.MapSlice(podList.Items, func(pod corev1.Pod) complexv1.Pod {
					return complexv1.Pod{
						Name: pod.Name,
						Images: util.MapSlice(pod.Spec.Containers, func(container corev1.Container) string {
							return container.Image
						}),
						Node: pod.Spec.NodeName,
					}
				}),
			})
		} else {
			const msg = "Failed to get pod list from k8s"
			app.Log.Error(c, msg, zap.Error(err))
			return c.JSON(utilv1.ErrorReply(msg))
		}
	}
}
