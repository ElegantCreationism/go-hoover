# go-hoover

### How to use this service 
Navigate to the Project root:
> run:\
> ``go run main.go``

Then curl the service using this:
```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"roomSize" : [5, 5],"coords" : [2, 2],"patches" : [[0, 1],[3, 1],[2, 0]],"instructions" : "N,N,W,S,S,S,W,N"}' \
  http://0.0.0.0:8080/roomba
```

> Instructions are given in comma-separated format for efficient filtering.\
> North: N\
> East:  E\
> South: S\
> West:  W


```go

```
