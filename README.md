## 说明
使用python写的一个小工具。
旨在帮助运维人员，维护多个服务器连接的快速连接方式。

## 安装说明
### 依赖

* ssh
* [sshpass](http://sourceforge.net/projects/sshpass/)

### 安装
修改goconfig文件，修改成自己的服务器连接方式。将其复制成当前用户根目录下的 `.goconfig` 文件

    cp goconfig ~/.goconfig

复制 `go` 程序文件到当前用户的 `bin` 目录下

    cp go ~/bin/

## 配置文件说明
配置中 `tag` 为自定义的连接方式的别名，可以直接根据 `tag` 快速连接
`host` 为 `use@host` 的方式，指向使用何用户名以及连接的目标服务器
当前支持三种服务器连接方式，依赖 `ssh` 以及 `sshpass`

### 直接连接
配置连接方式如下所示。

    {
        "tag": "direct_connect",
        "host": "user@example.com"
    },

### 使用密码进行连接
配置连接方式如下所示。

    {
        "tag": "ssh_use_passwd",
        "host": "user@example.com",
        "passwd": "123456"
    },

### 使用指定的rsa文件连接
配置连接方式如下所示。

    {
        "tag": "ssh_use_rsa",
        "host": "user@example.com",
        "rsa": "/home/user/.ssh/id_rsa_2"
    },

## 使用

* 直接敲 `go [tag|idx]` ，会根据配置的各连接方式的顺序 `idx` ，或者对应的每个连接方式的 `tag` ，可以直接连接
* 直接敲 `go` ，会列出各连接方式，供选择

    go ssh_user_rsa

    go 3

    go
