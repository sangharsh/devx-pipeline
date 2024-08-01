# Developer Experience

End to end tool chain for software building with high emphasis for Developer experience (DevX or DX). DX in this project's context includes all engineering functions (development, testing, deployment, operations)

> Developer Experience (DX) is all about how easily and effectively developers can get things done.
>
> -- <cite>Addy Osmani</cite>'s [book](https://read-dx.addy.ie/preface)

# Sample business requirement

- Enable user to place an order through their browser.
- Enable internal team (ops) to manage inventory.
- Track inventory as orders are placed.

# Goals

List is incomplete and WIP.

1. Apps should be easy to setup and test for someone new to the codebase, including pre-requisites, infra, networking and seed data.
1. Changes in code should be easy to test along with dependent services as needed. Possible solution: Ephermeral environments
1. Most setups and configurations should be visible and repeatable. Solution: Prefer code driven rather than configuring a tool through UI.

# Tech stack

Evaluation framework. Tool being used and alternatives considered are mentioned in [Techstack](/Techstack.md)

# Status

1. A basic golang web app for inventory
1. Testing using testcontainers
1. Build using dockerfile
1. Deployment using docker-compose, k8s yaml (minikube), helm
1. WIP: CI using concourse
1. Secret: Vault (Hashicorp)
