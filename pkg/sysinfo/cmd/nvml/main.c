#include <nvml.h>
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>

#define nvmlCheckError(expr) \
	do { \
		nvmlReturn_t ret = (expr); \
		if (ret != NVML_SUCCESS) { \
			fprintf(stderr, "%s\n", nvmlErrorString(ret)); \
			nvmlShutdown(); \
			exit(1); \
		} \
	} while (0);

#define for_range(val, max) for (size_t (val) = 0; (val) < (max); (val)++)

int main() {
	nvmlCheckError(nvmlInit());

	uint32_t device_count;
	nvmlCheckError(nvmlDeviceGetCount(&device_count));

	for_range (device_index, device_count) {
		nvmlDevice_t device;
		nvmlCheckError(nvmlDeviceGetHandleByIndex(device_index, &device));

		char device_name[NVML_DEVICE_NAME_BUFFER_SIZE];
		uint32_t gpu_cores;

		nvmlCheckError(nvmlDeviceGetName(device, device_name, NVML_DEVICE_NAME_BUFFER_SIZE));
		nvmlCheckError(nvmlDeviceGetNumGpuCores(device, &gpu_cores));

		printf("%u|%s\n", gpu_cores, device_name);
	}

	nvmlCheckError(nvmlShutdown());
	return 0;
}
