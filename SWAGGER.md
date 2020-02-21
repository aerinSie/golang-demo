若無法順利import依賴包，可以試看看下方指令
- cmd export 都是臨時的，要變成永久的話，要寫在設定檔裡
```
export GO111MODULE=on
export PATH=$PATH:$GOPATH/bin
```
- 升級所有依賴
```
go get -u
```