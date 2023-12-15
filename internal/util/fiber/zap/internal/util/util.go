/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package util

import (
	"go.uber.org/zap"
)

func RouteInfo(path, method string) (zap.Field, zap.Field) {
	return zap.String("path", path), zap.String("method", method)
}
