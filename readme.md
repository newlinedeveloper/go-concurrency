### Go Concurrency

Go language has built-in support for concurrency through goroutines and channels.

- A goroutine is a lightweight thread of execution that can be created with the "go" keyword. Goroutines run concurrently with other goroutines within the same program.

- Channels are a way to communicate between goroutines and synchronize their execution. A channel can be used to send and receive values between goroutines.

- This approach to concurrency allows for clear and efficient management of multiple tasks without the need for locks or other synchronization mechanisms.

#### Start the application

```
# To run main.go
go run main.go


# To check in browser
http://localhost:8080/book

```

#### Dockerize the go application

```
# To build docker image for go application 
docker build -t go-app .


# To run the docker image
docker run --name go-app -p 8080:8080 go-app

```


#### K8 deployment

```
cd k8

# To run deployment and service configuration for this docker image 

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

```