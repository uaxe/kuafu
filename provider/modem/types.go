package modem

const (
	CMCCProvider string = "cmcc"
	CUCCProvider string = "cucc"
	CTCCProvider string = "ctcc"
)

type SuperAdmin struct {
	Addr       string `json:"addr,omitempty"`
	Device     string `json:"device,omitempty"`
	MacAddr    string `json:"mac_addr,omitempty"`
	AdminName  string `json:"admin_name,omitempty"`
	AdminPwd   string `json:"admin_pwd,omitempty"`
	TelnetName string `json:"telnet_name,omitempty"`
	TelnetPwd  string `json:"telnet_pwd,omitempty"`
}
