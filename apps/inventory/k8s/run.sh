#!/bin/zsh
# Do not auto run. Few commands need a separate terminal

minikube start
helm install my-postgres bitnami/postgresql -f k8s/postgres-values.yaml
sleep 20
kubectl create configmap db-init-sql --from-file=db-init.sql=./db-init.sql
kubectl apply -f k8s/postgres-init-job.yaml
eval $(minikube docker-env)
docker build -t inventory:latest .
kubectl apply -f k8s/app-deployment.yaml
kubectl apply -f k8s/app-service.yaml

# In a different terminal
export SERVICE_NAME=inventory-service
minikube service ${SERVICE_NAME} --url
