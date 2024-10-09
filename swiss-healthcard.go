package main

import (
	"fmt"
	"strings"

	"github.com/ebfe/scard"
)

type HealthcardIDAnswer struct {
	Error     string `json:"error"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthdate string `json:"birthDate"`
	AhvNr     string `json:"ahvNr"`
	Gender    string `json:"gender"`
}

// APDU Specs
// See https://www.bag.admin.ch/bag/de/home/versicherungen/krankenversicherung/krankenversicherung-versicherte-mit-wohnsitz-in-der-schweiz/versichertenkarte.html
const (
	NAME_TAG  = 128
	BIRTHDATE = 130
	AHV_NR    = 131
	GENDER    = 132
)

var (
	SELECT_MAIN_FILE = []byte{0x00, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00}
	SELECT_ID        = []byte{0x00, 0xA4, 0x02, 0x00, 0x02, 0x2F, 0x06}
	READ_84_BYTES    = []byte{0x00, 0xB0, 0x00, 0x00, 0x54}
)

// End APDU Specs
func GetHealthcardData() HealthcardIDAnswer {

	if CardReader == nil {
		return HealthcardIDAnswer{
			Error: "no card reader found",
		}
	}

	// Connect to the card
	Card, err := Context.Connect(*CardReader, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		return HealthcardIDAnswer{
			Error: "no smartcard found",
		}
	}

	// APDU Commands
	var selector = [][]byte{
		SELECT_MAIN_FILE,
		SELECT_ID,
		READ_84_BYTES,
	}

	selectRsp, err := APDUGetRsp(Card, selector, false)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return HealthcardIDAnswer{
			Error: err.Error(),
		}
	}

	return ParseHealthcardResponse(selectRsp)

}

func ParseHealthcardResponse(response []byte) HealthcardIDAnswer {
	var lastname string = ""
	var firstname string = ""
	var birthdate string = ""
	var ahv_nr string = ""
	var gender string = ""
	i := 2 // First 2 tags are headers
	for i < len(response) {
		tag := response[i]
		length := int(response[i+1])
		value := response[i+2 : i+2+length]
		switch tag {
		case NAME_TAG:
			fullname := string(value)
			parts := strings.Split(fullname, ", ")
			lastname = parts[0]
			firstname = parts[1]
		case BIRTHDATE:
			birthdate = string(value)
		case AHV_NR:
			ahv_nr = string(value)
		case GENDER:
			gender = parseGender(value[0])
		}
		i += 2 + length
	}

	Error := ""
	return HealthcardIDAnswer{
		Error,
		firstname,
		lastname,
		birthdate,
		ahv_nr,
		gender,
	}
}

func parseGender(gender byte) string {
	if gender == 0 {
		return "female"
	}
	if gender == 1 {
		return "male"
	}
	return ""
}
