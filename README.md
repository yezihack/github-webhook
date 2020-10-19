# github-webhook
[![Build Status](https://travis-ci.org/yezihack/github-webhook.svg?branch=master)](https://travis-ci.org/yezihack/github-webhook)
[![](https://img.shields.io/github/release/yezihack/github-webhook?style=flat-square)](https://github.com/yezihack/github-webhook/releases)
[![](https://img.shields.io/github/license/yezihack/github-webhook?style=flat-square)](https://github.com/yezihack/github-webhook/blob/master/LICENSE)
![](https://img.shields.io/github/repo-size/yezihack/github-webhook?style=flat-square)


https://img.shields.io/appveyor/build/yezihack/github-webhook
## 📡 Overview
The Github-webhook is a webhook tool on github, 
That can trigger bash scripts after monitoring git's push behavior

The a line command handles the automatic build

github-webhook 是github webhook自动构建工具.能监听git push行为,自动触发脚本.

一条命令搞定webhook自动构建,无需复杂的配置.

## 📜 Usage
### 1. Install
`go get -u github.com/yezihack/github-webhook`
### 2. Download
- [releases](https://github.com/yezihack/github-webhook/releases)
- github release下载太慢, 试试这个
```
wget http://img.sgfoot.com/github-webhook1.4.1.linux-amd64.tar.gz
```

```shell script
cp ~/github-webhook /usr/bin
chmod u+x /usr/bin/github-webhook
```

## 3. Command
- Daemonize run:  `nohup github-webhook --bash /home/my.sh --secret mysecret -q &`  后台运行
- Monitor run: `github-webhook --bash /home/my.sh --secret mysecret`
- Quiet mode run: `github-webhook --bash /home/my.sh --secret mysecret --quiet`
- Custom port mode run: `github-webhook --bash /home/my.sh --secret mysecret --port 6100 --quiet`
- Hidden secret mode run: `github-webhook --bash /home/my.sh  --quiet` 

add systemd service
> /home/sh/hugo2www.sh is your script bash file
```shell script
cat > /lib/systemd/system/webhook << EOF
[Unit]
Description=github-webhook
Documentation=https://github.com/yezihack/github-webhook
After=network.target
 
[Service]
Type=simple
ExecStart=/usr/bin/github-webhook --bash /home/sh/hugo2www.sh --secret qweqwe
Restart=on-failure
RestartSec=42s
 
[Install]
WantedBy=multi-user.target
EOF
```
```shell script
systemctl daemon-reload
systemctl start webhook
systemctl status webhook
```


## 4. WebHook
- Default port: 2020
- Http path: /web-hook
- Test URL: `http://ip:2020/ping`
- WebHook URL: `http://ip:2020/web-hook`


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
   --port value, -p value    http port (default: 2020)
   --secret value, -s value  github hook secret
   --quiet, -q               quiet operation (default: false)
   --verbose, --vv           print verbose (default: false)
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```
中文
```text
GLOBAL OPTIONS:
   --bash value, -b value    Execute the script path. eg: /home/hook.sh 自定义脚本
   --port value, -p value    http port (default: 2020) 自定义端口,默认6666
   --secret value, -s value  github hook secret 自定义密码, 不允许为空
   --verbose, --vv           print verbose (default: false) 打印更多详细信息
   --quiet, -q               quiet operation (default: false) 安静模式,默认关闭. -q 开启,不输出任何信息
   --help, -h                show help (default: false) 
   --version, -v             print the version (default: false)

```
# How it works

![github-webhook](https://upload-images.jianshu.io/upload_images/13827699-49011566250e8250.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


- step 1:: Run your github-webhook server

  - notice: port default 2020, http-path: /web-hook
  - 注意: 端口默认为 2020, 可以更改, http的路由: /web-hook
  - 查看自己的外网Ip: `curp ip.sb`

  ![](assets/image-20200422194800401.png)

- step 2: Add webhook
  - 添加 webhook 参数
  
    ![配置第一步](https://upload-images.jianshu.io/upload_images/13827699-4aa2488f63658de4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
    
    ![配置第二步](https://upload-images.jianshu.io/upload_images/13827699-f3866693a5db8df2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
    
    ![配置第三步](https://upload-images.jianshu.io/upload_images/13827699-09a4de85b8b2b006.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- step 3: run shell script
 - notice: Make sure that the last line write: exit 0
 - shell脚本的最后一行一定要写上 `exit 0` 代码
```
#!/bin/bash
echo "hello webhook"
exit 0
```
## 👋 Thanks
- See [GitbookIO](https://github.com/GitbookIO/go-github-webhook)
