package main

import (
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// K8s config file for the client.
	kubeConfigFile = filepath.Join(os.Getenv("HOME"), ".kube", "config")

	graceSeconds = int64(0)
	deletePolicy = metav1.DeletePropagationForeground
)

func main() {
	client, err := CreateClient(kubeConfigFile)
	if err != nil {
		log.Fatalln("failed to create k8s client:", err.Error())
	}
	log.Infoln("Starting killing process.")

	attempts := 0
	for {
		attempts++
		time.Sleep(time.Second * 2)
		podList, err := client.CoreV1().Pods("kuberhealthy").List(metav1.ListOptions{})
		if err != nil {
			log.Warnln("Unable to list pods in the namespace.")
		}
		log.Debugln("Found", len(podList.Items), "pods on sweep", attempts)

		for _, pod := range podList.Items {
			for k, v := range pod.Labels {
				if k != "app" {
					continue
				}
				if v != "kuberhealthy-check" {
					continue
				}

				if pod.Status.Phase != corev1.PodPending {
					continue
				}

				log.Infoln("Found a pod to delete:", pod.Name)
				err := client.CoreV1().Pods("kuberhealthy").Delete(pod.Name, &metav1.DeleteOptions{
					GracePeriodSeconds: &graceSeconds,
					PropagationPolicy:  &deletePolicy,
				})
				if err != nil {
					log.Warnln("Unable to delete pod", pod.Name+":", err.Error())
				} else {
					log.Infoln(pod.Name, "deleted.")
				}
			}
		}
	}
}
