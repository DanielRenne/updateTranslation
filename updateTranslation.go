package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/DanielRenne/go-yandex-translate"
)

type yandexFile struct {
	Key       string   `json:"key"`
	Path      string   `json:"path"`
	Languages []string `json:"languages"`
}

func main() {

	var settings yandexFile
	err := readFileAndParse("yandex.json", &settings)
	if err != nil {
		fmt.Println("Error reading yandex.json.  " + err.Error())
		fmt.Println("A new yandex.json file has been created for you to add your API Key.  Please update.")
		var yf yandexFile
		yf.Languages = append(yf.Languages, "es")
		yf.Path = "web/pages"
		parseAndWriteFile("yandex.json", yf, 0755)
		return
	}

	tr := yandex_translate.New(settings.Key)

	page := os.Args[1]

	var translationsFile i18nFile

	//Read the existing json file if it exists.
	if doesFileExist(settings.Path + "/" + page + "/i18n.json") {
		err = readFileAndParse(settings.Path+"/"+page+"/i18n.json", &translationsFile)
		if err != nil {
			fmt.Println("Error reading i18n.json.  " + err.Error())
			return
		}
	} else {
		fmt.Println("No existing i18n.json file exists at " + settings.Path + "/" + page + " to update new translations.")
		return
	}

	for i := range settings.Languages {
		lang := settings.Languages[i]

		for key, value := range translationsFile.EN {

			translation, err := tr.Translate(lang, value)
			if err != nil {
				fmt.Println("Error failed to translate " + lang + "value.  " + err.Error())
				continue
			}
			switch lang {
			case "az":
				if translationsFile.AZ == nil {
					translationsFile.AZ = make(map[string]string)
				}
				translationsFile.AZ[key] = translation.Result()
			case "sq":
				if translationsFile.SQ == nil {
					translationsFile.SQ = make(map[string]string)
				}
				translationsFile.SQ[key] = translation.Result()
			case "am":
				if translationsFile.AM == nil {
					translationsFile.AM = make(map[string]string)
				}
				translationsFile.AM[key] = translation.Result()
			case "ar":
				if translationsFile.AR == nil {
					translationsFile.AR = make(map[string]string)
				}
				translationsFile.AR[key] = translation.Result()
			case "hy":
				if translationsFile.HY == nil {
					translationsFile.HY = make(map[string]string)
				}
				translationsFile.HY[key] = translation.Result()
			case "af":
				if translationsFile.AF == nil {
					translationsFile.AF = make(map[string]string)
				}
				translationsFile.AF[key] = translation.Result()
			case "be":
				if translationsFile.BE == nil {
					translationsFile.BE = make(map[string]string)
				}
				translationsFile.BE[key] = translation.Result()
			case "bg":
				if translationsFile.BG == nil {
					translationsFile.BG = make(map[string]string)
				}
				translationsFile.BG[key] = translation.Result()
			case "bs":
				if translationsFile.BS == nil {
					translationsFile.BS = make(map[string]string)
				}
				translationsFile.BG[key] = translation.Result()
			case "cy":
				if translationsFile.CY == nil {
					translationsFile.CY = make(map[string]string)
				}
				translationsFile.CY[key] = translation.Result()
			case "hu":
				if translationsFile.HU == nil {
					translationsFile.HU = make(map[string]string)
				}
				translationsFile.HU[key] = translation.Result()
			case "vi":
				if translationsFile.VI == nil {
					translationsFile.VI = make(map[string]string)
				}
				translationsFile.VI[key] = translation.Result()
			case "ht":
				if translationsFile.HT == nil {
					translationsFile.HT = make(map[string]string)
				}
				translationsFile.HT[key] = translation.Result()
			case "nf":
				if translationsFile.NF == nil {
					translationsFile.NF = make(map[string]string)
				}
				translationsFile.NF[key] = translation.Result()
			case "he":
				if translationsFile.HE == nil {
					translationsFile.HE = make(map[string]string)
				}
				translationsFile.HE[key] = translation.Result()
			case "id":
				if translationsFile.ID == nil {
					translationsFile.ID = make(map[string]string)
				}
				translationsFile.ID[key] = translation.Result()
			case "ga":
				if translationsFile.GA == nil {
					translationsFile.GA = make(map[string]string)
				}
				translationsFile.GA[key] = translation.Result()
			case "it":
				if translationsFile.IT == nil {
					translationsFile.IT = make(map[string]string)
				}
				translationsFile.IT[key] = translation.Result()
			case "is":
				if translationsFile.IS == nil {
					translationsFile.IS = make(map[string]string)
				}
				translationsFile.IS[key] = translation.Result()
			case "ko":
				if translationsFile.KO == nil {
					translationsFile.KO = make(map[string]string)
				}
				translationsFile.KO[key] = translation.Result()
			case "la":
				if translationsFile.LA == nil {
					translationsFile.LA = make(map[string]string)
				}
				translationsFile.LA[key] = translation.Result()
			case "de":
				if translationsFile.DE == nil {
					translationsFile.DE = make(map[string]string)
				}
				translationsFile.DE[key] = translation.Result()
			case "no":
				if translationsFile.NO == nil {
					translationsFile.NO = make(map[string]string)
				}
				translationsFile.NO[key] = translation.Result()
			case "fa":
				if translationsFile.FA == nil {
					translationsFile.FA = make(map[string]string)
				}
				translationsFile.FA[key] = translation.Result()
			case "pl":
				if translationsFile.PL == nil {
					translationsFile.PL = make(map[string]string)
				}
				translationsFile.PL[key] = translation.Result()
			case "pt":
				if translationsFile.PT == nil {
					translationsFile.PT = make(map[string]string)
				}
				translationsFile.PT[key] = translation.Result()
			case "ro":
				if translationsFile.RO == nil {
					translationsFile.RO = make(map[string]string)
				}
				translationsFile.RO[key] = translation.Result()
			case "ru":
				if translationsFile.RU == nil {
					translationsFile.RU = make(map[string]string)
				}
				translationsFile.RU[key] = translation.Result()
			case "sw":
				if translationsFile.SW == nil {
					translationsFile.SW = make(map[string]string)
				}
				translationsFile.SW[key] = translation.Result()
			case "th":
				if translationsFile.TH == nil {
					translationsFile.TH = make(map[string]string)
				}
				translationsFile.TH[key] = translation.Result()
			case "tr":
				if translationsFile.TR == nil {
					translationsFile.TR = make(map[string]string)
				}
				translationsFile.TR[key] = translation.Result()
			case "uk":
				if translationsFile.UK == nil {
					translationsFile.UK = make(map[string]string)
				}
				translationsFile.UK[key] = translation.Result()
			case "fi":
				if translationsFile.FI == nil {
					translationsFile.FI = make(map[string]string)
				}
				translationsFile.FI[key] = translation.Result()
			case "fr":
				if translationsFile.FR == nil {
					translationsFile.FR = make(map[string]string)
				}
				translationsFile.FR[key] = translation.Result()
			case "cs":
				if translationsFile.CS == nil {
					translationsFile.CS = make(map[string]string)
				}
				translationsFile.CS[key] = translation.Result()
			case "sv":
				if translationsFile.SV == nil {
					translationsFile.SV = make(map[string]string)
				}
				translationsFile.SV[key] = translation.Result()
			case "ja":
				if translationsFile.JA == nil {
					translationsFile.JA = make(map[string]string)
				}
				translationsFile.JA[key] = translation.Result()
			case "es":
				if translationsFile.ES == nil {
					translationsFile.ES = make(map[string]string)
				}
				translationsFile.ES[key] = translation.Result()
			}

		}
	}

	parseAndWriteFile(settings.Path+"/"+page+"/i18n.json", translationsFile, 0755)
}

