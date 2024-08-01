# Tech design

Build two API services for orders and inventory and a frontend for browser.

# Evaluation framework

Platforms, tools for end to end enablement are selected based on following factors:

1. **Free and open-source**. Tools being used should be free, preferably open-source or atleast open-core, making it adoptable for wider community. Also, I don't want incure any ongoing cost to host or run.
1. **Code driven**. Most of the things should be source code driven (gitops?).
1. **Vendor lock-in**: Avoid.
1. **Reproduciblity**
1. **Popularity** Prefer a set of popular tools which have wider adoption if other things are same

Continuously evaluate each step for developer experience.

# Requirement and Tools

## Project requirement documentation

- TODO: Doesn't exist currently. Evaluate need. Options: Redmine, Proprietary: Github, Confluence, Jira

## Code hosting

Code base is currently only persistent data in this project. Self hosting comes with a cost. Hence, opting to host it on a code host with free offering for open source projects. Github being the most popular tool is selected for this. Project should avoid using Github specific features like Actions to avoid vendor lock-in.

## Repository structure

Trying mono repo approach

## Development environment

- VSCode
- curl
- golang
- docker desktop, helm, minikube (k8s)
- Tried bit of nix, daytona. Plan to revisit them

## AI assitance

- Started with ChatGPT 4, 4o. Currently using Claude 3.5 Sonnet
- Immensely helpful for initial research, writing initial code
- Sometimes, it writes incorrect code, makes incorrect suggestions

## App development

| Function               | Choice                    |
| ---------------------- | ------------------------- |
| Protocol               | HTTP(S)                   |
| Data format            | JSON                      |
| API convention         | REST                      |
| Language and Framework | golang, Java (springboot) |
| Database               | Postgresql                |
| Cache                  | Redis or OSS alternative  |
| Message broker         | TBD                       |
| Web frontend           | React/Vue (TBD)           |

Develop a set of services which expose REST API with JSON request/response. Use golang for one and Java (springboot) for another. Both will have their own databases. (Try to) Add a cache (redis), a message broker

Use React/Vue for a frontend

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
1. HTTP Files: a post by Renato Athaydes [link](https://renato.athaydes.com/posts/writing-http-files-for-testing.html)
1. Tavern: Looks good. Python based

## Continuous integration

1. Jenkins: Tried. Powerful but complex. Configuration as code option is there. Potential candidate however due to steep learning curve, did not make a lot of progress
1. Drone: Acquired by Harness. Merged in Gitness
1. Gitness: Tried. Full stack solution from code hosting to CI etc. I don't want to manage code hosting
1. Gitlab: Full stack solution
1. Dagger:
1. Tekton: Kubernetes based only
1. **Concourse CI**: Entirely configuration driven which could be tracked with VCS. Github stars: ~8000. Selected.

## Secret

1. AWS, Azure, GCP offering: Avoid vendor lock-in. Prefer self-hosted solution
1. **Hashicorp Vault**: Selected. Seems to be most popular open-source, self-hosted tool

## Deployment

1. **Virtual Machines**. Teams should start with simple VM based deployments when they have a couple of services. VMs are simpler from learning perspective and easy to manage when service footprint is smaller.
1. **Containers and orchestration**. As service footprint increases, investment into container and orchestration makes sense.
   - **Kubernetes**. Most popular and open-source container orchestration framework. Defacto choice.

Opting for containers and Kubernetes because:

1. Easier to create and destroy.
1. Popular choice among mid to large companies with many microservices.

Current status:

- In local, have setup docker compose, k8s yamls, and helm based deployment
- TODO: Figure out production setup and intermediate lower environments

# TODOs

1. API docs
1. Container registry
1. Infrastructure as code
1. Production environment
1. Easy setup of lower environment (okteto, garden, skaffold etc.)
1. Logging, monitoring, alerts, incidents.
   1.Testing
   1. unit
   1. E2E
   1. resilience (chaos)
   1. Stress / load / performance
