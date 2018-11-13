# Start minikube
# minikube start

# Set docker env
eval $(minikube docker-env)
# eval $(minikube docker-env -u) -- to stop using minikube's docker environment

# Build image
docker build -t trening-backend:latest .

# Run in minikube
kubectl apply -f trening-backend-pod.yml --image-pull-policy=Never

# Check that it's running
kubectl get pods