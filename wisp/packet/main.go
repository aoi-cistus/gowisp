package packet

import (
	"encoding/binary"
	"fmt"
)

const (
	CONNECT uint8 = 0x01
	DATA uint8 = 0x02
	CONTINUE uint8 = 0x03
	CLOSE uint8 = 0x04
)

type Packet struct {
	PacketType uint8
	StreamID   uint32
	Payload    []byte
}

func getPacketTypeDescription(packetType uint8) int {
	switch packetType {
	case CONNECT:
		return 1
	case DATA:
		return 2
	case CONTINUE:
		return 3
	case CLOSE:
		return 4
	default:
		return 0
	}
}

func parseCloseReason(payload []byte) (uint8, error) {
	if len(payload) < 1 {
		return 0, fmt.Errorf("payload too short")
	}
	return payload[0], nil
}

func parsePacket(data []byte) (*Packet, error) {
	if len(data) < 5 {
		return nil, fmt.Errorf("data too short")
	}

	packet := &Packet{
		PacketType: data[0],
		StreamID:   binary.LittleEndian.Uint32(data[1:5]),
		Payload:    data[5:],
	}

	return packet, nil
}
