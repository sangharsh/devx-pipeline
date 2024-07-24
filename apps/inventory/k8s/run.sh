#!/bin/zsh
# Do not auto run. Few commands need a separate terminal

minikube start

# TODO: Consider replacing helm with k8s yamls for postgres
# May need pvc, deployment, secret, service
helm install inventory-db bitnami/postgresql -f k8s/postgres-values.yaml --version 15.5.16
sleep 20
kubectl create configmap db-init-sql --from-file=db-init.sql=./db-init.sql
kubectl apply -f k8s/postgres-init-job.yaml
# Point docker client to dockerd inside minikube
eval $(minikube -p minikube docker-env)
docker build -t inventory:latest .
kubectl apply -f k8s/app-deployment.yaml
kubectl apply -f k8s/app-service.yaml

# In a different terminal
export SERVICE_NAME=inventory-service
minikube service ${SERVICE_NAME} --url
