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
	model "github.com/funcid/web-k8s/pkg/model/v1"
)

func SuccessReply() model.BaseReply {
	// TODO: Add support of custom replies (not needed for now)
	return model.BaseReply{Success: true}
}

func ErrorReply(msg string) model.BaseReply {
	return model.BaseReply{
		Success: false,
		Error:   msg,
	}
}
