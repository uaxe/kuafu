package cucc

import (
	"fmt"
	"net"
	"regexp"
	"time"

	"github.com/uaxe/kuafu/provider/modem"
)

func (s *CUCCProvider) GetSuperAdmin() (*modem.SuperAdmin, error) {
	ret := modem.SuperAdmin{Addr: s.host, Device: s.device}

	if err := s.telnetAdminAndPwd(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *CUCCProvider) telnetAdminAndPwd(ret *modem.SuperAdmin) error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	ret.TelnetName = "root"
	ret.TelnetPwd = "Zte521"

	var buf [4096]byte
	_, err = conn.Read(buf[0:])
	if err != nil {
		return err
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

	_, err = conn.Write([]byte("sendcmd 1 DB p DevAuthInfo\n"))
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)
	_, err = conn.Read(buf[0:])
	if err != nil {
		return err
	}

	unames := RegAdminName.FindStringSubmatch(string(buf[0:]))
	if len(unames) > 0 {
		ret.AdminName = unames[1]
	}

	pwds := RegAdminPwd.FindStringSubmatch(string(buf[0:]))
	if len(pwds) > 0 {
		ret.AdminPwd = pwds[1]
	}

	return nil
}

var (
	RegAdminName = regexp.MustCompile(`<DM name="User" val="(.*)"/>`)

	RegAdminPwd = regexp.MustCompile(`<DM name="Pass" val="(.*)"/>`)
)

const (
	TIME_DELAY_AFTER_WRITE = 500
)
