package helper

import (
	"fmt"
	"strings"
)

type SublinkHelper struct {
	LinkSettings []Link
}

func (h *SublinkHelper) GetConfigs() string {
	var links []string
	for _, cfg := range h.LinkSettings {
		links = append(links, cfg.getConfig())
	}
	return strings.Join(links, "\n")
}

func (h *SublinkHelper) GetClashConfigs() string {
	clashTemplate := `
	mixed-port: 7890
	mode: rule
	ipv6: true
	proxies:
	@PROXIES@
	proxy-groups:
    	- name: "Fish VPN"
      	type: select
      	proxies:
	@LIST@
	rules:
    	- 'DOMAIN-SUFFIX,.local,DIRECT'
    	# - GEOIP,IR,DIRECT
    	- IP-CIDR,127.0.0.0/8,DIRECT
    	- 'MATCH,VPN'
  `

	var confs []string
	for _, conf := range h.LinkSettings {
		confs = append(confs, conf.getConfigClash())
	}
	res := strings.Replace(clashTemplate, "@PROXIES@", strings.Join(confs, "\n"), 1)

	var selectorList []string
	for _, conf := range h.LinkSettings {
		selectorList = append(selectorList, fmt.Sprintf(`				- %s`, conf.getName()))
	}
	res = strings.Replace(res, "@LIST@", strings.Join(selectorList, "\n"), 1)

	return res
}

type Link interface {
	getConfig() string
	getConfigClash() string
	getName() string
}

type TrojanLink struct {
	Remark string
	Addr   string
	Port   int
	Passwd string
	Path   string
	SNI    string
	Host   string
}

func (s *TrojanLink) getName() string {
	return s.Remark
}

func (s *TrojanLink) getConfig() string {
	return fmt.Sprintf(
		`trojan://%s@%s:%d?security=tls&alpn=http/1.1&host=%s&fp=chrome&type=ws&path=%s&sni=%s#%s`,
		s.Passwd,
		s.Addr,
		s.Port,
		s.Host,
		s.Path,
		s.SNI,
		s.Remark,
	)
}

func (s *TrojanLink) getConfigClash() string {
	return fmt.Sprintf(`
	- name: "%s"
	  type: trojan
	  port: %d
	  network: ws
	  ws-opts:
	      path: %s
	      headers:
	        Host: %s
	  ws-headers:
	      Host: %s
	  password: %s
	  udp: true
	  sni: %s
	`,
		s.Remark,
		s.Port,
		s.Path,
		s.Host,
		s.Host,
		s.Passwd,
		s.SNI,
	)
}
