package main

import (
	"fmt"
	"log"

	"go.mercari.io/go-emv-code/mpm"
	"go.mercari.io/go-emv-code/tlv"
)

func makeQR() {
	c := mpm.Code{
		PayloadFormatIndicator:      "01",
		PointOfInitiationMethod:     mpm.PointOfInitiationMethodDynamic,
		MerchantCategoryCode:        "4111",
		TransactionCurrency:         "156",
		CountryCode:                 "CN",
		MerchantName:                "BEST TRANSPORT",
		MerchantCity:                "BEIJING",
		PostalCode:                  "",
		AdditionalDataFieldTemplate: "030412340603***0708A60086670902ME",
		UnreservedTemplates: []tlv.TLV{
			{Tag: "80", Length: "36", Value: "003239401ff0c21a4543a8ed5fbaa30ab02e"},
		},
	}

	buf, err := mpm.Encode(&c)
	if err != nil {
		log.Fatal(err)
	}

	// You can convert this into QR Image for scanning
	fmt.Println(string(buf))

	// And this function decodes, here one can pass the text
	// value of QR image to get information in above structure
	// for easy reading

	dst, err := mpm.Decode(buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", dst)

}
