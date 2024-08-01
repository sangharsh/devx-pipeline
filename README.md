# Developer Experience

## Intent

1. Understand the best possible developer experience (DevX) across software lifecycle, from product specs to deployment, using available tools.
1. Identify improvement areas

## What is DevX

Referring to Addy Osmani's [book](https://read-dx.addy.ie/preface)

> Developer Experience (DX) is all about how easily and effectively developers can get things done.

Engineering team (developer, QA, Devops) should be able to deliver business value in most efficient way possible. Underlying tools, platforms should take care of everything possible by machines, provide timely support and feedback. Individual should be able to function as independently as possible without waiting/blocking on another individual.

Sub point note:

> Be able to access or setup an environment as per their need in a seamless way.

## Approach

Starting with web domain as I am most familiar with it. Usually startup begin with a monolith, over time some microservices are spun up to address new requirements or carve out the monolith. Typical companies have a bunch of services (1-few monoliths, 10 to 500 microservices, depending on org's size). Many services have their own datastore and communicate with each other to facilitate user flows, popularly HTTP request/response following REST convention and JSON data format.

Take a simplistic slice of popular system setup of startup/companies.
Build an end to end tool chain for software development using best available tools for the purpose. Prefer to use tools which are free, open-source, self-hosted. Keep all code, configuration for setup from scratch in VCS.

Many of cutting edge tools might be commercial or internal. Commercial tools might be considered for this project if there is a free tier which is not time-limited.

# Sample business requirement

Build a system to take simple e-commerce orders and manage inventory through a browser-based interface

# Tech design

Develop a set of services which expose REST API with JSON request/response. Use golang for one and Java (springboot) for another. Both will have their own databases. (Try to) Add a cache (redis), a message broker

Use React/Vue for a frontend

# Project setup

Tools being used should be free, preferably open-source or atleast open-core. Being a personal project, it is imperative to not have a recurring cost to host, run etc.
Most of the things should be source code driven (gitops?). Avoid vendor lock-in. Keep things reproducible.

Use a set of popular tools to build.

Continuously evaluate each step for developer experience.

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
- TODO: Figure out production setup and intermediate lower environments

# TODO

API docs, Container registry, IaC, production environment, easy setup of lower environment (okteto, garden etc.)
Logging, monitoring, alerts, incidents.

## Testing

- unit
- E2E
- resilience (chaos)
- Stress / load / performance
