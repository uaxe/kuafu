package cmcc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/uaxe/infra/cmd"
	"github.com/uaxe/kuafu/provider/modem"
)

func (s *CMCCProvider) GetSuperAdmin() (*modem.SuperAdmin, error) {
	ret := modem.SuperAdmin{}
	macAddr, err := s.getMacAddr()
	if err != nil {
		return nil, err
	}
	ret.MacAddr = macAddr
	if s.macAddr != "" && strings.Contains(strings.ToUpper(macAddr), strings.ToUpper(s.macAddr)) {
		return nil, fmt.Errorf("mac addr no match %s %s", s.macAddr, macAddr)
	}
	telnetMacAddr := ""
	for _, part := range strings.Split(macAddr, ":") {
		if len(part) == 1 {
			part = fmt.Sprintf("0%v", part)
		}
		telnetMacAddr += strings.ToUpper(part)
	}

	if err = s.telnetEnable(telnetMacAddr); err != nil {
		return nil, err
	}

	if err = s.telnetAdminAndPwd(telnetMacAddr, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

var (
	RegIP = regexp.MustCompile(`(\d{1,3}\.){3}\d{1,3}`)

	RegMAC = regexp.MustCompile(`([\d|\w]{1,3}:){5}[\d|\w]{1,3}`)
)

func (s *CMCCProvider) getMacAddr() (string, error) {
	output, err := cmd.NewExecCmd(s.options.ctx, cmd.WithName("arp"), cmd.WithArgs("-a")).CombinedOutput()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(bytes.NewReader(output))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		value := scanner.Text()
		if ip := RegIP.FindString(value); ip == s.host {
			return RegMAC.FindString(value), nil
		}
	}

	return "", fmt.Errorf("not found host %s", s.host)
}

func (s *CMCCProvider) telnetEnable(macAddr string) error {
	uri := fmt.Sprintf(s.telnetURL, s.host, macAddr)
	resp, err := s.hclient.Get(uri)
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
		return fmt.Errorf("%s %s", uri, string(raw))
	}

	return nil
}

func (s *CMCCProvider) telnetAdminAndPwd(macAddr string, ret *modem.SuperAdmin) error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	if len(macAddr) < 6 {
		return fmt.Errorf("mac addr not less size %s", macAddr)
	}

	ret.TelnetName = "admin"
	ret.TelnetPwd = fmt.Sprintf("Fh@%s", macAddr[len(macAddr)-6:])

	var buf [4096]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		return err
	}

	buf[1] = 252
	buf[4] = 252
	buf[7] = 252
	buf[10] = 252
	_, err = conn.Write(buf[0:n])
	if err != nil {
		return err
	}

	n, err = conn.Read(buf[0:])
	if err != nil {
		return err
	}

	if !strings.Contains(string(buf[0:]), "Login:") {
		buf[1] = 252
		buf[4] = 251
		buf[7] = 252
		buf[10] = 254
		buf[13] = 252

		_, err = conn.Write(buf[0:n])
		if err != nil {
			return err
		}

		n, err = conn.Read(buf[0:])
		if err != nil {
			return err
		}

		buf[1] = 252
		buf[4] = 252
		_, err = conn.Write(buf[0:n])
		if err != nil {
			return err
		}

		_, err = conn.Read(buf[0:])
		if err != nil {
			return err
		}
	}

	_, err = conn.Write([]byte(ret.TelnetName + "\n"))
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)
	_, err = conn.Read(buf[0:])
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(ret.TelnetPwd + "\n"))
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)
	_, err = conn.Read(buf[0:])
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte("load_cli factory\n"))
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)
	_, err = conn.Read(buf[0:])
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte("show admin_name\n"))
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)
	_, err = conn.Read(buf[0:])
	if err != nil {
		return err
	}
	unames := RegAdminName.FindStringSubmatch(string(buf[0:n]))
	if len(unames) > 0 {
		ret.AdminName = unames[1]
	}

	_, err = conn.Write([]byte("show admin_pwd\n"))
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)
	_, err = conn.Read(buf[0:])
	if err != nil {
		return err
	}
	pwds := RegAdminPwd.FindStringSubmatch(string(buf[0:n]))
	if len(pwds) > 0 {
		ret.AdminPwd = pwds[1]
	}

	return nil
}

var (
	RegAdminName = regexp.MustCompile(`admin_name=(.*)`)

	RegAdminPwd = regexp.MustCompile(`admin_pwd=(.*)`)
)

const (
	TIME_DELAY_AFTER_WRITE = 500
)
