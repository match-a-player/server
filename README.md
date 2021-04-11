# server
The go backend

## How to build

```
docker build -t matchaplayer.jfrog.io/app/server .
docker run --rm  -it -p 8080:8080 matchaplayer.jfrog.io/app/server
```

## Push to artifactory
* Use this only when it's definitely needed in artifactory. Our limits can get exhausted.

```
docker login matchaplayer.jfrog.io (Only the first time)
docker push matchaplayer.jfrog.io/app/server
```

