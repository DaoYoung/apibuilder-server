# 婚礼纪comet

## 安装
1. 下载并运行[Nats](https://www.nats.io)
2. 创建配置文件`config.yml`，可从`config.yml.example`拷贝
3. 运行`hlj-comet`

## 部署
1. 编译成linux可执行文件
```cmd
build.bat linux
```
2. 部署到测试服务器
```cmd
deploy.sh test
```
3. 在服务器上执行
```bash
./run.sh update
./run.sh restart
```