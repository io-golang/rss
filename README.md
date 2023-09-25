# rss
This Go program fetches RSS feeds from provided URLs and sends an email to the provided email address

## Run with Docker
1. First, create a secret.list file using the ./deployment/secret.list.example file as a guide.
2. Run `docker pull iogolang/rss:latest`
3. Run `docker images` (find and copy the <image_id> for iogolang/rss:latest)
4. Finally, run `docker run --env-file ./path-to-secret.list <image-id>`

## Deploy to Kubernetes with Helm
1. Make sure you have a [Kubernetes](https://kubernetes.io) cluster (or [Minikube](https://minikube.sigs.k8s.io/docs/)) and [Helm](https://helm.sh) configured
2. Run `helm upgrade --install rss <path/to/rss/deployment/helm/rss> -n app --create-namespace`

