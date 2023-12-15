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
	zapenv "github.com/funcid/web-k8s/internal/util/zap/internal/env"
	"go.uber.org/zap"
)

type (
	Logger = *zap.Logger
)

func CreateLogger() (Logger, error) {
	return zapenv.Config.Build()
}
