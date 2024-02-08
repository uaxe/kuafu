package cmcc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/uaxe/infra/cmd"
	"github.com/uaxe/kuafu/provider/modem"
)

func init() {
	if err := modem.Register(&CMCCProvider{
		providerType: modem.CMCCProvider}); err != nil {
		panic(err)
	}
}

var _ modem.Provider = (*CMCCProvider)(nil)

type CMCCProvider struct {
	providerType string
	options
}

type OptionFunc func(opt *options) error

type options struct {
	ctx       context.Context
	hclient   *http.Client
	host      string
	macAddr   string
	telnetURL string
	port      int
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

func WithMacAddr(addr string) OptionFunc {
	return func(opt *options) error {
		opt.macAddr = addr
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

func (s *CMCCProvider) New(ctx context.Context, f *modem.AdminFlag) (modem.Provider, error) {
	ss := &CMCCProvider{options: defaultOptions()}
	if ctx != nil {
		ss.ctx = ctx
	}
	if f.Host != "" {
		ss.host = f.Host
	}
	if f.MacAddr != "" {
		ss.macAddr = f.MacAddr
	}
	if f.Port > 0 {
		ss.port = f.Port
	}
	return ss, nil
}

func (s *CMCCProvider) Type() string {
	return s.providerType
}

func (s *CMCCProvider) GetSuperAdmin() (*modem.SuperAdmin, error) {
	macAddr, err := s.arp()
	if err != nil {
		return nil, err
	}

	ret := modem.SuperAdmin{MacAddr: macAddr}

	if err = s.telnetEnable(macAddr); err != nil {
		return nil, err
	}

	return &ret, nil
}

var (
	RegARP = regexp.MustCompile(`\((\d{1,3}\.){3}\d{1,3}\)\s{1,2}at\s{1,2}([\d|\w]{1,3}:){5}[\d|\w]{1,3}`)

	RegIP = regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}`)

	RegMAC = regexp.MustCompile(`([\d|\w]{1,3}:){5}[\d|\w]{1,3}`)
)

func (s *CMCCProvider) arp() (string, error) {
	output, err := cmd.NewExecCmd(s.options.ctx, cmd.WithName("arp"), cmd.WithArgs("-a")).CombinedOutput()
	if err != nil {
		return "", err
	}

	arps := RegARP.FindAllString(string(output), -1)
	for _, value := range arps {
		if ip := RegIP.FindString(value); ip == s.host {
			return RegMAC.FindString(value), nil
		}
	}

	return "", fmt.Errorf("not found host %s", s.host)
}

func (s *CMCCProvider) telnetEnable(macAddr string) error {
	resp, err := s.hclient.Get(fmt.Sprintf(s.telnetURL, s.host, macAddr))
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telnet start http not OK")
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if !strings.Contains(string(raw), "telnet") {
		return errors.New(string(raw))
	}

	return nil
}

func (s *CMCCProvider) telnetAdminAndPwd() error {

	return nil
}
