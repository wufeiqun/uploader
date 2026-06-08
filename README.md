# 项目介绍

`uploader`是用Golang编写的一个简单的文件上传下载服务, 用于CICD的过程中制品的存储和下载.

# 项目构建

参考`build.sh`

# 部署方式

* 命令行启动

`./uploader`

* systemd启动

```
/etc/systemd/system/uploader.service

[Unit]
Description="uploader"
After=network.target

[Service]
Type=simple

ExecStart=/data/server/uploader/uploader

LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

#### 使用方式

* 单文件上传

```
curl http://file.xxx.com/upload -F "file=@readme.md"
```

* 多文件上传

```
curl -X POST http://file.xxx.com/uploads -F "files=@test1.jpg" -F "files=@test2.png"

```

* 文件下载

```
wget http://file.xxx.com/readme.md
```


#### 定时删除历史文件

crontab -e

```
0 19 * * * find /data/data/uploader  -type f -mtime +14|xargs rm -rf  >/dev/null 2>&1
```
