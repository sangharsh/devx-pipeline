## Intent

1. Understand the best possible developer experience (DevX) across software lifecycle, from product specs to deployment, using available tools.
1. Identify improvement areas

Engineering team (developer, QA, Devops) should be able to deliver business value in most efficient way possible. Underlying tools, platforms should take care of everything possible by machines, provide timely support and feedback. Individual should be able to function as independently as possible without waiting/blocking on another individual.

Sub point note:

> Be able to access or setup an environment as per their need in a seamless way.

## Approach

Starting with web domain as I am most familiar with it. Usually startup begin with a monolith, over time some microservices are spun up to address new requirements or carve out the monolith. Typical companies have a bunch of services (1-few monoliths, 10 to 500 microservices, depending on org's size). Many services have their own datastore and communicate with each other to facilitate user flows, popularly HTTP request/response following REST convention and JSON data format.

Take a simplistic slice of popular system setup of startup/companies.
Build an end to end tool chain for software development using best available tools for the purpose. Prefer to use tools which are free, open-source, self-hosted. Keep all code, configuration for setup from scratch in VCS.

Many of cutting edge tools might be commercial or internal. Commercial tools might be considered for this project if there is a free tier which is not time-limited.
