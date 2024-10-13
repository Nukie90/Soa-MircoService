#!/bin/bash

# Port forward the svelte-app-service to localhost
kubectl port-forward service/svelte-app-service 5000:5000 &
