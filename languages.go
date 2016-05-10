package geoutils




func ListLanguages() map[string]language {
	list := map[string]language{
		"ab": language{
			Name:       "Abkhaz",
			NativeName: "аҧсуа",
		},
		"aa": language{
			Name:       "Afar",
			NativeName: "Afaraf",
		},
		"af": language{
			Name:       "Afrikaans",
			NativeName: "Afrikaans",
		},
		"ak": language{
			Name:       "Akan",
			NativeName: "Akan",
		},
		"sq": language{
			Name:       "Albanian",
			NativeName: "Shqip",
		},
		"am": language{
			Name:       "Amharic",
			NativeName: "አማርኛ",
		},
		"ar": language{
			Name:       "Arabic",
			NativeName: "العربية",
		},
		"an": language{
			Name:       "Aragonese",
			NativeName: "Aragonés",
		},
		"hy": language{
			Name:       "Armenian",
			NativeName: "Հայերեն",
		},
		"as": language{
			Name:       "Assamese",
			NativeName: "অসমীয়া",
		},
		"av": language{
			Name:       "Avaric",
			NativeName: "авар мацӀ, магӀарул мацӀ",
		},
		"ae": language{
			Name:       "Avestan",
			NativeName: "avesta",
		},
		"ay": language{
			Name:       "Aymara",
			NativeName: "aymar aru",
		},
		"az": language{
			Name:       "Azerbaijani",
			NativeName: "azərbaycan dili",
		},
		"bm": language{
			Name:       "Bambara",
			NativeName: "bamanankan",
		},
		"ba": language{
			Name:       "Bashkir",
			NativeName: "башҡорт теле",
		},
		"eu": language{
			Name:       "Basque",
			NativeName: "euskara, euskera",
		},
		"be": language{
			Name:       "Belarusian",
			NativeName: "Беларуская",
		},
		"bn": language{
			Name:       "Bengali",
			NativeName: "বাংলা",
		},
		"bh": language{
			Name:       "Bihari",
			NativeName: "भोजपुरी",
		},
		"bi": language{
			Name:       "Bislama",
			NativeName: "Bislama",
		},
		"bs": language{
			Name:       "Bosnian",
			NativeName: "bosanski jezik",
		},
		"br": language{
			Name:       "Breton",
			NativeName: "brezhoneg",
		},
		"bg": language{
			Name:       "Bulgarian",
			NativeName: "български език",
		},
		"my": language{
			Name:       "Burmese",
			NativeName: "ဗမာစာ",
		},
		"ca": language{
			Name:       "Catalan; Valencian",
			NativeName: "Català",
		},
		"ch": language{
			Name:       "Chamorro",
			NativeName: "Chamoru",
		},
		"ce": language{
			Name:       "Chechen",
			NativeName: "нохчийн мотт",
		},
		"ny": language{
			Name:       "Chichewa; Chewa; Nyanja",
			NativeName: "chiCheŵa, chinyanja",
		},
		"zh": language{
			Name:       "Chinese",
			NativeName: "中文 (Zhōngwén), 汉语, 漢語",
		},
		"cv": language{
			Name:       "Chuvash",
			NativeName: "чӑваш чӗлхи",
		},
		"kw": language{
			Name:       "Cornish",
			NativeName: "Kernewek",
		},
		"co": language{
			Name:       "Corsican",
			NativeName: "corsu, lingua corsa",
		},
		"cr": language{
			Name:       "Cree",
			NativeName: "ᓀᐦᐃᔭᐍᐏᐣ",
		},
		"hr": language{
			Name:       "Croatian",
			NativeName: "hrvatski",
		},
		"cs": language{
			Name:       "Czech",
			NativeName: "česky, čeština",
		},
		"da": language{
			Name:       "Danish",
			NativeName: "dansk",
		},
		"dv": language{
			Name:       "Divehi; Dhivehi; Maldivian;",
			NativeName: "ދިވެހި",
		},
		"nl": language{
			Name:       "Dutch",
			NativeName: "Nederlands, Vlaams",
		},
		"en": language{
			Name:       "English",
			NativeName: "English",
		},
		"eo": language{
			Name:       "Esperanto",
			NativeName: "Esperanto",
		},
		"et": language{
			Name:       "Estonian",
			NativeName: "eesti, eesti keel",
		},
		"ee": language{
			Name:       "Ewe",
			NativeName: "Eʋegbe",
		},
		"fo": language{
			Name:       "Faroese",
			NativeName: "føroyskt",
		},
		"fj": language{
			Name:       "Fijian",
			NativeName: "vosa Vakaviti",
		},
		"fi": language{
			Name:       "Finnish",
			NativeName: "suomi, suomen kieli",
		},
		"fr": language{
			Name:       "French",
			NativeName: "français, langue française",
		},
		"ff": language{
			Name:       "Fula; Fulah; Pulaar; Pular",
			NativeName: "Fulfulde, Pulaar, Pular",
		},
		"gl": language{
			Name:       "Galician",
			NativeName: "Galego",
		},
		"ka": language{
			Name:       "Georgian",
			NativeName: "ქართული",
		},
		"de": language{
			Name:       "German",
			NativeName: "Deutsch",
		},
		"el": language{
			Name:       "Greek, Modern",
			NativeName: "Ελληνικά",
		},
		"gn": language{
			Name:       "Guaraní",
			NativeName: "Avañeẽ",
		},
		"gu": language{
			Name:       "Gujarati",
			NativeName: "ગુજરાતી",
		},
		"ht": language{
			Name:       "Haitian; Haitian Creole",
			NativeName: "Kreyòl ayisyen",
		},
		"ha": language{
			Name:       "Hausa",
			NativeName: "Hausa, هَوُسَ",
		},
		"he": language{
			Name:       "Hebrew (modern)",
			NativeName: "עברית",
		},
		"hz": language{
			Name:       "Herero",
			NativeName: "Otjiherero",
		},
		"hi": language{
			Name:       "Hindi",
			NativeName: "हिन्दी, हिंदी",
		},
		"ho": language{
			Name:       "Hiri Motu",
			NativeName: "Hiri Motu",
		},
		"hu": language{
			Name:       "Hungarian",
			NativeName: "Magyar",
		},
		"ia": language{
			Name:       "Interlingua",
			NativeName: "Interlingua",
		},
		"id": language{
			Name:       "Indonesian",
			NativeName: "Bahasa Indonesia",
		},
		"ie": language{
			Name:       "Interlingue",
			NativeName: "Originally called Occidental; then Interlingue after WWII",
		},
		"ga": language{
			Name:       "Irish",
			NativeName: "Gaeilge",
		},
		"ig": language{
			Name:       "Igbo",
			NativeName: "Asụsụ Igbo",
		},
		"ik": language{
			Name:       "Inupiaq",
			NativeName: "Iñupiaq, Iñupiatun",
		},
		"io": language{
			Name:       "Ido",
			NativeName: "Ido",
		},
		"is": language{
			Name:       "Icelandic",
			NativeName: "Íslenska",
		},
		"it": language{
			Name:       "Italian",
			NativeName: "Italiano",
		},
		"iu": language{
			Name:       "Inuktitut",
			NativeName: "ᐃᓄᒃᑎᑐᑦ",
		},
		"ja": language{
			Name:       "Japanese",
			NativeName: "日本語 (にほんご／にっぽんご)",
		},
		"jv": language{
			Name:       "Javanese",
			NativeName: "basa Jawa",
		},
		"kl": language{
			Name:       "Kalaallisut, Greenlandic",
			NativeName: "kalaallisut, kalaallit oqaasii",
		},
		"kn": language{
			Name:       "Kannada",
			NativeName: "ಕನ್ನಡ",
		},
		"kr": language{
			Name:       "Kanuri",
			NativeName: "Kanuri",
		},
		"ks": language{
			Name:       "Kashmiri",
			NativeName: "कश्मीरी, كشميري‎",
		},
		"kk": language{
			Name:       "Kazakh",
			NativeName: "Қазақ тілі",
		},
		"km": language{
			Name:       "Khmer",
			NativeName: "ភាសាខ្មែរ",
		},
		"ki": language{
			Name:       "Kikuyu, Gikuyu",
			NativeName: "Gĩkũyũ",
		},
		"rw": language{
			Name:       "Kinyarwanda",
			NativeName: "Ikinyarwanda",
		},
		"ky": language{
			Name:       "Kirghiz, Kyrgyz",
			NativeName: "кыргыз тили",
		},
		"kv": language{
			Name:       "Komi",
			NativeName: "коми кыв",
		},
		"kg": language{
			Name:       "Kongo",
			NativeName: "KiKongo",
		},
		"ko": language{
			Name:       "Korean",
			NativeName: "한국어 (韓國語), 조선말 (朝鮮語)",
		},
		"ku": language{
			Name:       "Kurdish",
			NativeName: "Kurdî, كوردی‎",
		},
		"kj": language{
			Name:       "Kwanyama, Kuanyama",
			NativeName: "Kuanyama",
		},
		"la": language{
			Name:       "Latin",
			NativeName: "latine, lingua latina",
		},
		"lb": language{
			Name:       "Luxembourgish, Letzeburgesch",
			NativeName: "Lëtzebuergesch",
		},
		"lg": language{
			Name:       "Luganda",
			NativeName: "Luganda",
		},
		"li": language{
			Name:       "Limburgish, Limburgan, Limburger",
			NativeName: "Limburgs",
		},
		"ln": language{
			Name:       "Lingala",
			NativeName: "Lingála",
		},
		"lo": language{
			Name:       "Lao",
			NativeName: "ພາສາລາວ",
		},
		"lt": language{
			Name:       "Lithuanian",
			NativeName: "lietuvių kalba",
		},
		"lu": language{
			Name:       "Luba-Katanga",
			NativeName: "",
		},
		"lv": language{
			Name:       "Latvian",
			NativeName: "latviešu valoda",
		},
		"gv": language{
			Name:       "Manx",
			NativeName: "Gaelg, Gailck",
		},
		"mk": language{
			Name:       "Macedonian",
			NativeName: "македонски јазик",
		},
		"mg": language{
			Name:       "Malagasy",
			NativeName: "Malagasy fiteny",
		},
		"ms": language{
			Name:       "Malay",
			NativeName: "bahasa Melayu, بهاس ملايو‎",
		},
		"ml": language{
			Name:       "Malayalam",
			NativeName: "മലയാളം",
		},
		"mt": language{
			Name:       "Maltese",
			NativeName: "Malti",
		},
		"mi": language{
			Name:       "Māori",
			NativeName: "te reo Māori",
		},
		"mr": language{
			Name:       "Marathi (Marāṭhī)",
			NativeName: "मराठी",
		},
		"mh": language{
			Name:       "Marshallese",
			NativeName: "Kajin M̧ajeļ",
		},
		"mn": language{
			Name:       "Mongolian",
			NativeName: "монгол",
		},
		"na": language{
			Name:       "Nauru",
			NativeName: "Ekakairũ Naoero",
		},
		"nv": language{
			Name:       "Navajo, Navaho",
			NativeName: "Diné bizaad, Dinékʼehǰí",
		},
		"nb": language{
			Name:       "Norwegian Bokmål",
			NativeName: "Norsk bokmål",
		},
		"nd": language{
			Name:       "North Ndebele",
			NativeName: "isiNdebele",
		},
		"ne": language{
			Name:       "Nepali",
			NativeName: "नेपाली",
		},
		"ng": language{
			Name:       "Ndonga",
			NativeName: "Owambo",
		},
		"nn": language{
			Name:       "Norwegian Nynorsk",
			NativeName: "Norsk nynorsk",
		},
		"no": language{
			Name:       "Norwegian",
			NativeName: "Norsk",
		},
		"ii": language{
			Name:       "Nuosu",
			NativeName: "ꆈꌠ꒿ Nuosuhxop",
		},
		"nr": language{
			Name:       "South Ndebele",
			NativeName: "isiNdebele",
		},
		"oc": language{
			Name:       "Occitan",
			NativeName: "Occitan",
		},
		"oj": language{
			Name:       "Ojibwe, Ojibwa",
			NativeName: "ᐊᓂᔑᓈᐯᒧᐎᓐ",
		},
		"cu": language{
			Name:       "Old Church Slavonic, Church Slavic, Church Slavonic, Old Bulgarian, Old Slavonic",
			NativeName: "ѩзыкъ словѣньскъ",
		},
		"om": language{
			Name:       "Oromo",
			NativeName: "Afaan Oromoo",
		},
		"or": language{
			Name:       "Oriya",
			NativeName: "ଓଡ଼ିଆ",
		},
		"os": language{
			Name:       "Ossetian, Ossetic",
			NativeName: "ирон æвзаг",
		},
		"pa": language{
			Name:       "Panjabi, Punjabi",
			NativeName: "ਪੰਜਾਬੀ, پنجابی‎",
		},
		"pi": language{
			Name:       "Pāli",
			NativeName: "पाऴि",
		},
		"fa": language{
			Name:       "Persian",
			NativeName: "فارسی",
		},
		"pl": language{
			Name:       "Polish",
			NativeName: "polski",
		},
		"ps": language{
			Name:       "Pashto, Pushto",
			NativeName: "پښتو",
		},
		"pt": language{
			Name:       "Portuguese",
			NativeName: "Português",
		},
		"qu": language{
			Name:       "Quechua",
			NativeName: "Runa Simi, Kichwa",
		},
		"rm": language{
			Name:       "Romansh",
			NativeName: "rumantsch grischun",
		},
		"rn": language{
			Name:       "Kirundi",
			NativeName: "kiRundi",
		},
		"ro": language{
			Name:       "Romanian, Moldavian, Moldovan",
			NativeName: "română",
		},
		"ru": language{
			Name:       "Russian",
			NativeName: "русский язык",
		},
		"sa": language{
			Name:       "Sanskrit (Saṁskṛta)",
			NativeName: "संस्कृतम्",
		},
		"sc": language{
			Name:       "Sardinian",
			NativeName: "sardu",
		},
		"sd": language{
			Name:       "Sindhi",
			NativeName: "सिन्धी, سنڌي، سندھی‎",
		},
		"se": language{
			Name:       "Northern Sami",
			NativeName: "Davvisámegiella",
		},
		"sm": language{
			Name:       "Samoan",
			NativeName: "gagana faa Samoa",
		},
		"sg": language{
			Name:       "Sango",
			NativeName: "yângâ tî sängö",
		},
		"sr": language{
			Name:       "Serbian",
			NativeName: "српски језик",
		},
		"gd": language{
			Name:       "Scottish Gaelic; Gaelic",
			NativeName: "Gàidhlig",
		},
		"sn": language{
			Name:       "Shona",
			NativeName: "chiShona",
		},
		"si": language{
			Name:       "Sinhala, Sinhalese",
			NativeName: "සිංහල",
		},
		"sk": language{
			Name:       "Slovak",
			NativeName: "slovenčina",
		},
		"sl": language{
			Name:       "Slovene",
			NativeName: "slovenščina",
		},
		"so": language{
			Name:       "Somali",
			NativeName: "Soomaaliga, af Soomaali",
		},
		"st": language{
			Name:       "Southern Sotho",
			NativeName: "Sesotho",
		},
		"es": language{
			Name:       "Spanish",
			NativeName: "español, castellano",
		},
		"su": language{
			Name:       "Sundanese",
			NativeName: "Basa Sunda",
		},
		"sw": language{
			Name:       "Swahili",
			NativeName: "Kiswahili",
		},
		"ss": language{
			Name:       "Swati",
			NativeName: "SiSwati",
		},
		"sv": language{
			Name:       "Swedish",
			NativeName: "svenska",
		},
		"ta": language{
			Name:       "Tamil",
			NativeName: "தமிழ்",
		},
		"te": language{
			Name:       "Telugu",
			NativeName: "తెలుగు",
		},
		"tg": language{
			Name:       "Tajik",
			NativeName: "тоҷикӣ, toğikī, تاجیکی‎",
		},
		"th": language{
			Name:       "Thai",
			NativeName: "ไทย",
		},
		"ti": language{
			Name:       "Tigrinya",
			NativeName: "ትግርኛ",
		},
		"bo": language{
			Name:       "Tibetan Standard, Tibetan, Central",
			NativeName: "བོད་ཡིག",
		},
		"tk": language{
			Name:       "Turkmen",
			NativeName: "Türkmen, Түркмен",
		},
		"tl": language{
			Name:       "Tagalog",
			NativeName: "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔",
		},
		"tn": language{
			Name:       "Tswana",
			NativeName: "Setswana",
		},
		"to": language{
			Name:       "Tonga (Tonga Islands)",
			NativeName: "faka Tonga",
		},
		"tr": language{
			Name:       "Turkish",
			NativeName: "Türkçe",
		},
		"ts": language{
			Name:       "Tsonga",
			NativeName: "Xitsonga",
		},
		"tt": language{
			Name:       "Tatar",
			NativeName: "татарча, tatarça, تاتارچا‎",
		},
		"tw": language{
			Name:       "Twi",
			NativeName: "Twi",
		},
		"ty": language{
			Name:       "Tahitian",
			NativeName: "Reo Tahiti",
		},
		"ug": language{
			Name:       "Uighur, Uyghur",
			NativeName: "Uyƣurqə, ئۇيغۇرچە‎",
		},
		"uk": language{
			Name:       "Ukrainian",
			NativeName: "українська",
		},
		"ur": language{
			Name:       "Urdu",
			NativeName: "اردو",
		},
		"uz": language{
			Name:       "Uzbek",
			NativeName: "zbek, Ўзбек, أۇزبېك‎",
		},
		"ve": language{
			Name:       "Venda",
			NativeName: "Tshivenḓa",
		},
		"vi": language{
			Name:       "Vietnamese",
			NativeName: "Tiếng Việt",
		},
		"vo": language{
			Name:       "Volapük",
			NativeName: "Volapük",
		},
		"wa": language{
			Name:       "Walloon",
			NativeName: "Walon",
		},
		"cy": language{
			Name:       "Welsh",
			NativeName: "Cymraeg",
		},
		"wo": language{
			Name:       "Wolof",
			NativeName: "Wollof",
		},
		"fy": language{
			Name:       "Western Frisian",
			NativeName: "Frysk",
		},
		"xh": language{
			Name:       "Xhosa",
			NativeName: "isiXhosa",
		},
		"yi": language{
			Name:       "Yiddish",
			NativeName: "ייִדיש",
		},
		"yo": language{
			Name:       "Yoruba",
			NativeName: "Yorùbá",
		},
		"za": language{
			Name:       "Zhuang, Chuang",
			NativeName: "Saɯ cueŋƅ, Saw cuengh",
		},
	}

	return list
}

