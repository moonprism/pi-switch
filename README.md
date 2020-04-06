## raspberry pi switch app

使用 GPIO pin14 作为 IN 端实现的一个 web 开关程序。

  [blog](https://www.kicoe.com/article/id/27)

## usage

[releases](https://github.com/moonprism/pi-switch/releases) 页面下载最新版本，解压到树莓派运行：

```
chmod +x main
./main
```

浏览器访问 http://respi-ip:1323/

## dev (win~)

配置树莓派地址：

```
| - ftp.txt
| web 
|  | - .env.development
```

树莓派建立对应目录

```
sudo mkdir -p /data/switch/js
chown pi -R /data
```

build and deploy:

```
./build
```
