使用go module 需要 在settings 下 go - go modules 下勾选enable go modules

go module 使用 https://www.bilibili.com/video/av63052644/

go module 使用国内阿里云代理
go env -w GOPROXY="https://mirrors.aliyun.com/goproxy/"

在项目下 go mod init <projectName> 初始化 创建go.mod文件


用户登录用的是gin-session

登录以后会把用户登录信息加密以后写入到cookie里去，浏览器会保存这个cookie，下次请求的时候带上。可以设定加密秘钥。

服务端收到请求以后就会对加密的cookie（实际保存格式是gin-session）逆向生成用户id。。能成功生成就说明用户是登录了的。

目前发现的问题，
1.cookie过期依赖的是浏览器，如果客户端手动修改过期时间，服务端还是做为已经登录处理的。
2.登出实际上只是浏览器清除了cookie。客户每次请求也不会刷新cookie过期时间。
3.因为cookie数据完全存在客户端，所以无法做到强制退出，过期校验等功能。要做的话可能就要把相关数据持久化了。
  
