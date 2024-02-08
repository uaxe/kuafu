package modem

const (
	CMCCProvider string = "CMCC"
)

type SuperAdmin struct {
	MacAddr string `json:"mac_addr,omitempty"`
	Name    string `json:"name,omitempty"`
	Pwd     string `json:"pwd,omitempty"`
}
