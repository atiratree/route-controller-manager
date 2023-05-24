package main

import (
	"context"
	"os"
	"runtime"

	"github.com/openshift/library-go/pkg/serviceability"
	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"

	route_controller_manager "github.com/openshift/route-controller-manager/pkg/cmd/route-controller-manager"
	"github.com/openshift/route-controller-manager/pkg/version"
)

func main() {
	ctx := genericapiserver.SetupSignalContext()

	defer serviceability.BehaviorOnPanic(os.Getenv("OPENSHIFT_ON_PANIC"), version.Get())()
	defer serviceability.Profile(os.Getenv("OPENSHIFT_PROFILE")).Stop()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	command := NewRouteControllerManagerCommand(ctx)
	code := cli.Run(command)
	os.Exit(code)
}

func NewRouteControllerManagerCommand(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "route-controller-manager",
		Short: "Command for additional management of ingress and Route resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}
	start := route_controller_manager.NewRouteControllerManagerCommand("start", os.Stdout, os.Stderr, ctx)
	cmd.AddCommand(start)

	return cmd
}
