package collectd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_parsesTheHostname(t *testing.T) {

	fileBytes, err := ioutil.ReadFile("cpu_disk_packet.dat")

	if err != nil {
		fmt.Errorf("error encountered %v", err)
	}

	buffer := bytes.NewBuffer(fileBytes)

	var partType uint16
	var partLength uint16
	var content string

	err = binary.Read(buffer, binary.BigEndian, &partType)
	if err != nil {
		fmt.Errorf("error encountered %v", err)
	}

	err = binary.Read(buffer, binary.BigEndian, &partLength)
	if err != nil {
		fmt.Errorf("error encountered %v", err)
	}

	contentBytes := buffer.Next(int(partLength - 4))
	//Trim the null terminating byte from the string
	content = string(contentBytes[0 : len(contentBytes)-1])

	assert.Equal(t, content, "localhost", "contents does not equal localhost")
}