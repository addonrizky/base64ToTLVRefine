package base64ToTLVRefine

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"regexp"

	"github.com/addonrizky/base64ToTLVRefine/constant"
)

//GetTLV is a function to conver base64 QR into TLV EMV in QRIS rule
func GetTLV(base64QR string) (map[string]string, error) {

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println(err)
		}
	}()

	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("=+$")
	if err != nil {
		log.Fatal(err)
	}

	base64QRTrimmed := reg.ReplaceAllString(base64QR, "")

	decodedBase64, err := base64.RawStdEncoding.DecodeString(base64QRTrimmed)
	if err != nil {
		return nil, err
	}
	hexaSlice := hex.EncodeToString(decodedBase64)

	constant := constant.QrisConstant
	var constantRule map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(constant), &constantRule)

	runningIndex := 0
	//counter := 0
	tagMap := make(map[string]string)
	for {
		tagLabel := strings.ToUpper(hexaSlice[runningIndex : runningIndex+2])

		if tagLabel == "5F" || tagLabel == "9F" {
			tagLabel = strings.ToUpper(hexaSlice[runningIndex : runningIndex+4])
		}

		if constantRule["tag"+tagLabel] == nil {
			return nil, errors.New("unknown tag : " + tagLabel)
		}

		tagRule := constantRule["tag"+tagLabel].(map[string]interface{})

		tagLenExpected := ""
		if tagLabel[0:2] == "5F" || tagLabel[0:2] == "9F" {
			tagLenExpected = strings.ToUpper(hexaSlice[runningIndex+4 : runningIndex+6])
		} else {
			tagLenExpected = strings.ToUpper(hexaSlice[runningIndex+2 : runningIndex+4])
		}

		//issuer anomaly spec
		if tagLenExpected == "81" {
			runningIndex += 2
			if tagLabel[0:2] == "5F" || tagLabel[0:2] == "9F" {
				tagLenExpected = strings.ToUpper(hexaSlice[runningIndex+4 : runningIndex+6])
			} else {
				tagLenExpected = strings.ToUpper(hexaSlice[runningIndex+2 : runningIndex+4])
			}
		}

		tagLenDecimalExpected, _ := strconv.ParseUint(hexaNumberToInteger("0x"+tagLenExpected), 16, 64)

		tagValueStartIndex := 0
		if tagLabel[0:2] == "5F" || tagLabel[0:2] == "9F" {
			tagValueStartIndex = runningIndex + 6
		} else {
			tagValueStartIndex = runningIndex + 4
		}
		tagValueEndIndex := tagValueStartIndex + int(tagLenDecimalExpected*2)
		tagValue := hexaSlice[tagValueStartIndex:tagValueEndIndex]
		isRulelengthConsidered := false

		if tagRule["length"] != "" {
			isRulelengthConsidered = true
		}

		if isRulelengthConsidered == true && tagLenExpected != tagRule["length"] {
			return nil, errors.New("invalid length " + tagLabel)
		}

		if tagLabel == "61" || tagLabel == "63" {
			if tagValueEndIndex != len(hexaSlice) {
				return nil, errors.New("invalid length, expected and actual length not match for tag : " + tagLabel)
			}
			tagValueEndIndex = runningIndex + 4
		}
		runningIndex = tagValueEndIndex
		tagMap[tagLabel] = strings.ToUpper(tagValue)
		//counter++
		//
		//if counter == 2 {
		//	break
		//}

		if runningIndex == len(hexaSlice) {
			break
		}
	}

	tagMap["tlvemv"] = hexaSlice

	return tagMap, nil
}

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
