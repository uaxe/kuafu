
English | [中文](./README_cn.md)

### kuafu
This is a tool to obtain the router super administrator password

---
#### Install 

```shell
go get github.com/uaxe/kuafu
```

Commond help: 
```shell
kuafu -help
```

Commond help output:
```shell
Kuafu CLI 0.0.1

Available commands:

   version   The Kuafu CLI version
   modem     The Kuafu CLI modem

Flags:

  -help
    	Get help on the 'kuafu' command.
```
#### Get modem super administrator password

```shell
kuafu modem admin 
```

Commond modem admin output:
```shell
mac_addr: 7c:fc:fd:2:17:a0
admin_name: CMCCAdmin
admin_pwd: aDm8H%MdAD*5Vz2Hh
```
You can view the parameters
```shell
kuafu modem admin -help
```

View the parameters output:
```shell
Kuafu CLI 0.0.1

KuaFu modem admin - The Kuafu modem admin
Flags:

  -help
    	Get help on the 'kuafu modem admin' command.
  -host string
    	telnet host (default "192.168.1.1")
  -maddr string
    	mac addr
  -port int
    	telnet port (default 23)
  -type string
    	device type (default "CMCC")
```




