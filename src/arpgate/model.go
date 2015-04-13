package arpgate

type Configuration struct {
	Release         string  `yaml:"release"`
	Created         string  `yaml:"created"`
	ConfigVersion   int     `yaml:"configversion"`
	SoftwareVersion int     `yaml:"softwareversion"`
	Device          Device  `yaml:"device"`
	Dhcp            Dhcp    `yaml:"dhcp"`
	Dns             Dns     `yaml:"dns"`
	Vpn             Vpn     `yaml:"vpn"`
	Ntp             Ntp     `yaml:"ntp"`
	Haproxy         Haproxy `yaml:"haproxy"`
	Nginx           Nginx   `yaml:"nginx"`
	Mqtt            Mqtt    `yaml:"mqtt"`
	Snort           Snort   `yaml:"snort"`
}

type Device struct {
	Pwd        string `yaml:"pwd"`
	Ip         string `yaml:"ip"`
	Subnet     string `yaml:"subnet"`
	Gateway    string `yaml:"gateway"`
	ExternalIp string `yaml:"externalIp"`
	Email      string `yaml:"email"`
}

type Dhcp struct {
	Enabled        bool   `yaml:"enabled"`
	Pxebootenabled bool   `yaml:"pxebootenabled"`
	Rangelow       string `yaml:"subnet"`
	Rangehigh      string `yaml:"gateway"`
}

type DnsRecord struct {
	Type  string `yaml:"type"`
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Dns struct {
	Enabled    bool        `yaml:"enabled"`
	ForwarderA string      `yaml:"forwardera"`
	ForwarderB string      `yaml:"forwarderb"`
	DnsRecords []DnsRecord `yaml:"dnsrecords"`
}

type Vpn struct {
	Enabled      bool   `yaml:"enabled"`
	Sharedsecret string `yaml:"sharedsecret"`
	Username     string `yaml:"username"`
	Pwd          string `yaml:"pwd"`
}

type Ntp struct {
	Enabled bool   `yaml:"enabled"`
	ServerA string `yaml:"servera"`
	ServerB string `yaml:"serverb"`
}

type Node struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type Forwarding struct {
	Port  int    `yaml:"port"`
	Nodes []Node `yaml:"nodes"`
}

type Haproxy struct {
	Enabled     bool         `yaml:"enabled"`
	Forwardings []Forwarding `yaml:"forwardings"`
}

type Site struct {
	Domain string `yaml:"domain"`
	Path   string `yaml:"path"`
	Port   int    `yaml:"port"`
}

type Nginx struct {
	Enabled bool   `yaml:"enabled"`
	Sites   []Site `yaml:"sites"`
}

type Mqtt struct {
	Enabled bool `yaml:"enabled"`
}

type Snort struct {
	Enabled bool `yaml:"enabled"`
}
