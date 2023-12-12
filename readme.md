# Countdown API

This project provides an API to be providing the highest words for a list of letters implemented with various search algorithms. The first algorithm uses a brute force approach to finding a word given the letters. The returned array is a list of all of the words in the dictionary that could be considered a match.

## Prerequisites

To run the project:

```
go run main.go
```

In Postman or in a webbrowser send the following GET request:

```
http://localhost:3000/words/s;r;k;d;u;a;e;w;n
```

This will return with a list of words matching the letters.

Project created with:

```go mod init countdownapi
   go get -u github.com/gin-gonic/gin
```

## Build the project
To build the project run the following:
```
go build *.go
./main
```

## Tests to run

```
http://localhost:3000/words/t;h;e;r;d;a
http://localhost:3000/words/t;c;h;o;s;e;n
http://localhost:3000/words/s;r;k;d;u;a;e;w;n

```

Returns example:
```
{
    "definitions": [
        "Unawares; unexpectedly; -- sometimes preceded by at. [Obs.] Holinshed.",
        "To recant or recall, as an oath; to recall after having sworn; to abjure. J. Fletcher.\n\nTo recall an oath. Spenser.",
        "Toward the sun.",
        "A Dane. [Obs.] Inquire me first what Danskers are in Paris. Shak.",
        "Apart; separate from each other; into parts; in two; separately; into or in different pieces or places. I took my staff, even Beauty, and cut it asunder. Zech. xi. 10. As wide asunder as pole and pole. Froude."
    ],
    "dictionary": [
        "unwares",
        "unswear",
        "sunward",
        "dansker",
        "asunder"
    ],
    "userLetters": [
        "s",
        "r",
        "k",
        "d",
        "u",
        "a",
        "e",
        "w",
        "n"
    ]
}
```

## Setup as a server in AWS
In this example we build a docker container to run the application and then host this in AWS:

### Set up Docker image
We will use the golang alpine3.18 image from Dockerhub as the go installation. Pull the golang docker image from Dockerhub: 

```
docker pull golang:alpine3.18
```

The Dockerfile in the source code has been created to build the correct docker image to run an application on port 3000. Build the docker image using the following code:

```
docker build -t countdownapi .
```
And then to run this locally to test, run the following:

```
docker run -dp 127.0.0.1:3000:3000 countdownapi
```

Finally we need to push the container into Docker Hub to be using in AWS. You will need to log into Docker Hub in order to do this. Once logged in create a new repository in the web UI. This will give you a tag for the docker image. Then push this into Docker Hub:

```
docker tag bmctest/countdownapi:latest
docker push bmctest/countdownapi:latest
```

### Basic docker container running in AWS

We use Elastic Container Service to run the container. Log into AWS and navigate to the Elastic Container Service. To set up the service follow the steps below:

1. Create a new cluster in Elastic Container Service. Select AWS Fargate (serverless) as the infrastructure. Keep allother options to the default.

2. Once the cluster is set up, go to the Tasks tab and create a new task

3. Set the following defaults on the page:
    - Keep the default compute options
    - Set Deployment Configuration to Service - select the task family (if the task family is not set up then configure the task definition) The docker image is docker.io/bmctest/countdownapi:latest
    - Assign the service name
    - In networking, set up a new security group and keep Custom TCP to port 3000 open to all IPs
    - Add a load Balancer - select Application load balancer. Set the listener to port 3000
    - Use Auto Scaling

4. Example of the CloudFormation template is in the CloudFormationTemplate.json file

5. Once the task is running - navigate to the Cluster overview, select the Services tab and click on the service to view the load balancer metrics

6. Click on view load balancer. The URL to use with the API is the load balancer DNS name (A Record)

You can use the DNS name to send requests to the running service, for example:

```
http://EC2Con-EcsEl-MEcwhs3oRXvj-1468286161.eu-west-2.elb.amazonaws.com:3000/words/s;r;k;d;u;a;e;w
```
To update the container, go to the cluster service page in AWS and update the service. Selecting Force new deployment will download the latest image from Docker Hub again. 

## To do

- Add in the dictionary definition of the results (DONE)
- Add functionality to limit the size of the array being returned (DONE)
- Add instructions for running in AWS (DONE)
