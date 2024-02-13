
中文 | [English](./README.md)

### kuafu (夸父)
这是一个获取路由器超级管理员密码的工具

---
#### 安装

```shell
go get github.com/uaxe/kuafu
```

命令行帮助:
```shell
kuafu -help
```

命令行帮助输出:
```shell
Kuafu CLI 0.0.1

Available commands:

   version   The Kuafu CLI version
   modem     The Kuafu CLI modem

Flags:

  -help
    	Get help on the 'kuafu' command.
```
#### 获取光猫超管密码

```shell
kuafu modem admin
```

获取光猫超管密码输出:
```json
{
 "addr": "192.168.1.1",
 "device": "F607Za",
 "admin_name": "cuadmin",
 "admin_pwd": "cuadmin",
 "telnet_name": "root",
 "telnet_pwd": "Zte521"
}
```
查看参数信息
```shell
kuafu modem admin -help
```

查看参数信息输出:
```shell
Kuafu CLI 0.0.1

Kuafu modem admin - The Kuafu modem admin
Flags:

  -help
    	Get help on the 'kuafu modem admin' command.
  -host string
    	telnet host (default "192.168.1.1") 路由器IP
  -maddr string
    	mac addr   指定路由器Mac地址
  -port int
    	telnet port (default 23)  telnet默认的端口
  -type string
    	device type (default "CMCC") 设备类型，目前默认支持CMCC（移动）
```
