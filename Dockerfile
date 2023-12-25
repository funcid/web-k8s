FROM node:21-slim AS front
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app
COPY ./pkg/front/package.json ./pkg/front/pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
COPY ./pkg/front ./
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm run build

FROM docker.io/golang:1.21-alpine3.18 AS builder

WORKDIR /work

COPY go.mod go.sum ./webk8s/
RUN cd ./webk8s && go mod download

COPY . ./webk8s/
COPY --from=front /app/dist ./webk8s/pkg/front/dist
RUN set -euo pipefail ;\
  mkdir ./bin ;\
	cd ./webk8s ;\
	for cmd in ./cmd/*; do \
		CGO_ENABLED=0 go build -ldflags="-w -s" -v -o ../bin "$cmd" ;\
	done;

FROM docker.io/archlinux:base-devel AS nvml
WORKDIR /work
COPY ./pkg/sysinfo/cmd/nvml ./nvml
RUN set -euxo pipefail ;\
	pacman-key --init ;\
	pacman-key --populate archlinux ;\
	pacman -Syu --noconfirm nvidia-utils patchelf
RUN set -euxo pipefail ;\
	mkdir ./bin/; cd ./bin/ ;\
	gcc ../nvml/main.c -I ../nvml -lnvidia-ml -o ./nvml ;\
	INTERPRETER="$(patchelf --print-interpreter ./nvml)" ;\
	cp "${INTERPRETER}" ./ ;\
	patchelf --set-interpreter "/nvidia/$(basename "${INTERPRETER}")" ./nvml ;\
	patchelf --set-rpath /nvidia ./nvml ;\
	cp /usr/lib/libnvidia-ml.so.1 ./ ;\
	SMI_NEEDED="$(patchelf --print-needed ./nvml)" ;\
	ML_NEEDED="$(patchelf --print-needed ./libnvidia-ml.so.1)" ;\
	for lib in $(echo ${SMI_NEEDED} ${ML_NEEDED} | tr " " "\n" | sort | uniq); do \
		cp "/usr/lib/${lib}" ./ ;\
	done;

FROM scratch AS master
USER 65534:65534
COPY --from=builder /work/bin/master /master
ENTRYPOINT [ "/master" ]

FROM scratch AS worker
USER 65534:65534
COPY --from=builder /work/bin/worker /worker
COPY --from=nvml /work/bin/ /nvidia/
ENV PATH="/nvidia"
ENV WEBK8S_SYSINFO_NVML="/nvidia/nvml"
ENTRYPOINT [ "/worker" ]
