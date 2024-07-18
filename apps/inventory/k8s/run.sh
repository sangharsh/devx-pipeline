#!/bin/zsh
# Do not auto run. Few commands need a separate terminal

minikube start
helm install my-postgres bitnami/postgresql -f k8s/postgres-values.yaml
sleep 20
kubectl apply -f k8s/postgres-init-configmap.yaml
kubectl apply -f k8s/postgres-init-job.yaml
eval $(minikube docker-env)
docker build -t inventory:v1 .
kubectl apply -f k8s/app-deployment.yaml
kubectl apply -f k8s/app-service.yaml

# In a different terminal
export SERVICE_NAME=inventory-service
minikube service ${SERVICE_NAME} --url
