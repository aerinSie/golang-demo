# run 本地環境

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
curl http://localhost:9205
```

# run AWS ECS環境
## 建立docker image
- cd到Dockerfile這一層
- 建立一個名為 golang-demo 的image
```
docker build . -t golang-demo 
```
## 本地docker運行
- docker執行image為container 對外開8080對內9205
```
docker run -i -t -p 8080:9205 golang-demo
```
- 連網
```
curl http://localhost:8080
```
## AWS ECS運行
### 上傳到ECR
- 需要先安裝AWS CLI
- 需要先登入AWS CLI

進入 ECR > 存儲庫 > 專案名稱 > 查看推送命令

- 登入ecr 以 region=ap-northeast-2為例
```
aws ecr get-login --no-include-email --region ap-northeast-2
```
if ok 會顯示 Login Succeeded

if fail ，用下列的指令看看。以AWS帳戶ID=000000000000 為例
```
aws ecr get-login-password | docker login --username AWS --password-stdin 000000000000.dkr.ecr.ap-northeast-2.amazonaws.com
```
- 生成docker image 以專案名稱=golang-demo為例
```
docker build -t golang-demo
```
- 從本機推到ECR 
```
docker tag golang-demo:latest 000000000000.dkr.ecr.ap-northeast-2.amazonaws.com/golang-demo:latest

docker push 000000000000.dkr.ecr.ap-northeast-2.amazonaws.com/golang-demo:latest
```
如果成功 [ECR > 存儲庫 > 專案名稱] 底下會新增一個image檔案

### run ECS
集群 > 