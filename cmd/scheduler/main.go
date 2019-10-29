package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listersv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

const schedulerName = "minimum-pod-scheduler"

type Scheduler struct {
	clientset  *kubernetes.Clientset
	podQueue   chan *v1.Pod
	nodeLister listersv1.NodeLister
}

func NewScheduler(podQueue chan *v1.Pod, quit chan struct{}) Scheduler {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return Scheduler{
		clientset:  clientset,
		podQueue:   podQueue,
		nodeLister: initInformers(clientset, podQueue, quit),
	}
}

func initInformers(clientset *kubernetes.Clientset, podQueue chan *v1.Pod, quit chan struct{}) listersv1.NodeLister {
	factory := informers.NewSharedInformerFactory(clientset, 0)

	nodeInformer := factory.Core().V1().Nodes()
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			node, ok := obj.(*v1.Node)
			if !ok {
				log.Println("this is not a node")
				return
			}
			log.Printf("New Node Added to Store: %s", node.GetName())
		},
	})

	podInformer := factory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod, ok := obj.(*v1.Pod)
			if !ok {
				log.Println("this is not a pod")
				return
			}
			if pod.Spec.NodeName == "" && pod.Spec.SchedulerName == schedulerName {
				podQueue <- pod
			}
		},
	})

	factory.Start(quit)
	return nodeInformer.Lister()
}

func main() {
	fmt.Println("I'm a scheduler!")

	rand.Seed(time.Now().Unix())

	podQueue := make(chan *v1.Pod, 300)
	defer close(podQueue)

	quit := make(chan struct{})
	defer close(quit)

	scheduler := NewScheduler(podQueue, quit)
	scheduler.SchedulePods()
}

func (s *Scheduler) SchedulePods() error {

	for p := range s.podQueue {

		fmt.Println("found a pod to schedule:", p.Namespace, "/", p.Name)

		node, err := s.findFit()
		if err != nil {
			log.Println("cannot find node that fits pod", err.Error())
			continue
		}

		err = s.bindPod(p, node)
		if err != nil {
			log.Println("failed to bind pod", err.Error())
			continue
		}

		message := fmt.Sprintf("Placed pod [%s/%s] on %s\n", p.Namespace, p.Name, node.Name)

		err = s.emitEvent(p, message)
		if err != nil {
			log.Println("failed to emit scheduled event", err.Error())
			continue
		}

		fmt.Println(message)
	}
	return nil
}

func (s *Scheduler) findFit() (*v1.Node, error) {
      minPod := 110
      var bestNode v1.Node

      nodes, err := s.clientset.CoreV1().Nodes().List(metav1.ListOptions{})
      if err != nil {
          panic(err.Error())
      }


      for _, node := range nodes.Items {
        pods, err := s.clientset.CoreV1().Pods("").List(metav1.ListOptions{FieldSelector: "spec.nodeName=" +  node.Name})
        if err != nil {
            panic(err.Error())
        }

        if len(pods.Items) < minPod {
            bestNode = node
            minPod = len(pods.Items)
        }

        fmt.Printf("There are %d pods in the %s node \n", len(pods.Items), node.Name)
      }

      fmt.Printf("==== winner node is %s with %d pods ====\n", bestNode.Name, minPod)
      return &bestNode, nil
}

func (s *Scheduler) bindPod(p *v1.Pod, bestNode *v1.Node) error {
	return s.clientset.CoreV1().Pods(p.Namespace).Bind(&v1.Binding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Name,
			Namespace: p.Namespace,
		},
		Target: v1.ObjectReference{
			APIVersion: "v1",
			Kind:       "Node",
			Name:       bestNode.Name,
		},
	})
}

func (s *Scheduler) emitEvent(p *v1.Pod, message string) error {
	timestamp := time.Now().UTC()
	_, err := s.clientset.CoreV1().Events(p.Namespace).Create(&v1.Event{
		Count:          1,
		Message:        message,
		Reason:         "Scheduled",
		LastTimestamp:  metav1.NewTime(timestamp),
		FirstTimestamp: metav1.NewTime(timestamp),
		Type:           "Normal",
		Source: v1.EventSource{
			Component: schedulerName,
		},
		InvolvedObject: v1.ObjectReference{
			Kind:      "Pod",
			Name:      p.Name,
			Namespace: p.Namespace,
			UID:       p.UID,
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: p.Name + "-",
		},
	})
	if err != nil {
		return err
	}
	return nil
}
