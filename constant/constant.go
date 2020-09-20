package constant

//QrisConstant define the rule of QR TLVEMV
var QrisConstant string

func init() {
	QrisConstant = `
		{
			"tag85" : {
				"desc" : "payload format indicator",
				"length" : "05",
				"val" : "85"
			},
			"tag61" : {
				"desc" : "application template",
				"length" : "",
				"val" : "61"
			},
				"tag4F" : {
					"desc" : "application definition file (ADF) Name",
					"length" : "07",
					"val" : "4F",
					"defaultHex" : "A0 00 00 06 02 20 20"
				},
				"tag50" : {
					"desc" : "application label",
					"length" : "",
					"val" : "50"
				},
				"tag57" : {
					"desc" : "track 2 equivalent data",
					"length" : "",
					"val" : "57"
				},
				"tag5A" : {
					"desc" : "application PAN",
					"length" : "0A",
					"val" : "5A",
					"defaultHex" : "93 60"
				},
				"tag5F20" : {
					"desc" : "cardholders name",
					"length" : "",
					"val" : "5F20"
				},
				"tag5F2D" : {
					"desc" : "language preference",
					"length" : "",
					"val" : "5F2D"
				},
				"tag5F50" : {
					"desc" : "issuer url",
					"length" : "",
					"val" : "5F50"
				},
				"tag9F08" : {
					"desc" : "application version number",
					"length" : "2",
					"val" : "9F08"
				},
				"tag9F19" : {
					"desc" : "token requester ID",
					"length" : "11",
					"val" : "9F19"
				},
				"tag9F24" : {
					"desc" : "payment account reference",
					"length" : "29",
					"val" : "9F24"
				},
				"tag9F25" : {
					"desc" : "last 4 digits of PAN",
					"length" : "02",
					"val" : "9F25"
				},
				"tag63" : {
					"desc" : "application specific transparent template",
					"length" : "",
					"val" : "63"
				},
					"tag9F26" : {
						"desc" : "application cryptogram (AC)",
						"length" : "08", 
						"val" : "9F26"
					},
					"tag9F27" : {
						"desc" : "cryptogram information data (CID)",
						"length" : "01",
						"val" : "9F27"
					},
					"tag9F10" : {
						"desc" : "issuer application data (IAD)",
						"length" : "",
						"val" : "9F10"
					},
					"tag9F36" : {
						"desc" : "application transaction counter (ATC)",
						"length" : "02",
						"val" : "9F36"
					},
					"tag82" : {
						"desc" : "application interchange profile (AIP)",
						"length" : "02",
						"val" : "82"
					},
					"tag9F37" : {
						"desc" : "unpredictable number",
						"length" : "04",
						"val" : "9F37"
					},
					"tag9F74" : {
						"desc" : "indonesia qris",
						"length" : "",
						"val" : "9F74"
					},
		
			"status_qrcpm_inquiried" : 1,
			"status_qrcpm_flagged" : 2,
			"status_qrcpm_reversed" : 3,
			"trx_type_qrcpm" : "PurchaseQRCPM",
			"trx_type_reversal_qrcpm" : "ReversalPurchaseQRCPM",
			"trx_sch_qrcpm" : "SGRA",
			"third_party_qrcpm" : "JALIN",
			"account_type_qrcpm" : "Britama / IDR"
		}
	`
}
