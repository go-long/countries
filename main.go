package geoutils

import (
	"strings"
)

type(
	continent struct {
		Name   string
		CnName string
	}

	language struct {
		Name       string
		NativeName string
	}

	country struct {
		Name       string
		NativeName string
		CnName     string
		Phone      string  //国际电话区号
		Continent  string  //所在洲
		Capital    string  //首都
		Currency   string  //货币
		Languages  string  //使用语言
	}


)

func GetContinents(code string) continent {
	code = strings.ToUpper(code)
	return ListContinents()[code]
}


func GetLanguage(code string) language {
	code = strings.ToLower(code)
	return ListLanguages()[code]
}

func GetCountry(code string) country {
	code = strings.ToUpper(code)
	return ListCountries()[code]
}
