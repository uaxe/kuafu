package cmcc

import (
	"context"
	"fmt"
	"github.com/uaxe/infra/cmd"
	"github.com/uaxe/kuafu/provider/superadmin"
	"net/http"
	"regexp"
)

func init() {
	if err := superadmin.Register(&CMCCProvider{
		providerType: superadmin.CMCCProvider,
		options:      defaultOptions(),
	}); err != nil {
		panic(err)
	}
}

var _ superadmin.Provider = (*CMCCProvider)(nil)

type CMCCProvider struct {
	providerType string
	options
}

type OptionFunc func(opt *options) error

type options struct {
	ctx       context.Context
	hclient   *http.Client
	host      string
	port      int
	telnetURL string
}

func WithContext(ctx context.Context) OptionFunc {
	return func(opt *options) error {
		opt.ctx = ctx
		return nil
	}
}

func WithHclient(hclient *http.Client) OptionFunc {
	return func(opt *options) error {
		opt.hclient = hclient
		return nil
	}
}

func WithHost(host string) OptionFunc {
	return func(opt *options) error {
		opt.host = host
		return nil
	}
}

func WithTelnetURL(uri string) OptionFunc {
	return func(opt *options) error {
		opt.telnetURL = uri
		return nil
	}
}

func defaultOptions() options {
	return options{
		ctx:       context.Background(),
		hclient:   http.DefaultClient,
		host:      "192.168.1.1",
		port:      23,
		telnetURL: "http://%s/cgi-bin/telnetenable.cgi?telnetenable=1&key=%s",
	}
}

func (s *CMCCProvider) New(ctx context.Context, f *superadmin.AdminFlag) error {
	s = &CMCCProvider{options: defaultOptions()}
	if ctx != nil {
		s.ctx = ctx
	}
	if f.Host != "" {
		s.host = f.Host
	}

	if f.Port > 0 {
		s.port = f.Port
	}

	return nil
}

func (s *CMCCProvider) Type() string {
	return s.providerType
}

func (s *CMCCProvider) GetSuperAdmin() (ret *superadmin.SuperAdmin, e error) {
	_, err := s.arp()
	if err != nil {
		return nil, err
	}

	return
}

var (
	RegARP = regexp.MustCompile(`\((\d{1,3}\.){3}\d{1,3}\)\s{1,2}at\s{1,2}([\d|\w]{1,3}:){5}[\d|\w]{1,3}`)
	RegIP  = regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}`)
	RegMAC = regexp.MustCompile(`(\d{1,3}:){5}\d{1,3}`)
)

func (s *CMCCProvider) arp() (string, error) {
	output, err := cmd.NewExecCmd(
		s.options.ctx,
		cmd.WithName("arp"),
		cmd.WithArgs("-a"),
	).CombinedOutput()
	if err != nil {
		return "", err
	}
	arps := RegARP.FindAllString(string(output), 100)
	for _, value := range arps {
		if ip := RegIP.FindString(value); ip == s.host {
			return RegMAC.FindString(value), nil
		}
	}
	return "", fmt.Errorf("not found host %s", s.host)
}
