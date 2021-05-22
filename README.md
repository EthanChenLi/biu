# Golden Snitch
> 一个由golang编写的内网穿透小工具，目的在于将内网环境的WEB服务映射到公网。

## 更新日志
|版本号 |备注 |
|--|--|
|0.1|~|
|1.1|添加配置文件|
|1.2|优化日志打印|

## 使用方式
### server端
1、 建议依赖```NGINX```作为反向代理，假如我们的域名是```go.bilibili.com```，那么nginx的配置如下
```
server {
    listen 80;
    server_name go.bilibili.com;
    location / {
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;   
        proxy_pass http://127.0.0.1:12138; 
    }
}
```
SERVER端监听的HTTP端口是```12138```，当然你也可以直接修改```conf.ini```文件。  
SERVER监听的TCP端口是```22222```，显然也可以在```conf.ini```文件中修改。  
配置好```NGINX```以后（注意云服务开放相应的端口)
### 服务端conf.ini配置文件说明

|配置名称| 可选值 | 说明 |
| --|-- | --| 
| WEB| - |  WEB服务配置 |
|HTTP_PORT |-| http服务端访问端口|
|TCP_PORT| -|  tcp服务端端口|



### Client端
#### 客户端conf.ini配置文件说明
|配置名称| 可选值 | 说明 |
| --|-- | --| 
| BASE| - |  共用配置|
|KEY | * | 访问的顶级域名名称|
|TYPE |HTTP、TCP、ALL |  http为web内网穿透，tcp为tcp代理模式，all：支持以上2种服务|
|SERVER_IP| - |服务端的IP地址|
| WEB| - |  web服务配置|
|SERVER_PORT|-| 服务端WEB端口|
|WEB_ADDR | -|需要转发的目标web服务地址|


#### 该项目作为个人应用项目，不保证有其他BUG，不可作为生产使用。欢迎PR
