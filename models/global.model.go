package models

type Response struct {
	Rc      string      `json:"rc"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

const (
	ERR_CODE_00     = "00"
	ERR_CODE_00_MSG = "SUCCESS.."

	ERR_CODE_01     = "01"
	ERR_CODE_01_MSG = "FAILED.."

	ERR_CODE_03     = "03"
	ERR_CODE_03_MSG = "Error, unmarshall body Request"

	ERR_CODE_05     = "05"
	ERR_CODE_05_MSG = "File Format not csv"

	ERR_CODE_07     = "07"
	ERR_CODE_07_MSG = "File Empty / Already Uploaded"

	ERR_CODE_08     = "08"
	ERR_CODE_08_MSG = "RRN Responder no Set"

	ERR_CODE_09     = "09"
	ERR_CODE_09_MSG = "IN PROGRESS.."

	ERR_CODE_10     = "10"
	ERR_CODE_10_MSG = "Data on File is Can't be Empty"

	ERR_CODE_11     = "11"
	ERR_CODE_11_MSG = "Internal Server Error"

	ERR_CODE_12     = "12"
	ERR_CODE_12_MSG = "Value must Greater than 0"

	ERR_CODE_13     = "13"
	ERR_CODE_13_MSG = "Value has wrong format"
)
