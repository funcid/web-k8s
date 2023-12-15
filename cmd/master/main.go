/*
 * Copyright (c) 2023 Kamillaova. All rights reserved.
 *
 * This software is licensed under the WebK8S End-User License Agreement (EULA).
 * A copy of the EULA is included in the repository's LICENSE file.
 *
 * For non-commercial usage only. Modifications must be published under the same license.
 */

package main

import (
	"fmt"

	"github.com/funcid/web-k8s/internal/master"
	. "github.com/funcid/web-k8s/internal/util"
	fiberutil "github.com/funcid/web-k8s/internal/util/fiber"
	zaputil "github.com/funcid/web-k8s/internal/util/zap"
	"github.com/funcid/web-k8s/pkg/front"
	"go.uber.org/zap"
	k8srest "k8s.io/client-go/rest"
)

const (
	port uint16 = 8080
)

func main() {
	log := Must(zaputil.CreateLogger())
	kubeconfig := Must(k8srest.InClusterConfig())

	app := fiberutil.CreateApp("WebK8S {version} @ master")

	fiberutil.RegisterZap(app, log)
	master.Register(app, log, kubeconfig)
	fiberutil.RegisterFS(app, "/", front.Front, "dist")

	log.Info("Starting the server...", zap.Uint16("port", port))
	log.Fatal("Failed to start the server", zap.Error(app.Listen(fmt.Sprintf(":%d", port))))
}
