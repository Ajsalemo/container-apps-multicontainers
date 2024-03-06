az deployment group create \
  --resource-group "$RESOURCE_GROUP" \
  --template-file ./arm.json \
  --parameters \
      environment_name="$CONTAINERAPPS_ENVIRONMENT" \
      location="$LOCATION" \
      azureContainerRegistryServerUrl="$AZURE_CONTAINER_REGISTRY_SERVER_URL" \
      azureContainerRegistryPassword="$AZURE_CONTAINER_REGISTRY_PASSWORD" \
      azureContainerRegistryUsername="$AZURE_CONTAINER_REGISTRY_USERNAME" \
      azureContainerRegistryNginxImageName="$AZURE_CONTAINER_REGISTRY_NGINX_IMAGE_NAME" \
      azureContainerRegistryNginxImageTag="$AZURE_CONTAINER_REGISTRY_NGINX_IMAGE_TAG" \
      azureContainerRegistryGoImageName="$AZURE_CONTAINER_REGISTRY_GO_IMAGE_NAME" \
      azureContainerRegistryGoImageTag="$AZURE_CONTAINER_REGISTRY_GO_IMAGE_TAG" \
      azureContainerRegistryInitContainerImageTag="$AZURE_CONTAINER_REGISTRY_INIT_CONTAINER_IMAGE_TAG" \
      azureContainerRegistryInitContainerImageName="$AZURE_CONTAINER_REGISTRY_INIT_CONTAINER_IMAGE_NAME"