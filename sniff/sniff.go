package sniff

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// func StartUtil(dev string) {
// if handle, err := pcap.OpenLive(dev, 1600, true, pcap.BlockForever); err != nil {
//   panic(err)
// } else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
//   panic(err)
// } else {
//   packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
//   for packet := range packetSource.Packets() {
//     handlePacket(packet)
//   }
// }
//
// }
//
// func handlePacket(packet gopacket.Packet) error {
//   fmt.Println(packet.String())
//   // buff <- packet.String()
//   return nil
// }

func StartUtilAndGetPacketChannel(dev string)(chan gopacket.Packet){
  if handle, err := pcap.OpenLive(dev, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		return packetSource.Packets() 
	}
}

func StartUtil(dev string, buff chan string) {
	if handle, err := pcap.OpenLive(dev, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil {
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet, buff)
		}
	}

}

func handlePacket(packet gopacket.Packet, buff chan string) {
	// fmt.Println(packet.String())
	buff <- packet.String()
}
