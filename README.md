# ready-go
spell over golang



```sh
要使用go module 首先要设置GO111MODULE=on，GO111MODULE 有三个值，off、on、auto。
auto 会根据当前目录下是否有 go.mod 文件来判断是否使用 modules 功能。
平时 GO111MODULE = off，在需要使用的时候再开启，避免在已有项目中意外引入 go module。
命令：
set GO111MODULE=on
go env // 查看 GO111MODULE 选项为 on 代表修改成功

#初始化。先进入test项目下，然后执行此命令，项目根目录会出现一个 go.mod 文件
go mod init test 
#检测依赖。tidy会检测该文件夹目录下所有引入的依赖，写入 go.mod 文件，写入后你会发现 go.mod 文件有所变动
go mod tidy 
#下载依赖。我们需要将依赖下载至本地，而不是使用 go get
go mod download 
#导入依赖。此命令会将刚才下载至 GOPATH 下的依赖转移至该项目根目录下的 vendor(自动新建) 文件夹下, 此时我们就可以使用这些依赖了
go mod vendor 
#依赖更新：这里的更新不是指版本的更新，而是指引入新依赖，依赖更新请从检测依赖部分一直执行即可：
go mod tidy
go mod download
go mod vendor

go get -u github.com/gin-gonic/gin
```
