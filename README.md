# Take Home Project - QuickNode - Krithika Nagarajan

This is a git repo and the deployment is driven by GitHub workflow and Ansible. Workflow gets triggered whenever there's a push to the main branch.
Sample curl command: curl http://192.241.195.59:8080/test

#### Checklist: 
- Daemon responds with 329d4feb-c5c0-4de5-b10c-701b44fbec4f at the /test URI ✅
- Automate the deployment of the daemon ✅
- Bare minimum “productionized” ✅
- Survive machine restarts or crashes, etc.,. ✅
- Scalable inventory ✅

## Rationale and the major why's:

Most of the tools used in this project were chosen because they belong to QuickNode's current infra tech stack. Another major goal was to make the project as simple as possible by maintaining configurability and scalability.

#### Why Golang?
Portability, ease-of-use and simple, versatile and faster build time. 

#### Why Github?
Allows a lot more free workflow runs compared to the others like BitBucket. Another reason is Github's pipeline/workflows are simpler than other version control systems.

#### Why Ansible?
Used this majorly for scalability of inventory and while keeping deployments simple. Deploying with just Github Actions to 1000s of production servers via 10000 scp's is definitely not okay!

## What else would I have added if I weren't lazy?
I would have: 
- added Prometheus metric to at least count the number of incoming GET requests
- created separate repos for application code (Go, main pipeline/workflow) and infra code (playbooks)
- dedicated user and group for deployment instead of just using root directly
- named the daemon and some vars better
- better Go code (obviously)

Time taken to complete: 1 full day. This is because I was fairly new to Ansible and Github workflows and needed quite a bit of learning and digging. libcrypto error took most of my time!
