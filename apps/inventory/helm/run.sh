#!/bin/zsh

# To be run manually
if [[ -z RELEASE_NAME ]]; then
    echo "Setting RELEASE_NAME to r1"
    RELEASE_NAME=r1
fi

helm install ${RELEASE_NAME} helm/ --dependency-update

# Clean up
helm uninstall ${RELEASE_NAME}
# Delete PVC held by postgres
kubectl delete pvc data-${RELEASE_NAME}-postgresql-0
