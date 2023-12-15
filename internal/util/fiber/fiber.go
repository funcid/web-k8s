/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package fiber

import (
	"embed"
	"net/http"

	"github.com/bytedance/sonic"
	zaputil "github.com/funcid/web-k8s/internal/util/zap"
	"github.com/funcid/web-k8s/internal/version"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/swagger"
)

func CreateApp(name string) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: version.FormatVersion(name),

		DisableStartupMessage: true,

		Prefork: false, // Cannot be used with network "tcp"
		Network: fiber.NetworkTCP,

		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(etag.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	return app
}

func RegisterZap(app *fiber.App, log zaputil.Logger) {
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: log,
	}))
}

func RegisterFS(
	app *fiber.App,
	route string,
	fs embed.FS,
	pathPrefix string,
) {
	app.Use(route, filesystem.New(filesystem.Config{
		Root:       http.FS(fs),
		PathPrefix: pathPrefix,
	}))
}
