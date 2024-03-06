# container-apps-multicontainers
Two examples using Envoy and NGINX to forward requests to a 2-container based pod.

You can deploy and run this with the following:
1. Build the `Dockerfile` (s) located under `envoy-example/envoy` and `envoy-example/go` as well as `nginx-example/nginx` and `nginx-example/go`
2. Push these images to Azure Container Registry
3. Populate the environment variables referenced under `deploy/bash.sh`
4. Use the ARM template to deploy. Note, that this is implying a Container App Environment already exists. This template is only used to create or update a Container App - not any other resources.