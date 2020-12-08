package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

// 二进制封包格式
type Packet struct {
	Size uint16 // 包体大小
	Body []byte // 包体数据
}

// 从dataReader读取封包
func readPacket(dataReader io.Reader) (pkt Packet, err error) {

	// Size为uint16，占2字节
	var sizeBuffer = make([]byte, 2)

	// 持续读取Size直到读到为止
	_, err = io.ReadFull(dataReader, sizeBuffer)

	// 发生错误时返回
	if err != nil {
		return
	}

	// 使用bytes.Reader读取sizeBuffer中的数据
	sizeReader := bytes.NewReader(sizeBuffer)

	// 读取小端的uint16作为size
	err = binary.Read(sizeReader, binary.LittleEndian, &pkt.Size)

	if err != nil {
		return
	}

	// 分配包体大小
	pkt.Body = make([]byte, pkt.Size)

	// 读取包体数据
	_, err = io.ReadFull(dataReader, pkt.Body)

	return
}

// 将数据写入dataWriter
func writePacket(dataWriter io.Writer, data []byte) error {

	// 准备一个字节数组缓冲
	var buffer bytes.Buffer

	// 将Size写入缓冲
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}

	// 写入包体数据
	if _, err := buffer.Write(data); err != nil {
		return err
	}

	// 获得写入的完整数据
	out := buffer.Bytes()

	// 写入dataWriter
	if _, err := dataWriter.Write(out); err != nil {
		return err
	}

	return nil
}
