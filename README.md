# Fullstack Chat Room Application

This application allows users to create accounts, login, create chat room, switch chat room, send messages, and upvote or downvote messages.

## Prerequisites

Before you start, make sure you have installed:

<ul>
  <li>Golang</li>
  <li>React.js</li>
  <li>PostgreSQL</li>
  <li>Docker</li>
  <li>Minikube</li>
  <li>Kubectl</li>

</ul>

## Usage

### Backend & Database

1. Start the PostgreSQL database
2. Navigate to the backend
   ```console
   $ cd ./server
   ```
3. Install the required Go dependencies
   ```console
   $ go mod download
   ```
4. Copy .env.example to .env and update the variables to match your database configuration
   ```console
   $ cp .env.example .env
   ```
5. Start the Go server
   ```console
   $ go run main.go
   ```

### Frontend

1. Navigate to the frontend
   ```console
   $ cd ./client
   ```
2. Install the required Node.js dependencies
   ```console
   $ npm install
   ```
3. Start the React development server
   ```console
   $ npm start
   ```

## Building and Deploying with Docker and Kubernetes

1. Run Docker
2. Start Minikube
   ```console
    $ minikube start
   ```
3. Set the Docker Environment
   ```console
    $ eval $(minikube -p minikube docker-env)
   ```
4. Build Docker Images

   ```console
   # Go to the backend directory
   $ cd ./server

   # Build the Docker image for the backend
   $ docker build -t chatroom-backend .

   # Do the same for the frontend
   $ cd ./client
   $ docker build -t chatroom-frontend .
   ```

5. Create your deployments and services:
   ```console
   $ kubectl apply -f frontend-deployment.yaml
   $ kubectl apply -f backend-deployment.yaml
   $ kubectl apply -f frontend-service.yaml
   $ kubectl apply -f backend-service.yaml
   ```
6. Expose Services
   ```console
   $ minikube service chatroom-backend --url
   $ minikube service chatroom-frontend --url
   ```
