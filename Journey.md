TODOs:

1. Helm wait - Not supported; Unable to wait for postgres to be ready before deploying web app using helm hooks
1. Password is committed in plain-text. Fix it
1. DB connection in golang once closed is not re-created
1. Production deployment
1. CI/CD pipeline
1. Ephemeral environments

## Jul 31

Fixed tests and concourse build pipeline

## Jul 24

Optimized images sizes.
Move from golang:1.22 to golang:1.22-alpine3.20 and alpine:3.20.2 to alpine:3.20 to reuse cached layer

## Jul 19

Docker build needs to be done repeatedly.

1. Setup a registry to store images. Right now, minikube has its own docker and docker desktop - both maintain different images

## Jul 2

Postgres initialization is setup via docker compose and custom SQL mounted as volume

## Jun 19 - Jul 20

Setup initial golang web app
Evaluate different testing options. Yet to close
Evaluate CI options, finalized on Concourse CI
Evaluated secret managers, finalized on Hashicorp Vault
Deploymnet methods: For local, docker compose, k8s yamls, and helm are working at the moment
