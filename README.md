# rss
This RSS program fetches RSS feeds from provided URLs and sends an email to the provided email address

## Run with docker
1. First, create a secret.list file using the ./deployment/secret.list.example file as a guide.
2. Run `docker pull iogolang/rss:latest`
3. Run `docker images` (find and copy the <image_id> for iogolang/rss:latest)
4. Finally, run `docker run --env-file ./path-to-secret.list <image-id>`
