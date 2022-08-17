package storage

import (
	"os"
	"saving_service/process"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

type File struct {
	name string
}

func NewFile(name string) *File {

	return &File{
		name: name,
	}
}

func (f *File) Write(packets []process.Packet) error {
	file, err := os.Create(f.name)
	if err != nil {
		return err
	}
	defer file.Close()

	w := pcapgo.NewWriter(file)
	w.WriteFileHeader(65535, layers.LinkTypeEthernet)

	for _, pack := range packets {
		packet := gopacket.NewPacket(pack.GetData(), layers.LayerTypeEthernet, gopacket.Default)

		packet.Metadata().CaptureInfo.Timestamp = pack.GetCI().TimeStamp
		packet.Metadata().CaptureInfo.CaptureLength = pack.GetCI().CaptureLength
		packet.Metadata().CaptureInfo.Length = pack.GetCI().Length
		packet.Metadata().CaptureInfo.InterfaceIndex = pack.GetCI().InterfaceIndex
		packet.Metadata().CaptureInfo.AncillaryData = pack.GetCI().AccalaryData

		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	}
	return nil
}
