# go-github-webhook

## 📡 Overview
The Go-webhook is a webhook tool on github, 
That can trigger bash scripts after monitoring git's push behavior

## 📜 Usage


## 💌 Features
- Just run the binaries file 
- Custom your bash script
- Custom your enter secret
- Custom your port. 0 ~ 65535
- Print request any info

- 直接运行二进制文件
- 自定义脚本路径
- 自定义密码
- 自定义端口. 0 ~ 65535
- 打印请求的信息

```text
GLOBAL OPTIONS:
   --bash value, -b value    Execute the script path. eg: /home/hook.sh
   --port value, -p value    http port (default: 6666)
   --secret value, -s value  github hook secret
   --quiet, -q               only print errors (default: false)
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

```text
GLOBAL OPTIONS:
   --bash value, -b value    Execute the script path. eg: /home/hook.sh 自定义脚本
   --port value, -p value    http port (default: 6666) 自定义端口,默认6666
   --secret value, -s value  github hook secret 自定义密码, 不允许为空
   --quiet, -q               print any message (default: false) 安静模式,默认关闭. -q 开启,不输出任何信息
   --help, -h                show help (default: false) 
   --version, -v             print the version (default: false)

```
## 👋 Thanks
- See [GitbookIO](https://github.com/GitbookIO/go-github-webhook)
