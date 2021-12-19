package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	Version      int
	TypeId       int
	Value        int
	LengthTypeId int
	Length       int
	SubPackets   []Packet
}

func main() {
	b := readInput("input.txt")
	p := DecodeBinary(b)
	fmt.Printf("part one: %d\npart two: %d", p.VersionSum(), p.Value)
}

func (p *Packet) VersionSum() int {
	sum := p.Version
	for _, sub := range p.SubPackets {
		sum += sub.VersionSum()
	}
	return sum
}

func DecodeBinary(binary string) Packet {
	p := Packet{}
	p.Version, p.TypeId, binary = p.bitsToDec(binary[:3]), p.bitsToDec(binary[3:6]), binary[6:]
	if p.TypeId == 4 {
		p.Value = p.decodeLiteralPacket(binary)
	} else {
		p.LengthTypeId = p.bitsToDec(binary[0:1])
		binary = binary[1:]
		p.decodeOperatorPacket(binary)
	}
	switch p.TypeId {
	case 0:
		for _, sub := range p.SubPackets {
			p.Value += sub.Value
		}
	case 1:
		p.Value = 1
		for _, sub := range p.SubPackets {
			p.Value *= sub.Value
		}
	case 2:
		p.Value = p.SubPackets[0].Value
		for _, sub := range p.SubPackets {
			if sub.Value < p.Value {
				p.Value = sub.Value
			}
		}
	case 3:
		p.Value = p.SubPackets[0].Value
		for _, sub := range p.SubPackets {
			if sub.Value > p.Value {
				p.Value = sub.Value
			}
		}
	case 5:
		if p.SubPackets[0].Value > p.SubPackets[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case 6:
		if p.SubPackets[0].Value < p.SubPackets[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case 7:
		if p.SubPackets[0].Value == p.SubPackets[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	}
	return p
}

func (p *Packet) decodeOperatorPacket(binary string) []Packet {
	if p.LengthTypeId == 0 {
		length := p.bitsToDec(binary[:15])
		var totalLen int
		binary = binary[15:]
		for {
			subPacket := DecodeBinary(binary)
			p.SubPackets = append(p.SubPackets, subPacket)
			p.Length += subPacket.Length
			totalLen += subPacket.Length
			binary = binary[subPacket.Length:]
			if totalLen >= length {
				break
			}
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
	value := ""
	for {
		groupBits := binary[:5]
		value += groupBits[1:]
		binary = binary[5:]
		p.Length += 5
		if strings.HasPrefix(groupBits, "0") {
			break
		}
	}
	num, err := strconv.ParseInt(value, 2, 64)
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
