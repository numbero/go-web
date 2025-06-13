# 默认编译
go build -o go-web
# 交叉编译
GOOS=linux GOARCH=amd64 go build -o go-web_amd64 
# upx压缩
upx --best go-web_amd64       