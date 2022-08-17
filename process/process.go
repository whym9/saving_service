package process

import "time"

type Capture struct {
	TimeStamp      time.Time     `json: "time"`
	CaptureLength  int           `json: "caplength"`
	Length         int           `json: "length"`
	InterfaceIndex int           `json :  "index"`
	AccalaryData   []interface{} `json: "accalary"`
}

type Packet struct {
	ci   Capture
	data []byte
}

func (p Packet) GetCI() Capture {
	return p.ci
}

func (p Packet) GetData() []byte {
	return p.data
}

func NewPacket(ci Capture, data []byte) Packet {
	p := Packet{}

	p.ci = ci
	p.data = data
	return p
}

type Protocols struct {
	TCP  int `json: "TCP"`
	UDP  int `json: "UDP"`
	IPv4 int `json: "IPv4"`
	IPv6 int `json: "IPv6"`
}
