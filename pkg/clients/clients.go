package clients

import (
	"os"

	clientconfigv1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// GetKubeClientSetOrDie Add a native non-caching client for advanced CRUD operations
func GetKubeClientSetOrDie(cfg *rest.Config) kubernetes.Clientset {

	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		os.Exit(1)
	}
	return *clientSet
}

// GetConfigClientOrDie Add a configv1 client to the reconciler
func GetConfigClientOrDie(cfg *rest.Config) clientconfigv1.ConfigV1Client {

	client, err := clientconfigv1.NewForConfig(cfg)
	if err != nil {
		os.Exit(1)
	}
	return *client
}
