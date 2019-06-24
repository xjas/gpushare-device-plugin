package nvidia

import (
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

// MemoryUnit describes GPU Memory, now only supports Gi, Mi
type MemoryUnit string

const (
	resourceName  = "transwarp.io/gpu-mem"
	resourceCount = "transwarp.io/gpu-count"
	serverSock    = pluginapi.DevicePluginPath + "transwarp-gpu-share.sock"

	//0.7 is because tensorflow control gpu memory is not accurate, it is recommended to multiply by 0.7 to ensure that the upper limit is not exceeded.
	// AvailableNvidiaMemoryRatio = 0.7

	//used by tensorflow to control memory usage
	envGPUMemoryFraction = "GPU_MEMORY_FRACTION"

	OptimisticLockErrorMsg = "the object has been modified; please apply your changes to the latest version and try again"

	allHealthChecks             = "xids"
	containerTypeLabelKey       = "io.kubernetes.docker.type"
	containerTypeLabelSandbox   = "podsandbox"
	containerTypeLabelContainer = "container"
	containerLogPathLabelKey    = "io.kubernetes.container.logpath"
	sandboxIDLabelKey           = "io.kubernetes.sandbox.id"

	envNVGPU               = "NVIDIA_VISIBLE_DEVICES"
	envCUDAGPU             = "CUDA_VISIBLE_DEVICES"
	EnvResourceIndex       = "TRANSWARP_IO_GPU_MEM_IDX"
	EnvResourceByPod       = "TRANSWARP_IO_GPU_MEM_POD"
	EnvResourceByContainer = "TRANSWARP_IO_GPU_MEM_CONTAINER"
	EnvResourceByDev       = "TRANSWARP_IO_GPU_MEM_DEV"
	EnvAssignedFlag        = "TRANSWARP_IO_GPU_MEM_ASSIGNED"
	EnvResourceAssumeTime  = "TRANSWARP_IO_GPU_MEM_ASSUME_TIME"
	EnvResourceAssignTime  = "TRANSWARP_IO_GPU_MEM_ASSIGN_TIME"

	GiBPrefix = MemoryUnit("GiB")
	MiBPrefix = MemoryUnit("MiB")
)
