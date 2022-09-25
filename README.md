# Onboarding

This repo have a bunch of hands on assignments to onboard new joiners to Build and Deployment setup at Rapido.

> ## Create a new branch and start adding commands that you used, k8s manifest, dockerfiles and add any notes as a reference for future.

## We cover the following here:
- Docker
- Kubernetes
- Helm
- GoCD
- Hermes

## Sample Service Overview
![Sample Svc](/docs/images/svc.png)

## Product Details (PDS)
- Service that expose product data with following APIs:
  - `GET  /ping`  - Ping API/ can be used for liveliness check.
  - `GET  /products` - API to return all Product Details, this in turn calls the `product-reviews` service for review details.
  - `GET  /products/:id` - API to return Details of a particular product.
- Currently the data is hard-coded and the valid product IDs are `1` and `2`
- `export REVIEW_SVC_HOST=http://localhost:8080`this service uses this env var to talk to reviews service.
- Listens on port `9090`

## Product Reviews (PRS)
- Service that provide product review data, with following APIs:
  - `GET  /products/:id/reviews`
  - `GET  /ping`
- Data is available for both the product ids in this service.
- Listens on port `8080`

## Tasks

### Build Locally
- Build and Run both the services and test if the APIs are working as expected.

### Containerize
- Install [Docker for mac](https://docs.docker.com/docker-for-mac/install/).
- Add a dockerfile.
- Build a docker image.
- Push to [docker hub](https://hub.docker.com/).
- Run the containers and ensure PDS is able to talk to PRS.

### K8S
- Install [mini-kube](https://minikube.sigs.k8s.io/docs/start/) or [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) cluster on laptop.
- Install [kubectx](https://github.com/ahmetb/kubectx)
- Install [stern](https://github.com/wercker/stern)
- Run PRS in k8s as a stand alone pod.
  - try accessing the APIs by port-forwarding.
- Deploy a [Replica-Set](https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/) of PRS using manifest file.
  - Scale the PRS to 2 replicas.
  - Add a cluster-ip [service](https://kubernetes.io/docs/concepts/services-networking/service/) and access the PRS by port-forwarding on the svc object.
  - Stern on the logs of the pods and ensure requests are getting routed to different pods.
- Use [Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/) and run both the services using manifest files.
  - Set `REVIEW_SVC_HOST` as an ENV Var, use the services DNS name.
  - Use a wrong value for `REVIEW_SVC_HOST` and apply deployment changes.
  - Rollout restart the deployment.
  - Now Rollback to previous version and ensure things are working as expected.


### Helm 
- Install [helm](https://helm.sh/) cli - `brew install helm`
- Create one helm chart per service and drive deployment using it.
- Deploy the helm chart.
- Read the values for image and ENV variables from values file.
- Explore history and rollback command

### Deploy to lower env cluster & Istio
- Raise PR to infra-acm repo and get access to dev and staging environment.
- Create a dummy namespace and deploy the services from local to this environment.
- Add istio side-car injection support.
- Add a virtual service and expose the service using istio ingress gateway.
- Inject fault for 50% API requests using istio.
- Check the istio grafana dashboard for metrics on this service.
- Lookup the logs for these services in loki (only if deployed in staging).

### GoCD
- Add a build and deployment pipeline using gocd for these services (Docker build and helm deploy).
  - Create an elastic profile.
  - Add pipeline config file to repo.
  - Add it to config repo.

### Hermes
- Migrate this service to use hermes (docker build mode)
