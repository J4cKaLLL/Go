#!/bin/bash

hash_commit=$(git rev-parse HEAD | cut -c1-7)
service_name=obtener-cliente-dom
service_id=${service_name}-v1
image_name="lgaete/${service_name}:${hash_commit}"

# Solo para minikube
eval $(minikube docker-env)

echo "Tagging images as $image_name"

docker tag $service_name $image_name