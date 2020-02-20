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
集群 > 開始使用 
#### Container definition
- 選擇： custom
- 容器名稱： 自己命名
- 映像：貼上 ECR的URI
- 端口映射： 已本例為例 填9205

#### Task defination
- 負載均衡器類型：Application Load Balancer
- 監聽窗口： 9205

next > 創建集群 > loading > 查看服務

### 連網 
點擊 EC2 負載均衡器 > 找到 DNS 名稱後，加上端口名稱:9205
- 例如： 
EC2Co-EcsEl-1AWU1LI9XRXJ3-1697609552.ap-northeast-2.elb.amazonaws.com:9205
- 再加上指定網頁路徑 如 /user/info 後即可連上瀏覽器