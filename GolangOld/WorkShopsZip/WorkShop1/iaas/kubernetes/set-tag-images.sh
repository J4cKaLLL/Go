#!/bin/bash

hash_commit=$(git rev-parse HEAD | cut -c1-7)

echo "Tagging images as ${hash_commit}"
sed -i.bu "s/dom:[[:alnum:]]*/dom:${hash_commit}/" iaas/kubernetes/deploy.yaml