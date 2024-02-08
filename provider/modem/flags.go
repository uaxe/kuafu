package modem

type AdminFlag struct {
	Type    string `yaml:"type" name:"type" description:"device type"`
	MacAddr string `yaml:"maddr" name:"maddr" description:"mac addr"`
	Host    string `yaml:"host" name:"host" description:"telnet host"`
	Port    int    `yaml:"port" name:"port" description:"telnet port"`
}

func (*AdminFlag) Default() *AdminFlag {
	return &AdminFlag{
		Type: CMCCProvider,
		Host: "192.168.1.1",
		Port: 23,
	}
}
