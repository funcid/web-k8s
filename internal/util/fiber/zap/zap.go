/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package zap

import (
	"github.com/funcid/web-k8s/internal/util/fiber/zap/internal/util"
	zaputil "github.com/funcid/web-k8s/internal/util/zap"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Logger struct {
	log zaputil.Logger
}

func (l Logger) Error(c *fiber.Ctx, msg string, fields ...zap.Field) {
	path, method := util.RouteInfo(c.Path(), c.Method())
	l.log.Error(msg, append(fields, path, method)...)
}

func CreateLogger(log zaputil.Logger) Logger {
	return Logger{log}
}
