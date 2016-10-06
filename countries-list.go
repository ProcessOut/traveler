package countries

// CountryList returns the country list map which was built from
// https://en.wikipedia.org/wiki/ISO_3166-1 and is based on the Alpha2 for indexing.
// It is recommended you go through the Verify() or Information() if possible.
func CountryList() map[Alpha2]CountryInformation {
	return countryList
}

var countryList = map[Alpha2]CountryInformation{
	"AF": CountryInformation{"Afghanistan", "AFG", 4},
	"AX": CountryInformation{"Aland Islands !Åland Islands", "ALA", 248},
	"AL": CountryInformation{"Albania", "ALB", 8},
	"DZ": CountryInformation{"Algeria", "DZA", 12},
	"AS": CountryInformation{"American Samoa", "ASM", 16},
	"AD": CountryInformation{"Andorra", "AND", 20},
	"AO": CountryInformation{"Angola", "AGO", 24},
	"AI": CountryInformation{"Anguilla", "AIA", 660},
	"AQ": CountryInformation{"Antarctica", "ATA", 10},
	"AG": CountryInformation{"Antigua and Barbuda", "ATG", 28},
	"AR": CountryInformation{"Argentina", "ARG", 32},
	"AM": CountryInformation{"Armenia", "ARM", 51},
	"AW": CountryInformation{"Aruba", "ABW", 533},
	"AU": CountryInformation{"Australia", "AUS", 36},
	"AT": CountryInformation{"Austria", "AUT", 40},
	"AZ": CountryInformation{"Azerbaijan", "AZE", 31},
	"BS": CountryInformation{"Bahamas", "BHS", 44},
	"BH": CountryInformation{"Bahrain", "BHR", 48},
	"BD": CountryInformation{"Bangladesh", "BGD", 50},
	"BB": CountryInformation{"Barbados", "BRB", 52},
	"BY": CountryInformation{"Belarus", "BLR", 112},
	"BE": CountryInformation{"Belgium", "BEL", 56},
	"BZ": CountryInformation{"Belize", "BLZ", 84},
	"BJ": CountryInformation{"Benin", "BEN", 204},
	"BM": CountryInformation{"Bermuda", "BMU", 60},
	"BT": CountryInformation{"Bhutan", "BTN", 64},
	"BO": CountryInformation{"Bolivia (Plurinational State of)", "BOL", 68},
	"BQ": CountryInformation{"Bonaire, Sint Eustatius and Saba", "BES", 535},
	"BA": CountryInformation{"Bosnia and Herzegovina", "BIH", 70},
	"BW": CountryInformation{"Botswana", "BWA", 72},
	"BV": CountryInformation{"Bouvet Island", "BVT", 74},
	"BR": CountryInformation{"Brazil", "BRA", 76},
	"IO": CountryInformation{"British Indian Ocean Territory", "IOT", 86},
	"BN": CountryInformation{"Brunei Darussalam", "BRN", 96},
	"BG": CountryInformation{"Bulgaria", "BGR", 100},
	"BF": CountryInformation{"Burkina Faso", "BFA", 854},
	"BI": CountryInformation{"Burundi", "BDI", 108},
	"CV": CountryInformation{"Cabo Verde", "CPV", 132},
	"KH": CountryInformation{"Cambodia", "KHM", 116},
	"CM": CountryInformation{"Cameroon", "CMR", 120},
	"CA": CountryInformation{"Canada", "CAN", 124},
	"KY": CountryInformation{"Cayman Islands", "CYM", 136},
	"CF": CountryInformation{"Central African Republic", "CAF", 140},
	"TD": CountryInformation{"Chad", "TCD", 148},
	"CL": CountryInformation{"Chile", "CHL", 152},
	"CN": CountryInformation{"China", "CHN", 156},
	"CX": CountryInformation{"Christmas Island", "CXR", 162},
	"CC": CountryInformation{"Cocos (Keeling) Islands", "CCK", 166},
	"CO": CountryInformation{"Colombia", "COL", 170},
	"KM": CountryInformation{"Comoros", "COM", 174},
	"CG": CountryInformation{"Congo", "COG", 178},
	"CD": CountryInformation{"Congo (Democratic Republic of the)", "COD", 180},
	"CK": CountryInformation{"Cook Islands", "COK", 184},
	"CR": CountryInformation{"Costa Rica", "CRI", 188},
	"CI": CountryInformation{"Cote d'Ivoire !Côte d'Ivoire", "CIV", 384},
	"HR": CountryInformation{"Croatia", "HRV", 191},
	"CU": CountryInformation{"Cuba", "CUB", 192},
	"CW": CountryInformation{"Curacao !Curaçao", "CUW", 531},
	"CY": CountryInformation{"Cyprus", "CYP", 196},
	"CZ": CountryInformation{"Czechia", "CZE", 203},
	"DK": CountryInformation{"Denmark", "DNK", 208},
	"DJ": CountryInformation{"Djibouti", "DJI", 262},
	"DM": CountryInformation{"Dominica", "DMA", 212},
	"DO": CountryInformation{"Dominican Republic", "DOM", 214},
	"EC": CountryInformation{"Ecuador", "ECU", 218},
	"EG": CountryInformation{"Egypt", "EGY", 818},
	"SV": CountryInformation{"El Salvador", "SLV", 222},
	"GQ": CountryInformation{"Equatorial Guinea", "GNQ", 226},
	"ER": CountryInformation{"Eritrea", "ERI", 232},
	"EE": CountryInformation{"Estonia", "EST", 233},
	"ET": CountryInformation{"Ethiopia", "ETH", 231},
	"FK": CountryInformation{"Falkland Islands (Malvinas)", "FLK", 238},
	"FO": CountryInformation{"Faroe Islands", "FRO", 234},
	"FJ": CountryInformation{"Fiji", "FJI", 242},
	"FI": CountryInformation{"Finland", "FIN", 246},
	"FR": CountryInformation{"France", "FRA", 250},
	"GF": CountryInformation{"French Guiana", "GUF", 254},
	"PF": CountryInformation{"French Polynesia", "PYF", 258},
	"TF": CountryInformation{"French Southern Territories", "ATF", 260},
	"GA": CountryInformation{"Gabon", "GAB", 266},
	"GM": CountryInformation{"Gambia", "GMB", 270},
	"GE": CountryInformation{"Georgia", "GEO", 268},
	"DE": CountryInformation{"Germany", "DEU", 276},
	"GH": CountryInformation{"Ghana", "GHA", 288},
	"GI": CountryInformation{"Gibraltar", "GIB", 292},
	"GR": CountryInformation{"Greece", "GRC", 300},
	"GL": CountryInformation{"Greenland", "GRL", 304},
	"GD": CountryInformation{"Grenada", "GRD", 308},
	"GP": CountryInformation{"Guadeloupe", "GLP", 312},
	"GU": CountryInformation{"Guam", "GUM", 316},
	"GT": CountryInformation{"Guatemala", "GTM", 320},
	"GG": CountryInformation{"Guernsey", "GGY", 831},
	"GN": CountryInformation{"Guinea", "GIN", 324},
	"GW": CountryInformation{"Guinea-Bissau", "GNB", 624},
	"GY": CountryInformation{"Guyana", "GUY", 328},
	"HT": CountryInformation{"Haiti", "HTI", 332},
	"HM": CountryInformation{"Heard Island and McDonald Islands", "HMD", 334},
	"VA": CountryInformation{"Holy See", "VAT", 336},
	"HN": CountryInformation{"Honduras", "HND", 340},
	"HK": CountryInformation{"Hong Kong", "HKG", 344},
	"HU": CountryInformation{"Hungary", "HUN", 348},
	"IS": CountryInformation{"Iceland", "ISL", 352},
	"IN": CountryInformation{"India", "IND", 356},
	"ID": CountryInformation{"Indonesia", "IDN", 360},
	"IR": CountryInformation{"Iran (Islamic Republic of)", "IRN", 364},
	"IQ": CountryInformation{"Iraq", "IRQ", 368},
	"IE": CountryInformation{"Ireland", "IRL", 372},
	"IM": CountryInformation{"Isle of Man", "IMN", 833},
	"IL": CountryInformation{"Israel", "ISR", 376},
	"IT": CountryInformation{"Italy", "ITA", 380},
	"JM": CountryInformation{"Jamaica", "JAM", 388},
	"JP": CountryInformation{"Japan", "JPN", 392},
	"JE": CountryInformation{"Jersey", "JEY", 832},
	"JO": CountryInformation{"Jordan", "JOR", 400},
	"KZ": CountryInformation{"Kazakhstan", "KAZ", 398},
	"KE": CountryInformation{"Kenya", "KEN", 404},
	"KI": CountryInformation{"Kiribati", "KIR", 296},
	"KP": CountryInformation{"Korea (Democratic People's Republic of)", "PRK", 408},
	"KR": CountryInformation{"Korea (Republic of)", "KOR", 410},
	"KW": CountryInformation{"Kuwait", "KWT", 414},
	"KG": CountryInformation{"Kyrgyzstan", "KGZ", 417},
	"LA": CountryInformation{"Lao People's Democratic Republic", "LAO", 418},
	"LV": CountryInformation{"Latvia", "LVA", 428},
	"LB": CountryInformation{"Lebanon", "LBN", 422},
	"LS": CountryInformation{"Lesotho", "LSO", 426},
	"LR": CountryInformation{"Liberia", "LBR", 430},
	"LY": CountryInformation{"Libya", "LBY", 434},
	"LI": CountryInformation{"Liechtenstein", "LIE", 438},
	"LT": CountryInformation{"Lithuania", "LTU", 440},
	"LU": CountryInformation{"Luxembourg", "LUX", 442},
	"MO": CountryInformation{"Macao", "MAC", 446},
	"MK": CountryInformation{"Macedonia (the former Yugoslav Republic of)", "MKD", 807},
	"MG": CountryInformation{"Madagascar", "MDG", 450},
	"MW": CountryInformation{"Malawi", "MWI", 454},
	"MY": CountryInformation{"Malaysia", "MYS", 458},
	"MV": CountryInformation{"Maldives", "MDV", 462},
	"ML": CountryInformation{"Mali", "MLI", 466},
	"MT": CountryInformation{"Malta", "MLT", 470},
	"MH": CountryInformation{"Marshall Islands", "MHL", 584},
	"MQ": CountryInformation{"Martinique", "MTQ", 474},
	"MR": CountryInformation{"Mauritania", "MRT", 478},
	"MU": CountryInformation{"Mauritius", "MUS", 480},
	"YT": CountryInformation{"Mayotte", "MYT", 175},
	"MX": CountryInformation{"Mexico", "MEX", 484},
	"FM": CountryInformation{"Micronesia (Federated States of)", "FSM", 583},
	"MD": CountryInformation{"Moldova (Republic of)", "MDA", 498},
	"MC": CountryInformation{"Monaco", "MCO", 492},
	"MN": CountryInformation{"Mongolia", "MNG", 496},
	"ME": CountryInformation{"Montenegro", "MNE", 499},
	"MS": CountryInformation{"Montserrat", "MSR", 500},
	"MA": CountryInformation{"Morocco", "MAR", 504},
	"MZ": CountryInformation{"Mozambique", "MOZ", 508},
	"MM": CountryInformation{"Myanmar", "MMR", 104},
	"NA": CountryInformation{"Namibia", "NAM", 516},
	"NR": CountryInformation{"Nauru", "NRU", 520},
	"NP": CountryInformation{"Nepal", "NPL", 524},
	"NL": CountryInformation{"Netherlands", "NLD", 528},
	"NC": CountryInformation{"New Caledonia", "NCL", 540},
	"NZ": CountryInformation{"New Zealand", "NZL", 554},
	"NI": CountryInformation{"Nicaragua", "NIC", 558},
	"NE": CountryInformation{"Niger", "NER", 562},
	"NG": CountryInformation{"Nigeria", "NGA", 566},
	"NU": CountryInformation{"Niue", "NIU", 570},
	"NF": CountryInformation{"Norfolk Island", "NFK", 574},
	"MP": CountryInformation{"Northern Mariana Islands", "MNP", 580},
	"NO": CountryInformation{"Norway", "NOR", 578},
	"OM": CountryInformation{"Oman", "OMN", 512},
	"PK": CountryInformation{"Pakistan", "PAK", 586},
	"PW": CountryInformation{"Palau", "PLW", 585},
	"PS": CountryInformation{"Palestine, State of", "PSE", 275},
	"PA": CountryInformation{"Panama", "PAN", 591},
	"PG": CountryInformation{"Papua New Guinea", "PNG", 598},
	"PY": CountryInformation{"Paraguay", "PRY", 600},
	"PE": CountryInformation{"Peru", "PER", 604},
	"PH": CountryInformation{"Philippines", "PHL", 608},
	"PN": CountryInformation{"Pitcairn", "PCN", 612},
	"PL": CountryInformation{"Poland", "POL", 616},
	"PT": CountryInformation{"Portugal", "PRT", 620},
	"PR": CountryInformation{"Puerto Rico", "PRI", 630},
	"QA": CountryInformation{"Qatar", "QAT", 634},
	"RE": CountryInformation{"Reunion !Réunion", "REU", 638},
	"RO": CountryInformation{"Romania", "ROU", 642},
	"RU": CountryInformation{"Russian Federation", "RUS", 643},
	"RW": CountryInformation{"Rwanda", "RWA", 646},
	"BL": CountryInformation{"Saint Barthelemy !Saint Barthélemy", "BLM", 652},
	"SH": CountryInformation{"Saint Helena, Ascension and Tristan da Cunha", "SHN", 654},
	"KN": CountryInformation{"Saint Kitts and Nevis", "KNA", 659},
	"LC": CountryInformation{"Saint Lucia", "LCA", 662},
	"MF": CountryInformation{"Saint Martin (French part)", "MAF", 663},
	"PM": CountryInformation{"Saint Pierre and Miquelon", "SPM", 666},
	"VC": CountryInformation{"Saint Vincent and the Grenadines", "VCT", 670},
	"WS": CountryInformation{"Samoa", "WSM", 882},
	"SM": CountryInformation{"San Marino", "SMR", 674},
	"ST": CountryInformation{"Sao Tome and Principe", "STP", 678},
	"SA": CountryInformation{"Saudi Arabia", "SAU", 682},
	"SN": CountryInformation{"Senegal", "SEN", 686},
	"RS": CountryInformation{"Serbia", "SRB", 688},
	"SC": CountryInformation{"Seychelles", "SYC", 690},
	"SL": CountryInformation{"Sierra Leone", "SLE", 694},
	"SG": CountryInformation{"Singapore", "SGP", 702},
	"SX": CountryInformation{"Sint Maarten (Dutch part)", "SXM", 534},
	"SK": CountryInformation{"Slovakia", "SVK", 703},
	"SI": CountryInformation{"Slovenia", "SVN", 705},
	"SB": CountryInformation{"Solomon Islands", "SLB", 90},
	"SO": CountryInformation{"Somalia", "SOM", 706},
	"ZA": CountryInformation{"South Africa", "ZAF", 710},
	"GS": CountryInformation{"South Georgia and the South Sandwich Islands", "SGS", 239},
	"SS": CountryInformation{"South Sudan", "SSD", 728},
	"ES": CountryInformation{"Spain", "ESP", 724},
	"LK": CountryInformation{"Sri Lanka", "LKA", 144},
	"SD": CountryInformation{"Sudan", "SDN", 729},
	"SR": CountryInformation{"Suriname", "SUR", 740},
	"SJ": CountryInformation{"Svalbard and Jan Mayen", "SJM", 744},
	"SZ": CountryInformation{"Swaziland", "SWZ", 748},
	"SE": CountryInformation{"Sweden", "SWE", 752},
	"CH": CountryInformation{"Switzerland", "CHE", 756},
	"SY": CountryInformation{"Syrian Arab Republic", "SYR", 760},
	"TW": CountryInformation{"Taiwan, Province of China", "TWN", 158},
	"TJ": CountryInformation{"Tajikistan", "TJK", 762},
	"TZ": CountryInformation{"Tanzania, United Republic of", "TZA", 834},
	"TH": CountryInformation{"Thailand", "THA", 764},
	"TL": CountryInformation{"Timor-Leste", "TLS", 626},
	"TG": CountryInformation{"Togo", "TGO", 768},
	"TK": CountryInformation{"Tokelau", "TKL", 772},
	"TO": CountryInformation{"Tonga", "TON", 776},
	"TT": CountryInformation{"Trinidad and Tobago", "TTO", 780},
	"TN": CountryInformation{"Tunisia", "TUN", 788},
	"TR": CountryInformation{"Turkey", "TUR", 792},
	"TM": CountryInformation{"Turkmenistan", "TKM", 795},
	"TC": CountryInformation{"Turks and Caicos Islands", "TCA", 796},
	"TV": CountryInformation{"Tuvalu", "TUV", 798},
	"UG": CountryInformation{"Uganda", "UGA", 800},
	"UA": CountryInformation{"Ukraine", "UKR", 804},
	"AE": CountryInformation{"United Arab Emirates", "ARE", 784},
	"GB": CountryInformation{"United Kingdom of Great Britain and Northern Ireland", "GBR", 826},
	"US": CountryInformation{"United States of America", "USA", 840},
	"UM": CountryInformation{"United States Minor Outlying Islands", "UMI", 581},
	"UY": CountryInformation{"Uruguay", "URY", 858},
	"UZ": CountryInformation{"Uzbekistan", "UZB", 860},
	"VU": CountryInformation{"Vanuatu", "VUT", 548},
	"VE": CountryInformation{"Venezuela (Bolivarian Republic of)", "VEN", 862},
	"VN": CountryInformation{"Viet Nam", "VNM", 704},
	"VG": CountryInformation{"Virgin Islands (British)", "VGB", 92},
	"VI": CountryInformation{"Virgin Islands (U.S.)", "VIR", 850},
	"WF": CountryInformation{"Wallis and Futuna", "WLF", 876},
	"EH": CountryInformation{"Western Sahara", "ESH", 732},
	"YE": CountryInformation{"Yemen", "YEM", 887},
	"ZM": CountryInformation{"Zambia", "ZMB", 894},
	"ZW": CountryInformation{"Zimbabwe", "ZWE", 716},
}
