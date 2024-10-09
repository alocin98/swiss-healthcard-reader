package main

import (
	"fmt"
	"time"

	"github.com/ebfe/scard"
)

type CardReaderResponse struct {
	isReaderPresent bool   `json: "isReaderPresent"`
	reader          string `json: "reader"`
}

// Transmitter is an interface that wrap the command to communicate with smart card via application protocol data unit (APDU) according to ISO/IEC 7816.
type Transmitter interface {
	Transmit([]byte) ([]byte, error)
}

// Check for card reader if we have none.
func ConnectCardReader() {
	for {
		time.Sleep(1 * time.Second)
		var err error
		// Establish a PC/SC context
		Context, err = scard.EstablishContext()
		if err != nil {
			fmt.Println("Error EstablishContext:", err)
			return
		}

		// List available readers
		Readers, err := Context.ListReaders()
		if err != nil {
			fmt.Println("Error ListReaders:", err)
			CardReader = nil
			continue
		}

		if CardReader == nil {
			CardReader = &Readers[0]
			fmt.Println("Card reader connected: ", *CardReader)
			continue
		}

	}
}

func GetCardReader() CardReaderResponse {
	if CardReader == nil {
		return CardReaderResponse{
			isReaderPresent: false,
			reader:          "",
		}
	} else {
		return CardReaderResponse{
			isReaderPresent: true,
			reader:          *CardReader,
		}
	}
}

// APDUGetRsp Send list of APDU and get last command response
// ispadzeroOptional is optional(default = true) to replace adpu tail section
func APDUGetRsp(card Transmitter, apducmds [][]byte, ispadzeroOptional ...bool) ([]byte, error) {
	var resp []byte

	ispadzero := true
	if len(ispadzeroOptional) > 0 {
		ispadzero = ispadzeroOptional[0]
	}

	// Send command APDU: apducmds
	for _, apducmd := range apducmds {
		rsp, err := card.Transmit(apducmd)
		if err != nil {
			fmt.Println("Error Transmit:", err)
			return nil, err
		}
		//printRsp(rsp)
		resp = rsp
	}

	// pad zero
	if ispadzero == true {
		dlen := len(resp)
		resp[dlen-2] = 0
	}

	return resp, nil
}
