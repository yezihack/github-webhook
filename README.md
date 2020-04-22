# go-github-webhook

## 📡 Overview
The Go-webhook is a webhook tool on github, 
That can trigger bash scripts after monitoring git's push behavior

## 📜 Usage
Golang install:`go get github.com/yezihack/go-webhook`
Or Download: [Release](https://github.com/yezihack/go-webhook/releases)

- Default Run: `go-webhook --bash /home/my.sh --secret mysecret`
- Quiet Mode Run: `go-webhook --bash /home/my.sh --secret mysecret --quiet`
- Custom Port Mode Run: `go-webhook --bash /home/my.sh --secret mysecret --port 6100 --quiet`
- Hidden Secret Mode Run: `go-webhook --bash /home/my.sh --quiet`


## 💌 Features
- Just run the binaries file 
- Custom your bash script
- Custom your enter secret
- Custom your port. 0 ~ 65535
- Quiet operation

中文 
- 直接运行二进制文件
- 自定义脚本路径
- 自定义密码
- 自定义端口. 0 ~ 65535
- 安静模式

```text
GLOBAL OPTIONS:
   --bash value, -b value    Execute the script path. eg: /home/hook.sh
   --port value, -p value    http port (default: 6666)
   --secret value, -s value  github hook secret
   --quiet, -q               quiet operation (default: false)
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```
中文
```text
GLOBAL OPTIONS:
   --bash value, -b value    Execute the script path. eg: /home/hook.sh 自定义脚本
   --port value, -p value    http port (default: 6666) 自定义端口,默认6666
   --secret value, -s value  github hook secret 自定义密码, 不允许为空
   --quiet, -q               quiet operation (default: false) 安静模式,默认关闭. -q 开启,不输出任何信息
   --help, -h                show help (default: false) 
   --version, -v             print the version (default: false)

```
# HOW DOING

![](help/ae3edeb82083683a.jpg)


- step 1:: Run your go-webhook server

  notice: port default 6666, http-path: /web-hook
  注意: 端口默认为6666, 可以更改, http的路由: /web-hook
  查看自己的外网Ip: `curp ip.sb`

  ![image-20200422194800401](assets/image-20200422194800401.png)

- step 2: Add webhook
  添加 webhook 参数
  
  ![image-20200422194224139](assets/image-20200422194224139.png)
  ![image-20200422195200683](assets/image-20200422195200683.png)



## 👋 Thanks

- See [GitbookIO](https://github.com/GitbookIO/go-github-webhook)
