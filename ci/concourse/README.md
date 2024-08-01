# Concourse for CI

## Run
From `concourse` directory run following command to start concourse with postgres
```
docker compose up -d
```

## Install Fly CLI
https://concourse-ci.org/quick-start.html#install-fly

## Setup pipeline
> TODO: Use long-form name for command flags
```
fly -t tutorial login -c http://localhost:8080 -u test -p test
fly -t tutorial set-pipeline -p inventory -c ../../apps/inventory/build.yml

# pipelines are paused when first created
fly -t tutorial unpause-pipeline -p inventory

# trigger the job and watch it run to completion
fly -t tutorial trigger-job --job inventory/build --watch
```
