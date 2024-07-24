#!/bin/zsh
# Do not auto run. Few commands need a separate terminal

minikube start

# TODO: Consider replacing helm with k8s yamls for postgres
# May need pvc, deployment, secret, service
helm install inventory bitnami/postgresql -f k8s/postgres-values.yaml --version 15.5.16
sleep 20
kubectl create configmap inventory-init-data --from-file=db-init.sql=./db-init.sql
kubectl apply -f k8s/init-job.yaml
# Point docker client to dockerd inside minikube
eval $(minikube -p minikube docker-env)
docker build -t inventory:latest .
kubectl apply -f k8s/app-deployment.yaml
kubectl apply -f k8s/app-service.yaml

# In a different terminal
export SERVICE_NAME=inventory-service
minikube service ${SERVICE_NAME} --url

# Uninstall and clean up
kubectl delete deploy inventory
kubectl delete service inventory
kubectl delete job inventory-init-job
kubectl delete configmap inventory-init-data
helm uninstall inventory
kubectl delete pvc data-inventory-postgresql-0
