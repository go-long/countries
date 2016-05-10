package geoutils

//7大洲
func ListContinents() map[string]continent {
	return map[string]continent{
		"AF": continent{
			Name:       "Africa",
			CnName: "非洲",
		},
		"AN": continent{
			Name:  "Antarctica",
			CnName: "南极洲",
		},
		"AS": continent{
			Name:  "Asia",
			CnName: "亚洲",
		},
		"EU": continent{
			Name:  "Europe",
			CnName: "欧洲",
		},
		"NA": continent{
			Name:  "North America",
			CnName: "北美洲",
		},
		"OC": continent{
			Name:  "Oceania",
			CnName: "大洋洲",
		},
		"SA": continent{
			Name:  "South America",
			CnName: "南美洲",
		},
	}
}