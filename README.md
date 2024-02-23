
[![Build Status](https://github.com/uaxe/kuafu/workflows/GoTests/badge.svg?branch=master)](https://github.com/uaxe/kuafu/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/uaxe/kuafu/graph/badge.svg)](https://codecov.io/gh/uaxe/kuafu)
[![GoDoc](https://pkg.go.dev/badge/github.com/uaxe/kuafu?status.svg)](https://pkg.go.dev/github.com/uaxe/kuafu?tab=doc)


English | [中文](./README_CN.md)

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
You can view the parameters
```shell
kuafu modem admin -help
```

View the parameters output:
```shell
Kuafu CLI 0.0.1

Kuafu modem admin - The Kuafu modem admin
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
