# Go example

Setup development environment using Docker

Source : https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/

## Compile

### Linux

``` 
make PLATFORM=linux/amd64
```

#### Notes

Docker Buildx shall be installed 
``` 
docker buildx install
``` 

### Windows

``` 
make PLATFORM=windows/amd64
```

## Unit tests

```
make unit-test
```