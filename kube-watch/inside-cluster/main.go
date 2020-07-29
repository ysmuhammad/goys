package main

import (
	"log"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		log.Panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err.Error())
	}

	factory := informers.NewSharedInformerFactory(clientset, 0)

	// informer := factory.Core().V1().Pods().Informer()
	informer := factory.Core().V1().Events().Informer()
	stopper := make(chan struct{})
	defer close(stopper)

	defer runtime.HandleCrash()

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// "k8s.io/apimachinery/pkg/apis/meta/v1" provides an Object
			// interface that allows us to get metadata easily
			// mObj := obj.(v1.Object)
			mObj := obj.(*corev1.Event)

			log.Printf("============\n")
			log.Printf("Event Message: %s", mObj.Message)
			log.Printf("Reason: %s", mObj.Reason)
		},
	})

	informer.Run(stopper)

}
