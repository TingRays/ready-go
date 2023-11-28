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

#### 关于 .mod 文件
module 表示模块名称
require 依赖包列表以及版本，是不需要自己手动去修改的，当运行代码的时候，会根据代码中用到的包自动去下载导入。
exclude 禁止依赖包列表，不下载和引用哪些包(仅在当前模块为主模块时生效)
replace 替换依赖包列表和引用路径(仅在当前模块为主模块时生
replace 他可以将代码中使用，但国内被墙的代码替换成 github上的下载路径，例如：golang.org/x/ 下的包，全都替换成 github地址上的包，版本使用 latest 即可。
replace 指令可以将依赖的模块替换为另一个模块，例如由公共库替换为内部私有仓

replace golang.org/x/net v1.2.3 => example.com/fork/net v1.4
replace (
	golang.org/x/net => github.com/golang/net latest
	golang.org/x/tools => github.com/golang/tools latest
	golang.org/x/crypto => github.com/golang/crypto latest
	golang.org/x/sys => github.com/golang/sys latest
	golang.org/x/text => github.com/golang/text latest
	golang.org/x/sync => github.com/golang/sync latest
)

indirect 表示这个库是间接引用进来的。
使用 go list -m all 可以查看到所有依赖列表，也可以使用 go list -json -m all 输出 json格式的打印结果。

除了 go.mod 之外，go 命令行工具还维护了一个 go.sum 文件，它包含了指定的模块的版本内容的哈希值作为校验参考：
go 命令行工具使用 go.sum 文件来确保你的项目依赖的模块不会发生变化——无论是恶意的，还是意外的，或者是其它的什么原因。go.mod 文件和 go.sum 文件都应该保存到你的代码版本控制系统里面去。

go.sum 这个文件记录了源码的直接依赖和间接依赖包的相关版本的 hash 值，用来校验本地包的真实性。在构建的时候，如果本地依赖包的 hash 值与 go.sum 文件中记录的不一致，就会被拒绝构建，这样可以确保你的项目所依赖的 module 内容，不会被恶意或意外篡改。