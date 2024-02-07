package superadmin

type ProviderType string

const (
	CMCCProvider ProviderType = "CMCC"
)

type SuperAdmin struct {
	Mac  string `json:"mac,omitempty"`
	Name string `json:"name,omitempty"`
	Pwd  string `json:"pwd,omitempty"`
}
