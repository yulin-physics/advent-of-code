package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// type Message struct {
// 	RawHex      string
// 	Binary      string
// 	HexToBinary map[string]string
// 	Packets     []Packet
// 	VersionSum  int
// 	Length      int
// }

//type id = 4: literal value packet, encodes a binary number, len % 4 ==0
//type id != 4: operator packet
type Packet struct {
	Version      int
	TypeId       int
	Number       int
	LengthTypeId int
	Length       int
	SubPackets   []Packet
}

func main() {
	b := readInput("test.txt")
	p := DecodeBinary(b)
	fmt.Println(p.VersionSum())

}

func (p *Packet) VersionSum() int {
	var sum int
	for {
		sum += p.Version
		if p.SubPackets == nil {
			break
		}
		p = &p.SubPackets[0]
	}
	return sum
}

func DecodeBinary(binary string) Packet {
	p := Packet{}
	fmt.Println(binary)
	p.Version, p.TypeId, binary = p.bitsToDec(binary[:3]), p.bitsToDec(binary[3:6]), binary[6:]
	if p.TypeId == 4 {
		p.Number = p.decodeLiteralPacket(binary)
	} else {
		p.LengthTypeId = p.bitsToDec(binary[0:1])
		binary = binary[1:]
		p.decodeOperatorPacket(binary)
	}
	return p
}

func (p *Packet) decodeOperatorPacket(binary string)[]Packet {
	if p.LengthTypeId == 0 {
		length := p.bitsToDec(binary[:15])
		binary = binary[15:]
		for {
			if length-11 < 0 {
				break
			}
			subPacket := DecodeBinary(binary)
			p.SubPackets = append(p.SubPackets, subPacket)
			p.Length += subPacket.Length
			binary = binary[subPacket.Length:]
			length -= 11
		}
	} else if p.LengthTypeId == 1 {
		num := p.bitsToDec(binary[:11])
		binary = binary[11:]
		for range make([]struct{}, num) {
			subPacket := DecodeBinary(binary)
			p.Length += subPacket.Length
			p.SubPackets = append(p.SubPackets, subPacket)
			binary = binary[subPacket.Length:]
		}
	}
	return nil
}

func (p *Packet) decodeLiteralPacket(binary string) int {
	number := ""
	for {
		groupBits := binary[:5]
		number += groupBits[1:]
		binary = binary[5:]
		p.Length += 5
		if strings.HasPrefix(groupBits, "0") {
			break
		}
	}
	num, err := strconv.ParseInt(number, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(num)
}

func (p *Packet) bitsToDec(bits string) int {
	dec, err := strconv.ParseInt(bits, 2, 64)
	p.Length += len(bits)
	if err != nil {
		log.Fatal(err)
	}
	return int(dec)
}

func hexToBits(rawHex string, hexToBinary map[string]string) string {
	binary := ""
	for _, r := range rawHex {
		binary += hexToBinary[string(r)]
	}
	return binary
}

func readInput(fname string) string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rawHex := ""
	hexToBinary := make(map[string]string)
	first := true
	for scanner.Scan() {
		if first {
			rawHex = scanner.Text()
			first = false
		} else if scanner.Text() == "" {
			continue
		} else {
			row := strings.Split(scanner.Text(), " = ")
			hexToBinary[row[0]] = row[1]
		}
	}
	binary := hexToBits(rawHex, hexToBinary)
	return binary
}
