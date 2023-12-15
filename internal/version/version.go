/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package version

import (
	"strings"
)

var (
	version = "0.0.0"
)

func Version() string { return version }

func FormatVersion(str string) string {
	return strings.ReplaceAll(str, "{version}", "v"+version)
}
