build:
	docker build -t jonnydawg/k8s-pending-killer:v0.0.7 .

push:
	docker push jonnydawg/k8s-pending-killer:v0.0.7