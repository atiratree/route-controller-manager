package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/openshift/library-go/pkg/serviceability"
	"github.com/openshift/route-controller-manager/pkg/version"
	cache2 "k8s.io/apimachinery/pkg/util/cache"
	"k8s.io/apiserver/pkg/authentication/token/cache"
	genericapiserver "k8s.io/apiserver/pkg/server"
	clock2 "k8s.io/utils/clock"
)

func main() {
	_ = genericapiserver.SetupSignalContext()

	defer serviceability.BehaviorOnPanic(os.Getenv("OPENSHIFT_ON_PANIC"), version.Get())()
	defer serviceability.Profile(os.Getenv("OPENSHIFT_PROFILE")).Stop()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	for {
		fmt.Printf("creating RealClock\n")
		clock := clock2.RealClock{}
		fmt.Printf("created RealClock %#v\n")

		fmt.Printf("creating NewExpiringWithClock\n")
		expiring := cache2.NewExpiringWithClock(clock)
		fmt.Printf("created NewExpiringWithClock %#v\n", expiring)

		fmt.Printf("creating cached token authenticator\n")
		cachedAuthenticator := cache.New(nil, false, 5*time.Minute, 5*time.Minute)
		fmt.Printf("created cached token authenticator %#v\n", cachedAuthenticator)
	}
}