func readFileAndParse(path string, v interface{}) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &v)
	return
}

func parseAndWriteFile(path string, v interface{}, perm os.FileMode) (err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return
	}

	err = writeToFile(string(data), path, perm)
	return
}

func writeToFile(value string, path string, perm os.FileMode) error {
	err := ioutil.WriteFile(path, []byte(value), perm)
	if err != nil {
		log.Println("Error writing file " + path + ":  " + err.Error())
		return err
	}
	return nil
}

func doesFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

type i18nFile struct {
	AZ map[string]string `json:"az,omitempty"`
	SQ map[string]string `json:"sq,omitempty"`
	AM map[string]string `json:"am,omitempty"`
	AR map[string]string `json:"ar,omitempty"`
	HY map[string]string `json:"hy,omitempty"`

	AF map[string]string `json:"af,omitempty"`
	BE map[string]string `json:"be,omitempty"`
	BG map[string]string `json:"bg,omitempty"`
	BS map[string]string `json:"bs,omitempty"`

	EN map[string]string `json:"en,omitempty"`
	ES map[string]string `json:"es,omitempty"`
	ZH map[string]string `json:"zh,omitempty"`

	CY map[string]string `json:"cy,omitempty"`
	HU map[string]string `json:"hu,omitempty"`
	VI map[string]string `json:"vi,omitempty"`
	HT map[string]string `json:"ht,omitempty"`
	NF map[string]string `json:"nf,omitempty"`
	HE map[string]string `json:"he,omitempty"`
	ID map[string]string `json:"id,omitempty"`
	GA map[string]string `json:"ga,omitempty"`
	IT map[string]string `json:"it,omitempty"`
	IS map[string]string `json:"is,omitempty"`
	KO map[string]string `json:"ko,omitempty"`
	LA map[string]string `json:"la,omitempty"`
	DE map[string]string `json:"de,omitempty"`
	NO map[string]string `json:"no,omitempty"`
	FA map[string]string `json:"fa,omitempty"`
	PL map[string]string `json:"pl,omitempty"`
	PT map[string]string `json:"pt,omitempty"`
	RO map[string]string `json:"ro,omitempty"`
	RU map[string]string `json:"ru,omitempty"`
	SW map[string]string `json:"sw,omitempty"`
	TH map[string]string `json:"th,omitempty"`
	TR map[string]string `json:"tr,omitempty"`
	UK map[string]string `json:"uk,omitempty"`
	FI map[string]string `json:"fi,omitempty"`
	FR map[string]string `json:"fr,omitempty"`
	CS map[string]string `json:"cs,omitempty"`
	SV map[string]string `json:"sv,omitempty"`
	JA map[string]string `json:"js,omitempty"`
}
