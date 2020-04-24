## k8s-pending-killer

This is a simple go program that kills pods that are in status `Pending` to help
debug [issue](https://github.com/Comcast/kuberhealthy/issues/436).

Please be careful when using this!

For safe usage, make sure you are using a local k8s cluster -- _NOT PROD_.

#### Usage

Apply the deployment file found in this directory: 

> `kubectl apply -f deployment.yaml` 

_*Don't forget to clean up after you are done!*_

> `kubectl delete -f deployment.yaml`