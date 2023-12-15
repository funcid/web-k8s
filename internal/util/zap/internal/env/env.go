/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package env

import (
	"fmt"

	osutil "github.com/funcid/web-k8s/internal/util/os"
	"go.uber.org/zap"
)

const (
	Production  = "production"
	Development = "development"
)

var (
	Config = selectEnvConfig()
)

func selectEnvConfig() *zap.Config {
	env := osutil.GetEnv("WEBK8S_ENV", Production)
	switch env {
	case Production:
		return &productionConfig
	case Development:
		return &developmentConfig
	default:
		panic(fmt.Sprintf("WEBK8S_ENV = %s (unknown environment)", env))
	}
}
