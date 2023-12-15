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
	. "github.com/funcid/web-k8s/internal/util"
	"github.com/funcid/web-k8s/internal/util/zap/internal/encoding"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var developmentConfig = zap.Config{
	Level: zap.NewAtomicLevelAt(zap.DebugLevel),

	Development:       true,
	DisableCaller:     true,
	DisableStacktrace: true,
	Sampling:          nil,

	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},

	Encoding: encoding.Console,

	EncoderConfig: Apply(zap.NewDevelopmentEncoderConfig(), func(cfg *zapcore.EncoderConfig) {
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncodeTime = zapcore.TimeEncoderOfLayout("02.1.2006 15:04 MST")
	}),
}
