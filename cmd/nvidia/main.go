package main

import (
	"flag"

	"github.com/AliyunContainerService/gpushare-device-plugin/pkg/gpu/nvidia"
	log "github.com/golang/glog"
)

var (
	mps         = flag.Bool("mps", false, "Enable or Disable MPS")
	healthCheck = flag.Bool("health-check", false, "Enable or disable Health check")
	memoryUnit  = flag.String("memory-unit", "GiB", "Set memoryUnit of the GPU Memroy, support 'GiB' and 'MiB'")
	gpuRatio    = flag.Float64("gpu-ratio", 0.7, "set available ratio for tf request, because tensorflow control gpu memory is not accurate,")
)

func main() {
	flag.Parse()
	log.V(1).Infoln("Start gpushare device plugin")
	ngm := nvidia.NewSharedGPUManager(*mps, *healthCheck, translatememoryUnits(*memoryUnit), float32(*gpuRatio))
	err := ngm.Run()
	if err != nil {
		log.Fatalf("Failed due to %v", err)
	}
}

func translatememoryUnits(value string) nvidia.MemoryUnit {
	memoryUnit := nvidia.MemoryUnit(value)
	switch memoryUnit {
	case nvidia.MiBPrefix:
	case nvidia.GiBPrefix:
	default:
		log.Warningf("Unsupported memory unit: %s, use memoryUnit Gi as default", value)
		memoryUnit = nvidia.GiBPrefix
	}

	return memoryUnit
}
