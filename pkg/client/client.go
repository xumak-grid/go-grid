package client

import (
	"net"
	"os"

	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // Google Cloud Platform authentication
	"k8s.io/client-go/rest"
)

// MustNewKubeClient returns the kubernetes client with cluster config.
func MustNewKubeClient() kubernetes.Interface {
	cfg, err := InClusterConfig()
	if err != nil {
		panic(err)
	}
	return kubernetes.NewForConfigOrDie(cfg)
}

// InClusterConfig gets a client based on the ServiceAccount token provided by the cluster.
func InClusterConfig() (*rest.Config, error) {
	// Work around https://github.com/kubernetes/kubernetes/issues/40973
	// See https://github.com/coreos/etcd-operator/issues/731#issuecomment-283804819
	if len(os.Getenv("KUBERNETES_SERVICE_HOST")) == 0 {
		addrs, err := net.LookupHost("kubernetes.default.svc")
		if err != nil {
			panic(err)
		}
		os.Setenv("KUBERNETES_SERVICE_HOST", addrs[0])
	}
	if len(os.Getenv("KUBERNETES_SERVICE_PORT")) == 0 {
		os.Setenv("KUBERNETES_SERVICE_PORT", "443")
	}
	return rest.InClusterConfig()
}

// BuildKubeConfig allows to build a k8s kubernetes from a path or in cluster config.
func BuildKubeConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return InClusterConfig()
}
