# ansible-go-app-proto-guide

#### Description:

Sample bare minimum “productionized” working prototype for guiding deployments driven by GitHub actions workflow and Ansible. Daemon responds with a code to a GET request at the /test URI ( eg http://192.241.195.59/test ). Deployment of the daemon to the inventory Linux server is automated. Workflow is triggered whenever there's a push to the main branch.
Sample curl command: curl http://192.241.195.59:8080/test

#### Checklist: 
- Daemon responds with 329d4feb-c5c0-4de5-b10c-701b44fbec4f at the /test URI ✅
- Automate the deployment of the daemon ✅
- Bare minimum “productionized” ✅
- Survive machine restarts or crashes, etc.,. ✅
- Scalable inventory ✅

## Rationale and the major why's:

Major goal was to make the project as simple as possible by maintaining configurability and scalability.

#### Why Golang?
Portability, ease-of-use, simple, versatile and faster build time. 

#### Why Github?
Allows a lot more free workflow runs compared to the others like BitBucket. Another reason is Github's pipeline/workflows are simpler than other version control systems.

#### Why Ansible?
Used this majorly for the scalability of inventory and while keeping deployments simple. Deploying with just Github Actions to 1000s of production servers via 10000 scp's is definitely not okay!

## Possible improvisations
- added Prometheus metric to at least count the number of incoming GET requests
- created separate repos for application code (Go, main pipeline/workflow) and infra code (playbooks)
- dedicated user and group for deployment instead of just using root directly
- named the daemon and some vars better
- better Go code (obviously)
