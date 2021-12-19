package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Message struct {
	RawHex      string
	Binary      string
	HexToBinary map[string]string
	Packets     []Packet
	VersionSum  int
	Length      int
}

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
	m := readInput("test.txt")
	m.DecodeBinary()
}

func (p *Packet) VersionSum() int {
	var sum int
	for _, p := range p.SubPackets {
		sum += p.Version
	}
	return sum
}

func (m *Message) DecodeBinary() []Packet {
	fmt.Println("start", m.Binary)
	if len(strings.TrimLeft(m.Binary, "0")) == 0 {
		return m.Packets
	}
	if strings.HasPrefix(m.Binary, "0000") {
		fmt.Println("why", m.Binary)
		m.Binary = m.Binary[4:]
		fmt.Println("why", m.Binary)
	}
	packet := Packet{}
	packet.Version, packet.TypeId, m.Binary = m.bitsToDec(m.Binary[:3]), m.bitsToDec(m.Binary[3:6]), m.Binary[6:]
	m.VersionSum += packet.Version
	if packet.TypeId == 4 {
		packet.Number = m.decodeLiteralPacket()
	} else {
		lengthTypeId := m.bitsToDec(m.Binary[0:1])
		m.Binary = m.Binary[1:]
		packet.Number = m.decodeOperatorPacket(lengthTypeId)
	}
	m.Packets = append(m.Packets, packet)

	fmt.Println("---", m.VersionSum, packet.Version)
	return m.DecodeBinary()

}

func (m *Message) decodeOperatorPacket(lengthTypeId int) int {
	number := ""
	if lengthTypeId == 0 {
		fmt.Println("here", m.Binary)
		length := m.bitsToDec(m.Binary[:15])
		m.Binary = m.Binary[15:]
		for {
			if length-11 < 0 {
				number += m.Binary[:length]
				m.Binary = m.Binary[(m.Length-(len(m.Binary)-length))%4+length:]
				m.Length = len(m.Binary)
				break
			}
			fmt.Println(m.Binary, length)
			number += m.Binary[:11]
			m.Binary = m.Binary[11:]
			length -= 11
		}
	} else if lengthTypeId == 1 {
		num := m.bitsToDec(m.Binary[:11])
		fmt.Println("now", num)
		m.Binary = m.Binary[11:]
		for {
			if num == 0 {
				m.Binary = m.Binary[(m.Length-len(m.Binary))%4:]
				m.Length = len(m.Binary)
				break
			}
			number += m.Binary[:11]
			m.Binary = m.Binary[11:]
			num -= 1
		}
	}
	fmt.Println("here", m.Binary)
	return m.bitsToDec(number)
}

func (m *Message) decodeLiteralPacket() int {
	number := ""
	for {
		number += m.Binary[1:5]
		m.Binary = m.Binary[5:]
		if m.Binary[0:1] == "0" {
			// number += m.Binary[1:5]
			// m.Binary = m.Binary[(m.Length-(len(m.Binary)-5))%4+5:]
			// m.Length = len(m.Binary)
			break
		}
	}
	return m.bitsToDec(number)
}

func (m *Message) bitsToDec(bits string) int {
	dec, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(dec)
}

func (m *Message) hexToBits() {
	for _, r := range m.RawHex {
		m.Binary += m.HexToBinary[string(r)]
	}
}

func readInput(fname string) Message {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := Message{}
	m.HexToBinary = make(map[string]string)
	first := true
	for scanner.Scan() {
		if first {
			m.RawHex = scanner.Text()
			first = false
		} else if scanner.Text() == "" {
			continue
		} else {
			row := strings.Split(scanner.Text(), " = ")
			m.HexToBinary[row[0]] = row[1]
		}
	}
	m.hexToBits()
	m.Length = len(m.Binary)
	return m
}
