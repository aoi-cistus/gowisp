package main
import (
	"encoding/binary"
	"fmt"
	"log"
)

import "github.com/aoi-cistus/gowisp/packet"

const (
	CONNECT uint8 = 0x01
)

func main() {
	// 例: バイナリデータを作成 (パケットタイプ: 0x01, ストリームID: 12345, ペイロード: "Hello")
	packetType := CONNECT
	streamID := uint32(12345)
	payload := []byte("Hello")

	// バイナリデータを構築
	data := make([]byte, 5+len(payload))
	data[0] = packetType
	binary.LittleEndian.PutUint32(data[1:5], streamID)
	copy(data[5:], payload)

	// パケットをパース
	parsedPacket, err := packet.ParsePacket(data)
	if err != nil {
		log.Fatalf("Error parsing packet: %v", err)
	}

	// パケットの情報を表示
	fmt.Printf("Parsed Packet: %+v\n", parsedPacket)
	fmt.Printf("Packet Type Description: %s\n", packet.GetPacketTypeDescription(parsedPacket.PacketType))
}
