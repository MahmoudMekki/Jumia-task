package model

/*Phone model that holds phone number data for countries*/
type Phone struct {
	Country string `json:"country"`
	Regex   string `json:"-"`
	Code    string `json:"code"`
}

/*AvailableCodes for countries provided in the document*/
var AvailableCodes = map[string]Phone{
	"(237)": {Code: "(237)", Regex: `\(237\)\ ?[2368]\d{7,8}$`, Country: "Cameron"},
	"(251)": {Code: "(251)", Regex: `\(251\)\ ?[1-59]\d{8}$`, Country: "Ethiopia"},
	"(212)": {Code: "(212)", Regex: `\(212\)\ ?[5-9]\d{8}$`, Country: "MOROCCO"},
	"(258)": {Code: "(258)", Regex: `\(258\)\ ?[28]\d{7,8}$`, Country: "Cameron"},
	"(256)": {Code: "(256)", Regex: `\(256\)\ ?\d{9}$`, Country: "UGANDA"},
}
