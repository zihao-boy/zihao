package dns

import (
	"errors"
	dnsMap2 "github.com/zihao-boy/zihao/entity/dto/dnsMap"
	"golang.org/x/net/dns/dnsmessage"
	"log"
	"net"
	"strings"
)

const (
	// DNS server default port
	udpPort int = 53
	// DNS packet max length
	packetLen int = 512
)

var dnsMap map[string]*dnsMap2.DnsMapDto
var dnsIp string

type DnsServer struct {
	conn       *net.UDPConn
	memo       addrBag
	forwarders []net.UDPAddr
}

// Packet carries DNS packet payload and sender address.
type Packet struct {
	addr    net.UDPAddr
	message dnsmessage.Message
}

// Listen starts a DNS server on port 53
func (s *DnsServer) Listen() {
	var err error
	s.conn, err = net.ListenUDP("udp", &net.UDPAddr{Port: udpPort})
	if err != nil {
		log.Fatal(err)
	}
	defer s.conn.Close()

	for {
		buf := make([]byte, packetLen)
		_, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		var m dnsmessage.Message
		err = m.Unpack(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(m.Questions) == 0 {
			continue
		}
		go s.Query(Packet{*addr, m})
	}
}

// Query lookup answers for DNS message.
func (s *DnsServer) Query(p Packet) {
	// got response from forwarder, send it back to client
	if p.message.Header.Response {
		pKey := pString(p)
		if addrs, ok := s.memo.get(pKey); ok {
			for _, addr := range addrs {
				go sendPacket(s.conn, p.message, addr)
			}
			s.memo.remove(pKey)
		}
		return
	}

	// was checked before entering this routine
	q := p.message.Questions[0]

	// answer the question
	//val := dnsMap[qString(q)]

	val, err := s.GetDnsMap(qString(q))

	if err == nil {
		p.message.Answers = append(p.message.Answers, *val)
		go sendPacket(s.conn, p.message, p.addr)
	} else {
		// forwarding
		for i := 0; i < len(s.forwarders); i++ {
			s.memo.set(pString(p), p.addr)
			go sendPacket(s.conn, p.message, s.forwarders[i])
		}
	}
}

func (s *DnsServer) GetDnsMap(hostName string) (*dnsmessage.Resource, error) {
	dnsMapDto := dnsMap[hostName]

	if dnsMapDto != nil{
		resource,err := toResource(hostName,*dnsMapDto)
		return &resource,err
	}

	// add like
	for k, v := range dnsMap {
		if !strings.HasPrefix(k,"*."){
			continue
		}
		kPos := strings.Index(k,"*.")
		k = k[kPos:]
		if k == hostName{
			resource,err := toResource(hostName,*v)
			return &resource,err
		}
		hostNamePos := strings.Index(hostName,".")
		tmpHostName := hostName[hostNamePos:]
		if k == tmpHostName{
			resource,err := toResource(hostName,*v)
			return &resource,err
		}
	}
	return nil, errors.New("不存在")
}

/**
refresh config
 */
func FreshDnsConfig(dnsDataDto dnsMap2.DnsDataDto) {
	dnsIp = dnsDataDto.DnsIp
	if dnsDataDto.Maps == nil || len(dnsDataDto.Maps)<1{
		return ;
	}
	dnsMap = map[string]*dnsMap2.DnsMapDto{}
	for _,tmpDnsMap := range dnsDataDto.Maps{
		dnsMap[tmpDnsMap.Host] = tmpDnsMap
	}
}

func sendPacket(conn *net.UDPConn, message dnsmessage.Message, addr net.UDPAddr) {
	packed, err := message.Pack()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn.WriteToUDP(packed, &addr)
	if err != nil {
		log.Println(err)
	}
}

func toResource(host string, dnsMapDto dnsMap2.DnsMapDto) (dnsmessage.Resource, error) {
	rName, err := dnsmessage.NewName(host)
	none := dnsmessage.Resource{}
	if err != nil {
		return none, err
	}

	var rType dnsmessage.Type
	var rBody dnsmessage.ResourceBody

	rType = dnsmessage.TypeA
	ip := net.ParseIP(dnsMapDto.Value)
	if ip == nil {
		return none, errors.New("invalid IP address")
	}
	rBody = &dnsmessage.AResource{A: [4]byte{ip[12], ip[13], ip[14], ip[15]}}
	return dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  rName,
			Type:  rType,
			Class: dnsmessage.ClassINET,
			TTL:   600,
		},
		Body: rBody,
	}, nil
}
