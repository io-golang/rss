# rss
This RSS program fetches RSS feeds from provided URLs and sends an email to the provided email address

## Build and run with docker
cd to root of project
docker build -f ./deployment/Dockerfile . --tag io-golang-rss
docker run --env-file ./deployment/secret.list <image-id>
