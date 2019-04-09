### Linux安装go

- 设置PATH和GOPATH环境变量
   vim 编辑~/.bashrc, 再文件后面添加:
   ````
   export PATH=$PATH:/usr/local/go/bin
   export GOPATH=/root/go/mydocker
   ````
   然后执行`source ~/.bashrc`
   
### 工程设置
1. 新建工程目录/root/go/mydocker(也就是GOPATH变量)
2. 在$GOPATH下新建3个目录：
    - bin
    - pkg
    - src
3. 安装cli包
    在src目录下，输入以下命令安装cli包：
    ````
    go get github.com/urfave/cli
    ````
4. 安装logrus包
    在src目录下，输入以下命令安装logrus包，一般情况下会报错：
    ````
    go get github.com/Sirupsen/logrus
    package golang.org/x/sys/unix: unrecognized import path "golang.org/x/sys/unix" (https fetch: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
    ````
    logrus 包依赖于 golang官方的 golang.org/x/sys/unix 由于某些原因，不能直接获取 golang.org/x/sys/unix 
    
    但是golang 在 github 上建立了一个镜像库，如 https://github.com/golang/sys/unix 即是 golang.org/x/sys/unix 的镜像库
    
    1. 在src目录下，新建目录 golang.org/x
        ````
        mkdir -p golang.org/x
        ````
    2. 进入目录
        ````
        cd golang.org/x
        ````
    3. 拉取sys库源代码
        ````
        git clone git clone https://github.com/golang/sys.git
        ````
        成功拉取源代码后，会有sys目录
    4. 进入sys/unix
        ````
        cd sys/unix
        ````
    5. 安装unix包
        ````
        go install
        ````
    6. 重新尝试安装logrus包
        ````
        go get github.com/Sirupsen/logrus
        ````
        
    如果嫌步骤3、4麻烦，也可以直接使用脚本安装
    ````
    chmod +x install.sh && ./install.sh
    ````
    
    
### 运行程序
在src目录下，输入命令:
````
go build -o mydocker .
````
