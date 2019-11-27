az aks get-credentials --resource-group <ResourceGroup> --name <AKSCluster>

kubectl get ns

kubectl -n <namespace> get all

kubectl -n <namespace> get pods

kubectl -n kn-seti-dev-tf get pods

kubectl -n <namespace> logs <pod spec>

kubectl -n <namespace> exec -it <pod name> -- /bin/bash

kubectl delete ns <namespace>
