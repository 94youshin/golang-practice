package main

import (
	"github.com/usmhk/gopractice/etcd_kubernetes"
	"github.com/usmhk/gopractice/https"
)

func main() {
	https.Https()

	etcd_kubernetes.Etcd_Kubernetes()
}
