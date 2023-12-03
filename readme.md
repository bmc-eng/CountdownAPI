# Countdown API

This project provides an API to be providing the highest words for a list of letters implemented with various search algorithms. The first algorithm uses a brute force approach to finding a word given the letters. The returned array is a list of all of the words in the dictionary that could be considered a match.

## Prerequisites

To run the project:

```
go run main.go
```

In Postman or in a webbrowser send the following GET request:

```
http://localhost:8080/words/s;r;k;d;u;a;e;w;n
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
http://localhost:8080/words/t;h;e;r;d;a
http://localhost:8080/words/t;c;h;o;s;e;n
http://localhost:8080/words/s;r;k;d;u;a;e;w;n

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

### Basic docker container running in AWS

## To do

- Add in the dictionary definition of the results (DONE)
- Add functionality to limit the size of the array being returned (DONE)
- Add instructions for running in AWS


## Contributions