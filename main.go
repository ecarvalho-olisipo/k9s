package main

import (
	"flag"
	"os/exec"

	"github.com/derailed/k9s/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"
)

func init() {
	klog.InitFlags(nil)

	if err := flag.Set("logtostderr", "false"); err != nil {
		panic(err)
	}
	if err := flag.Set("alsologtostderr", "false"); err != nil {
		panic(err)
	}
	if err := flag.Set("stderrthreshold", "fatal"); err != nil {
		panic(err)
	}
	if err := flag.Set("v", "0"); err != nil {
		panic(err)
	}
}

func main() {
	shell := exec.Command("osprey", "user", "login", "-g", "dev")
	shell.Run()
	cmd.Execute()
}
