package k8s

import (
	"context"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Node struct {
	Hostname string
	Ip       string
}

func GetNodeList() ([]Node, error) {
	var simpleList []Node

	config, err := rest.InClusterConfig()
	if err != nil {
		return simpleList, err
	}

	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		return simpleList, err
	}

	nodes, err := c.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{
		Limit: 500,
	})
	if err != nil {
		return simpleList, err
	}

	for _, item := range nodes.Items {
		node := Node{
			Hostname: item.Name,
		}

		for _, nodeAddress := range item.Status.Addresses {
			if nodeAddress.Type == "InternalIP" {
				node.Ip = nodeAddress.Address
				break
			}
		}
		simpleList = append(simpleList, node)
	}

	return simpleList, nil
}

func GetNodeName() (string, error) {
	name := os.Getenv("NODE_NAME")

	var err error
	if name == "" {
		name, err = os.Hostname()
	}

	return name, err

}
