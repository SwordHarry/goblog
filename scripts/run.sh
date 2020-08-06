# 带命令行参数的运行
go run ../main.go -port=8000 -mode=release -config=configs/
# 带版本信息的编译
go build -ldflags \
"-X main.buildTime=`date +%Y-%m-%d, %H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"