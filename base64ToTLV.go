package base64totlvrefine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"brimo.bri.co.id/base64ToTLVRefine/constant"
)

//GetTLV is a function to conver base64 QR into TLV EMV in QRIS rule
func GetTLV(base64QR string) map[string]string {

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println(err)
		}
	}()

	constant := constant.QrisConstant
	var constantRule map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(constant), &constantRule)

	runningIndex := 0
	tagMap := make(map[string]string)
	for {
		tagLabel := strings.ToUpper(base64QR[runningIndex : runningIndex+2])

		if tagLabel == "5F" || tagLabel == "9F" {
			tagLabel = strings.ToUpper(base64QR[runningIndex : runningIndex+4])
		}

		if constantRule["tag"+tagLabel] == nil {
			fmt.Println("unknown tag : " + tagLabel)
			tagMap[tagLabel] = "tag not found"
			return tagMap
		}

		tagRule := constantRule["tag"+tagLabel].(map[string]interface{})

		tagLenExpected := ""
		if tagLabel[0:2] == "5F" || tagLabel[0:2] == "9F" {
			tagLenExpected = strings.ToUpper(base64QR[runningIndex+4 : runningIndex+6])
		} else {
			tagLenExpected = strings.ToUpper(base64QR[runningIndex+2 : runningIndex+4])
		}

		tagLenDecimalExpected, _ := strconv.ParseUint(hexaNumberToInteger("0x"+tagLenExpected), 16, 64)

		tagValueStartIndex := 0
		if tagLabel[0:2] == "5F" || tagLabel[0:2] == "9F" {
			tagValueStartIndex = runningIndex + 6
		} else {
			tagValueStartIndex = runningIndex + 4
		}

		tagValueEndIndex := tagValueStartIndex + int(tagLenDecimalExpected*2)
		tagValue := base64QR[tagValueStartIndex:tagValueEndIndex]
		isRulelengthConsidered := false

		if tagRule["length"] != "" {
			isRulelengthConsidered = true
		}

		if isRulelengthConsidered == true && tagLenExpected != tagRule["length"] {
			tagMap[tagLabel] = "invalid length " + tagLabel
			return tagMap
		}

		if tagLabel == "61" || tagLabel == "63" {
			if tagValueEndIndex != len(base64QR) {

				fmt.Println("invalid 61 " + strconv.Itoa(tagValueEndIndex) + " " + strconv.Itoa(len(base64QR)))
				tagMap[tagLabel] = "invalid 61"
				return tagMap
			}
			tagValueEndIndex = runningIndex + 4
		}
		runningIndex = tagValueEndIndex
		tagMap[tagLabel] = tagValue

		if runningIndex == len(base64QR) {
			break
		}
	}

	return tagMap
}

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
