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
	"bytes"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/bytedance/sonic"
	. "github.com/funcid/web-k8s/internal/util"
	model "github.com/funcid/web-k8s/pkg/model/v1"
	nodev1 "github.com/funcid/web-k8s/pkg/model/v1/master/node"
	"github.com/funcid/web-k8s/pkg/model/v1/resources"
	complexv1 "github.com/funcid/web-k8s/pkg/model/v1/resources/complex"
	"github.com/funcid/web-k8s/pkg/sysinfo"
	fsinfo "github.com/funcid/web-k8s/pkg/sysinfo/fs"
	nvidiainfo "github.com/funcid/web-k8s/pkg/sysinfo/gpu/nvidia"
)

func main() {
	const (
		url = "http://webk8s-master" + nodev1.Path
	)

	for {
		var gpus []resources.GPU

		if nvidiaGpus, err := nvidiainfo.Gpus(); err == nil {
			for _, nvidiaGpu := range nvidiaGpus {
				gpu := resources.GPU(nvidiaGpu)
				gpus = append(gpus, gpu)
			}
		}

		var diskFree uint64 = 0

		if fs, err := fsinfo.FS("/host"); err == nil {
			diskFree = fs.BytesAvailable
		}

		node := complexv1.Node{
			CPU:      resources.CPU(sysinfo.CPU()),
			GPUs:     gpus,
			Memory:   sysinfo.TotalMemory(),
			Uptime:   sysinfo.Uptime(),
			DiskFree: diskFree,
		}

		request := nodev1.UpdateNodeRequest{
			BaseRequest: model.BaseRequest{},
			Name:        os.Getenv("WEBK8S_NODE_NAME"),
			Node:        node,
		}

		res, err := http.Post(
			url,
			"application/json",
			bytes.NewBuffer(Must(sonic.Marshal(&request))),
		)
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		println(string(body))

		time.Sleep(5 * time.Second)
	}
}
