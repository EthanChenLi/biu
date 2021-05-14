# Golden Snitch
> 一个由golang编写的内网穿透小工具，目的在于将内网环境的WEB服务映射到公网。

## 更新日志
|版本号 |备注 |
|--|--|
|0.1|~|

## 使用方式
### server端
1、 需要依赖```NGINX```作为反向代理，假如我们的域名是```gogogo.bilibili.com```，那么nginx的配置如下
```
server {
    listen 80;
    server_name gogogo.bilibili.com;
    location / {
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;   

        proxy_pass http://127.0.0.1:12138; 
    }
}
```
SERVER端监听的HTTP端口是```12138```，当然你也可以直接修改代码里面的端口。  
SERVER监听的TCP端口是```22222```，显然也可以在代码中修改。  
配置好```NGINX```以后（注意云服务开放相应的端口)

### Client端
显然，目前需要你有一个golang的环境，因为现在暂时不支持命令行配置SERVER端的IP。
在```main.go```文件中修改服务端的IP地址
```golang
const TcpAddr = "[您的TCP端口地址]:22222"
```
然后运行```go build``` ，程序编译好以后即可执行
#### 相关命令
* host  ： 唯一的KEY，这是一个二级域名的转发KEY，在当前案例中应该是设置```gogogo```
* ip : 默认127.0.0.1 你需要映射的WEB服务的IP地址
* port : 默认80,你需要映射的WEB服务的端口  
举个栗子:
``` main -host gogogo -port 8080```

#### 该项目作为个人应用项目，不保证有其他BUG，不可作为生产使用。欢迎PR