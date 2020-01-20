package discovery

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)

// Selector specifies a list of key,value pairs.
type Selector map[string]string

// Service represents the list of properties of services both from spec and last seen status.
type Service struct {
	Name      string           `json:"name"`
	ClusterIP string           `json:"clusterIp"`
	Ports     []v1.ServicePort `json:"ports"`
	PodList   []Pod            `json:"podList"`
	Balancer  string           `json:"balancer"`
}

// Pod represents the basic information about a k8s Pod.
type Pod struct {
	HostIP string            `json:"hostIp"`
	PodIP  string            `json:"podIp"`
	Labels map[string]string `json:"labels"`
	Name   string            `json:"name"`
}

func containsSelector(s Selector, srvc v1.Service, exact bool) bool {
	srvcSelector := srvc.Spec.Selector
	for k, v := range s {
		sv, ok := srvcSelector[k]
		if !ok {
			return false
		}
		if v != sv {
			return false
		}
	}
	if !exact {
		return true
	}
	// If exact is not enabled ensure that both have the same number of selectors
	// otherwise users probably would see services that they don't want
	if len(s) != len(srvcSelector) {
		return false
	}
	return true
}

// Locate locates a service based on input selectors which
// are a list of key,value pairs that needs to match.
func Locate(cli kubernetes.Interface, s Selector, ns string, exact bool) ([]Service, error) {
	services, err := cli.CoreV1().Services(ns).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	servicesList := make([]Service, 0)
	for _, item := range services.Items {
		if containsSelector(s, item, exact) {
			k8Service := Service{
				Name:      item.Name,
				Ports:     item.Spec.Ports,
				ClusterIP: item.Spec.ClusterIP,
			}
			labelsSelector := labels.SelectorFromSet(item.Spec.Selector)
			pods, _ := cli.CoreV1().Pods(ns).List(metav1.ListOptions{
				LabelSelector: labelsSelector.String(),
				FieldSelector: fields.Everything().String(),
			})
			k8Service.PodList = GetPodList(pods)
			if len(item.Status.LoadBalancer.Ingress) > 0 {
				k8Service.Balancer = item.Status.LoadBalancer.Ingress[0].Hostname
				if k8Service.Balancer == "" {
					k8Service.Balancer = item.Status.LoadBalancer.Ingress[0].IP
				}
			}
			servicesList = append(servicesList, k8Service)
		}
	}
	return servicesList, nil
}

// GetPodList maps a k8s PodList to an array of Pods.
func GetPodList(pods *v1.PodList) []Pod {
	list := make([]Pod, 0)
	for _, pod := range pods.Items {
		p := Pod{
			HostIP: pod.Status.HostIP,
			PodIP:  pod.Status.PodIP,
			Labels: pod.Labels,
			Name:   pod.Name,
		}
		list = append(list, p)
	}
	return list
}
