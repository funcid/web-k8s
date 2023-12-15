/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package fs

import (
	"github.com/funcid/web-k8s/pkg/sysinfo"
	provider "github.com/funcid/web-k8s/pkg/sysinfo/internal/provider/fs"
)

func FS(path string) (sysinfo.FSType, error) {
	return provider.FS(path)
}
