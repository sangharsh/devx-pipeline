#!/bin/zsh

# docker compose up -d

fly -t tutorial login -c http://localhost:8080 -u test -p test
fly -t tutorial set-pipeline -p inventory -c ../../inventory/build.yml
# pipelines are paused when first created
fly -t tutorial unpause-pipeline -p inventory
# trigger the job and watch it run to completion
fly -t tutorial trigger-job --job inventory/build --watch