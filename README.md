# Developer experience setup

# Business requirement

Build a system to take simple e-commerce orders and manage inventory through a browser-based interface

# Tech design

Develop a set of services which expose REST API with JSON request/response. Use golang for one and Java (springboot) for another. Both will have their own databases. (Try to) Add a cache (redis), a message broker

Use React/Vue for a frontend

# Project setup

Tools being used should be free, preferably open-source or atleast open-core. Being a personal project, it is imperative to not have a recurring cost to host, run etc.
Most of the things should be source code driven (gitops?). Avoid vendor lock-in. Keep things reproducible.

Use a set of popular tools to build.

Continually evaluate each step for developer experience.

## Project requirement

- TODO: Doesn't exist currently. Evaluate need. Options: Redmine, Propritary: Github, Confluence, Jira

## Repository

- Host on github.com but avoid Github specific features as much as possible
- Trying mono repo approach

## Development environment

- VSCode
- curl
- golang
- docker desktop, helm, minikube (k8s)
- Tried bit of nix, daytona. Plan to revisit them

## AI assitance

- Started with ChatGPT 4, 4o. Currently using Claude 3.5 Sonnet
- Immensely helpful for initial research, writing initial code
- Sometimes write incorrect code, make incorrect suggestions

## Code formatting

- Only VSCode workspace setting right now. Committed to git.
- TODO: Evaluate centralized checking

## Code quality

- TODO: code style, security,

## Testing

### Functional / Integration

Yet to finalize on tool. Have briefly evaluated the following.

1. Selenium: For browser testing, not for API
1. Postman/Newman: Test script in collection file is not in a editable format outside of Postman.
1. Cucumber: Step definitions need to be defined in Java. Feels too heavy.
1. Karate: Common step definitions are pre-built. More can be added using Javascript. Revisit.
1. Keploy:
1. dredd: Looks good. Revisit
1. HTTP Files: a post by Renato Athaydes
1. Tavern: Looks good. Python based

## Continuous integration

1. Jenkins: Tried. Powerful but complex. Configuration as code option is there. Potential candidate however due to steep learning curve, did not make a lot of progress
1. Drone: Acquired by Harness. Merged in Gitness
1. Gitness: Tried. Full stack solution from code hosting to CI etc. I don't want to manage code hosting
1. Gitlab: Full stack solution
1. Dagger:
1. Tekton: Kubernetes based only
1. Concourse CI: Entirely configuration driven which could be tracked with VCS. Github stars: ~8000. Selected.

## Secret

1. AWS, Azure, GCP offering: Avoid vendor lock-in. Prefer self-hosted solution
1. Hashicorp Vault: Selected. Seems to be most popular open-source, self-hosted tool

## Deployment
- In local, have setup docker compose, k8s yamls, and helm based deployment
- TODO: Figure out production approach

# TODO

API docs, Container registry, IaC, production environment, easy setup of lower environment (okteto, garden etc.)

## Testing

- unit
- E2E
- resilience (chaos)
- Stress / load / performance
