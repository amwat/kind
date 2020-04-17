package cli

import (
	"os"

	"sigs.k8s.io/kind/pkg/log"
)

func OverrideDefaultName(logger log.Logger, defaultName *string) {
	if name := os.Getenv("KIND_CLUSTER_NAME"); name != "" {
		*defaultName = name
		logger.V(1).Infof("using %q due to KIND_CLUSTER_NAME", name)
	}
}
