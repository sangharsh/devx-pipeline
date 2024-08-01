# Install app using Helm

## Install helm

https://helm.sh/docs/intro/install/

## Install minikube

Instructions [here](../k8s/README.md#install-minikube)

## Set release name

```
export RELEASE_NAME=r1 # Replace with your chosen release name
```

## Install charts

```
helm install ${RELEASE_NAME} helm/ --dependency-update
```

## Uninstall and clean up

```
helm uninstall ${RELEASE_NAME}
```

Data attached with postgres is not destroyed during `helm uninstall`. Concerned PVC needs to be deleted explicitly.

```
kubectl delete pvc data-${RELEASE_NAME}-postgresql-0
```
