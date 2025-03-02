## Simple Startup Instructions

### Dependencies

- Minikube
- Helm
- Kubernetes

### Helpful Commands

Below is a list of helpful commands to get this helm chart deployment up and running on your local machine:

Ingress:

- `helm repo add nginx-stable https://helm.nginx.com/stable` (this is the ingress controller you will need for this project)
- `helm repo update`
- `helm install nginx-ingress nginx-stable/nginx-ingress --set rbac.create=true`
- You will need to create the expected project DNS for reverse proxying. Add `127.0.0.1 helm.test.cody` to your /etc/hosts file.
- `minikube tunnel` (to be used once you have started your ingress controller and applied the ingress.yaml template)

Helm Installation:
From the root directory:

- `helm install learn helm_charts` (this will create a namespace called "learn" and install the charts there)
- `kubectl apply -f helm_charts/templates/ingress.yaml`

### Installation

Order of operations to install the project and view in browser:

- `minikube start` (to start your local cluster. You can verify its running by executing `kubectl cluster-info`)
- `helm repo add nginx-stable https://helm.nginx.com/stable` (add the ingress controller)
- `helm repo update`
- `helm install nginx-ingress nginx-stable/nginx-ingress --set rbac.create=true` (start the ingress controller)
- `helm install learn helm_charts` (install the primary set of templates in helm_charts)
- `kubectl apply -f helm_charts/templates/ingress.yaml` (apply the ingress template to the ingress controller)
- `minikube tunnel` (this exposes the ingress controller to your local machine)
- Vist [helm.test.cody](http://helm.test.cody) to see the application running.
