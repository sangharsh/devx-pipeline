# Inventory service

Written in golang with postgres database

## Install

Use one of the below methods

1. Docker compose
   ```
   docker compose up
   ```
1. [K8s yamls](k8s/README.md)
1. [Helm](helm/README.md)

## Test

If using k8s (yaml/helm) via minikube, get external URL for service:

```
export SERVICE_NAME=inventory-service
minikube service ${SERVICE_NAME} --url
```

Set service URL

```
export INVENTORY_SERVICE_URL=localhost:8080 # Or output of above minikube command
```

Then execute:

```
./test.sh
```
