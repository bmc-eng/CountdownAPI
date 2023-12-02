# Countdown API

This project provides an API to be providing the highest words for a list of letters implemented with various search algorithms. The first algorithm uses a brute force approach to finding a word given the letters. The returned array is a list of all of the words in the dictionary that could be considered a match.

## Prerequisites

To run the project:

```
go run main.go
```

In Postman or in a webbrowser send the following GET request:

```http://localhost:8080/words/t;h;e;r;d;a```

This will return with a list of words matching the letters.

Project created with:

```go mod init countdownapi
   go get -u github.com/gin-gonic/gin
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
    "dictionary": [
        "thread",
        "hatred",
        "dearth"
    ],
    "test": [
        "a",
        "a",
        "a",
        "a",
        "a",
        "a"
    ]
}
```

## To do

- Add in the dictionary definition of the results
- Add 


## Contributions