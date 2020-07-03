package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 	var kubeconfig *string
	// 	if home := homeDir(); home != "" {
	// 		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// 	} else {
	// 		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// 	}
	// 	flag.Parse()

	// 	// use the current context in kubeconfig
	// 	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	// create the clientset
	// 	clientset, err := kubernetes.NewForConfig(config)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	for {
	// 		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// 		if err != nil {
	// 			panic(err.Error())
	// 		}
	// 		fmt.Printf("There are %v pods in the cluster\n", len(pods.Items))

	// 		// Examples for error handling:
	// 		// - Use helper functions like e.g. errors.IsNotFound()
	// 		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// 		namespace := "kube-system"
	// 		pod := "coredns-6955765f44-9bct4"
	// 		//_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	// 		var eve string
	// 		_, err = clientset.CoreV1().Events("kube-system").Get(context.TODO(), "",metav1.GetOptions{})
	// 		if errors.IsNotFound(err) {
	// 			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// 		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// 				pod, namespace, statusError.ErrStatus.Message)
	// 		} else if err != nil {
	// 			panic(err.Error())
	// 		} else {
	// 			fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// 		}
	// 		fmt.Println(eve)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }

	// func homeDir() string {
	// 	if h := os.Getenv("HOME"); h != "" {
	// 		return h
	// 	}
	// 	return os.Getenv("USERPROFILE") // windows
	// }

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
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
			log.Printf("Event Message: %s", mObj.Message)
			log.Printf("Reason: %s", mObj.Reason)
		},
	})

	informer.Run(stopper)

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
