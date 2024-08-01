# Deploy service on Kubernetes with yaml files

## Install minikube
Follow minikube installtion instructions at https://minikube.sigs.k8s.io/docs/start/. Post installation, start minikube with:
```
minikube start
```

## Deploy Postgres and web app
> TODO: Consider replacing helm with k8s yamls for postgres
May need pvc, deployment, secret, service

Deploy postgres and seed data
```
helm install inventory bitnami/postgresql -f k8s/postgres-values.yaml --version 15.5.16
sleep 20
kubectl create configmap inventory-init-data --from-file=db-init.sql=./db-init.sql
kubectl apply -f k8s/init-job.yaml
```

Build and deploy web app
```
# Point docker client to dockerd inside minikube
eval $(minikube -p minikube docker-env)

docker build -t inventory:latest .
kubectl apply -f k8s/app-deployment.yaml
kubectl apply -f k8s/app-service.yaml
```

In a different terminal
```
export SERVICE_NAME=inventory-service
minikube service ${SERVICE_NAME} --url
```

# Uninstall and clean up
```
kubectl delete deploy inventory
kubectl delete service inventory
kubectl delete job inventory-init-job
kubectl delete configmap inventory-init-data
helm uninstall inventory
kubectl delete pvc data-inventory-postgresql-0
```
