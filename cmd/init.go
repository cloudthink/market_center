package main

import (
	"context"
	"fmt"
	"github.com/goex-top/market_center/api"
	"github.com/goex-top/market_center/config"
	"github.com/goex-top/market_center/data"
	"os"
	"runtime"
	"strconv"
)

// AdjustGoMaxProcs adjusts the maximum processes that the CPU can handle.
func AdjustGoMaxProcs() {
	fmt.Println("Adjusting runtime performance..")
	maxProcsEnv := os.Getenv("GOMAXPROCS")
	maxProcs := runtime.NumCPU()
	fmt.Println("Number of CPU's detected:", maxProcs)

	if maxProcsEnv != "" {
		fmt.Println("GOMAXPROCS env =", maxProcsEnv)
		env, err := strconv.Atoi(maxProcsEnv)
		if err != nil {
			fmt.Println("Unable to convert GOMAXPROCS to int, using", maxProcs)
		} else {
			maxProcs = env
		}
	}
	if i := runtime.GOMAXPROCS(maxProcs); i != maxProcs {
		fmt.Println("Go Max Procs were not set correctly.")
	}
	fmt.Println("Set GOMAXPROCS to:", maxProcs)
}

func init() {
	AdjustGoMaxProcs()
	Ctx, Cancel = context.WithCancel(context.Background())
	Cfg = config.NewConfig()
	Data = data.NewData()
	Api = api.NewApi(Ctx, Cfg, Data)
	Api.EnableDebug()
}
