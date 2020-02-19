# run本地環境

- 若下載package於 $GOPATH
```
go get github.com/gin-gonic/gin
```
- 執行主程式
```
go run main.go
```
- 連網
```
curl http://localhost:8080
```

# run AWS ECS環境
## 建立docker image
- cd到Dockerfile這一層
- 建立一個名為 golang-demo 的image
```
docker build . -t golang-demo 
```
- docker執行image為container 對外開8080對內9205
```
docker run -i -t -p 8080:9205 golang-demo
```