// Package timezone provides helpers to work with time zone.
package timezone

import (
	"errors"
	"github.com/choonsiong/golib/logger/jsonlog"
)

type Timezone struct {
	Logger    *jsonlog.Logger
	Timezones map[string]string
}

var (
	ErrTimezoneIsEmpty    = errors.New("timezone is empty")
	ErrNoMatchingTimezone = errors.New("no matching timezone found")
)

var timezone = map[string]string{
	"Australia/ACT":                    "ACT",
	"Africa/Abidjan":                   "Abidjan",
	"Africa/Accra":                     "Accra",
	"America/Adak":                     "Adak",
	"Africa/Addis_Ababa":               "Addis Ababa",
	"Australia/Adelaide":               "Adelaide",
	"Asia/Aden":                        "Aden",
	"Africa/Algiers":                   "Algiers",
	"Asia/Almaty":                      "Almaty",
	"Asia/Amman":                       "Amman",
	"Europe/Amsterdam":                 "Amsterdam",
	"Asia/Anadyr":                      "Anadyr",
	"America/Anchorage":                "Anchorage",
	"Europe/Andorra":                   "Andorra",
	"America/Anguilla":                 "Anguilla",
	"Indian/Antananarivo":              "Antananarivo",
	"America/Antigua":                  "Antigua",
	"Pacific/Apia":                     "Apia",
	"Asia/Aqtau":                       "Aqtau",
	"Asia/Aqtobe":                      "Aqtobe",
	"America/Araguaina":                "Araguaina",
	"America/Argentina/Buenos_Aires":   "Argentina - Buenos Aires",
	"America/Argentina/Catamarca":      "Argentina - Catamarca",
	"America/Argentina/ComodRivadavia": "Argentina - ComodRivadavia",
	"America/Argentina/Cordoba":        "Argentina - Cordoba",
	"America/Argentina/Jujuy":          "Argentina - Jujuy",
	"America/Argentina/La_Rioja":       "Argentina - La Rioja",
	"America/Argentina/Mendoza":        "Argentina - Mendoza",
	"America/Argentina/Rio_Gallegos":   "Argentina - Rio Gallegos",
	"America/Argentina/Salta":          "Argentina - Salta",
	"America/Argentina/San_Juan":       "Argentina - San Juan",
	"America/Argentina/San_Luis":       "Argentina - San Luis",
	"America/Argentina/Tucuman":        "Argentina - Tucuman",
	"America/Argentina/Ushuaia":        "Argentina - Ushuaia",
	"America/Aruba":                    "Aruba",
	"Asia/Ashgabat":                    "Ashgabat",
	"Asia/Ashkhabad":                   "Ashkhabad",
	"AST":                              "Alaska Standard Time",
	"Africa/Asmara":                    "Asmara",
	"Africa/Asmera":                    "Asmera",
	"America/Asuncion":                 "Asuncion",
	"Europe/Athens":                    "Athens",
	"America/Atikokan":                 "Atikokan",
	"America/Atka":                     "Atka",
	"Pacific/Auckland":                 "Auckland",
	"Atlantic/Azores":                  "Azores",
	"Asia/Baghdad":                     "Baghdad",
	"America/Bahia":                    "Bahia",
	"Asia/Bahrain":                     "Bahrain",
	"Asia/Baku":                        "Baku",
	"Africa/Bamako":                    "Bamako",
	"Asia/Bangkok":                     "Bangkok",
	"Africa/Bangui":                    "Bangui",
	"Africa/Banjul":                    "Banjul",
	"America/Barbados":                 "Barbados",
	"Asia/Beirut":                      "Beirut",
	"America/Belem":                    "Belem",
	"Europe/Belfast":                   "Belfast",
	"Europe/Belgrade":                  "Belgrade",
	"America/Belize":                   "Belize",
	"Europe/Berlin":                    "Berlin",
	"Atlantic/Bermuda":                 "Bermuda",
	"Asia/Bishkek":                     "Bishkek",
	"Africa/Bissau":                    "Bissau",
	"America/Blanc-Sablon":             "Blanc-Sablon",
	"Africa/Blantyre":                  "Blantyre",
	"America/Boa_Vista":                "Boa Vista",
	"America/Bogota":                   "Bogota",
	"America/Boise":                    "Boise",
	"Europe/Bratislava":                "Bratislava",
	"Africa/Brazzaville":               "Brazzaville",
	"Australia/Brisbane":               "Brisbane",
	"Australia/Broken_Hill":            "Broken Hill",
	"Asia/Brunei":                      "Brunei",
	"Europe/Brussels":                  "Brussels",
	"Europe/Bucharest":                 "Bucharest",
	"Europe/Budapest":                  "Budapest",
	"America/Buenos_Aires":             "Buenos Aires",
	"Africa/Bujumbura":                 "Bujumbura",
	"Africa/Cairo":                     "Cairo",
	"Asia/Calcutta":                    "Calcutta",
	"America/Cambridge_Bay":            "Cambridge Bay",
	"America/Campo_Grande":             "Campo Grande",
	"Atlantic/Canary":                  "Canary",
	"Australia/Canberra":               "Canberra",
	"America/Cancun":                   "Cancun",
	"Atlantic/Cape_Verde":              "Cape Verde",
	"America/Caracas":                  "Caracas",
	"Africa/Casablanca":                "Casablanca",
	"Antarctica/Casey":                 "Casey",
	"America/Catamarca":                "Catamarca",
	"America/Cayenne":                  "Cayenne",
	"America/Cayman":                   "Cayman",
	"CAT":                              "Central African Time",
	"CET":                              "Central European Time",
	"ECT":                              "Central European Time",
	"CST":                              "Central Standard Time",
	"Africa/Ceuta":                     "Ceuta",
	"Indian/Chagos":                    "Chagos",
	"Pacific/Chatham":                  "Chatham",
	"America/Chicago":                  "Chicago",
	"America/Chihuahua":                "Chihuahua",
	"CTT":                              "China Standard Time",
	"PRC":                              "China Standard Time",
	"Europe/Chisinau":                  "Chisinau",
	"Asia/Choibalsan":                  "Choibalsan",
	"Asia/Chongqing":                   "Chongqing",
	"Indian/Christmas":                 "Christmas",
	"Asia/Chungking":                   "Chungking",
	"Indian/Cocos":                     "Cocos",
	"Asia/Colombo":                     "Colombo",
	"Indian/Comoro":                    "Comoro",
	"Africa/Conakry":                   "Conakry",
	"Europe/Copenhagen":                "Copenhagen",
	"America/Coral_Harbour":            "Coral Harbour",
	"America/Cordoba":                  "Cordoba",
	"America/Costa_Rica":               "Costa Rica",
	"America/Cuiaba":                   "Cuiaba",
	"America/Curacao":                  "Curacao",
	"Australia/Currie":                 "Currie",
	"Asia/Dacca":                       "Dacca",
	"Africa/Dakar":                     "Dakar",
	"Asia/Damascus":                    "Damascus",
	"America/Danmarkshavn":             "Danmarkshavn",
	"Africa/Dar_es_Salaam":             "Dar es Salaam",
	"Australia/Darwin":                 "Darwin",
	"Antarctica/Davis":                 "Davis",
	"America/Dawson":                   "Dawson",
	"America/Dawson_Creek":             "Dawson Creek",
	"America/Denver":                   "Denver",
	"America/Detroit":                  "Detroit",
	"Asia/Dhaka":                       "Dhaka",
	"Asia/Dili":                        "Dili",
	"Africa/Djibouti":                  "Djibouti",
	"America/Dominica":                 "Dominica",
	"Africa/Douala":                    "Douala",
	"Asia/Dubai":                       "Dubai",
	"Europe/Dublin":                    "Dublin",
	"Antarctica/DumontDUrville":        "DumontDUrville",
	"Asia/Dushanbe":                    "Dushanbe",
	"Pacific/Easter":                   "Easter",
	"EAT":                              "Eastern African Time",
	"ART":                              "Eastern European Time",
	"EET":                              "Eastern European Time",
	"Egypt":                            "Eastern European Time",
	"EST":                              "Eastern Standard Time",
	"America/Edmonton":                 "Edmonton",
	"Pacific/Efate":                    "Efate",
	"America/Eirunepe":                 "Eirunepe",
	"Africa/El_Aaiun":                  "El Aaiun",
	"America/El_Salvador":              "El Salvador",
	"Pacific/Enderbury":                "Enderbury",
	"America/Ensenada":                 "Ensenada",
	"Australia/Eucla":                  "Eucla",
	"Atlantic/Faeroe":                  "Faeroe",
	"Pacific/Fakaofo":                  "Fakaofo",
	"Atlantic/Faroe":                   "Faroe",
	"Pacific/Fiji":                     "Fiji",
	"America/Fort_Wayne":               "Fort Wayne",
	"America/Fortaleza":                "Fortaleza",
	"Africa/Freetown":                  "Freetown",
	"Pacific/Funafuti":                 "Funafuti",
	"Africa/Gaborone":                  "Gaborone",
	"Pacific/Galapagos":                "Galapagos",
	"Pacific/Gambier":                  "Gambier",
	"Asia/Gaza":                        "Gaza",
	"Europe/Gibraltar":                 "Gibraltar",
	"America/Glace_Bay":                "Glace Bay",
	"America/Godthab":                  "Godthab",
	"America/Goose_Bay":                "Goose Bay",
	"America/Grand_Turk":               "Grand Turk",
	"GMT":                              "Greenwich Mean Time",
	"GMT0":                             "GMT+00:00",
	"Etc/GMT":                          "Greenwich Mean Time",
	"Etc/GMT-1":                        "GMT+01:00",
	"Etc/GMT-2":                        "GMT+02:00",
	"Etc/GMT-3":                        "GMT+03:00",
	"Etc/GMT-4":                        "GMT+04:00",
	"Etc/GMT-5":                        "GMT+05:00",
	"Etc/GMT-6":                        "GMT+06:00",
	"Etc/GMT-7":                        "GMT+07:00",
	"Etc/GMT-8":                        "GMT+08:00",
	"Etc/GMT-9":                        "GMT+09:00",
	"Etc/GMT-10":                       "GMT+10:00",
	"Etc/GMT-11":                       "GMT+11:00",
	"Etc/GMT-12":                       "GMT+12:00",
	"Etc/GMT-13":                       "GMT+13:00",
	"Etc/GMT-14":                       "GMT+14:00",
	"Etc/GMT+1":                        "GMT-01:00",
	"Etc/GMT+2":                        "GMT-02:00",
	"Etc/GMT+3":                        "GMT-03:00",
	"Etc/GMT+4":                        "GMT-04:00",
	"Etc/GMT+5":                        "GMT-05:00",
	"Etc/GMT+6":                        "GMT-06:00",
	"Etc/GMT+7":                        "GMT-07:00",
	"Etc/GMT+8":                        "GMT-08:00",
	"Etc/GMT+9":                        "GMT-09:00",
	"Etc/GMT+10":                       "GMT-10:00",
	"Etc/GMT+11":                       "GMT-11:00",
	"Etc/GMT+12":                       "GMT-12:00",
	"Greenwich":                        "Greenwich Mean Time",
	"Iceland":                          "Greenwich Mean Time",
	"America/Grenada":                  "Grenada",
	"Pacific/Guadalcanal":              "Guadalcanal",
	"America/Guadeloupe":               "Guadeloupe",
	"Pacific/Guam":                     "Guam",
	"America/Guatemala":                "Guatemala",
	"America/Guayaquil":                "Guayaquil",
	"Europe/Guernsey":                  "Guernsey",
	"America/Guyana":                   "Guyana",
	"America/Halifax":                  "Halifax",
	"Africa/Harare":                    "Harare",
	"Asia/Harbin":                      "Harbin",
	"America/Havana":                   "Havana",
	"Europe/Helsinki":                  "Helsinki",
	"America/Hermosillo":               "Hermosillo",
	"Asia/Ho_Chi_Minh":                 "Ho Chi Minh",
	"Australia/Hobart":                 "Hobart",
	"Asia/Hong_Kong":                   "Hong Kong",
	"Pacific/Honolulu":                 "Honolulu",
	"Asia/Hovd":                        "Hovd",
	"HST":                              "Hawaii Standard Time",
	"America/Indiana/Indianapolis":     "Indiana - Indianapolis",
	"America/Indiana/Knox":             "Indiana - Knox",
	"America/Indiana/Marengo":          "Indiana - Marengo",
	"America/Indiana/Petersburg":       "Indiana - Petersburg",
	"America/Indiana/Tell_City":        "Indiana - Tell City",
	"America/Indiana/Vevay":            "Indiana - Vevay",
	"America/Indiana/Vincennes":        "Indiana - Vincennes",
	"America/Indiana/Winamac":          "Indiana - Winamac",
	"America/Indianapolis":             "Indianapolis",
	"VST":                              "Indochina Time",
	"America/Inuvik":                   "Inuvik",
	"America/Iqaluit":                  "Iqaluit",
	"Iran":                             "Iran Standard Time",
	"Asia/Irkutsk":                     "Irkutsk",
	"Europe/Isle_of_Man":               "Isle of Man",
	"Israel":                           "Israel Standard Time",
	"IST":                              "India Standard Time",
	"Asia/Istanbul":                    "Istanbul",
	"Europe/Istanbul":                  "Istanbul",
	"Asia/Jakarta":                     "Jakarta",
	"America/Jamaica":                  "Jamaica",
	"Atlantic/Jan_Mayen":               "Jan Mayen",
	"Japan":                            "Japan Standard Time",
	"JST":                              "Japan Standard Time",
	"Asia/Jayapura":                    "Jayapura",
	"Europe/Jersey":                    "Jersey",
	"Asia/Jerusalem":                   "Jerusalem",
	"Africa/Johannesburg":              "Johannesburg",
	"Pacific/Johnston":                 "Johnston",
	"America/Jujuy":                    "Jujuy",
	"America/Juneau":                   "Juneau",
	"Asia/Kabul":                       "Kabul",
	"Europe/Kaliningrad":               "Kaliningrad",
	"Asia/Kamchatka":                   "Kamchatka",
	"Africa/Kampala":                   "Kampala",
	"Asia/Karachi":                     "Karachi",
	"Asia/Kashgar":                     "Kashgar",
	"Asia/Kathmandu":                   "Kathmandu",
	"Asia/Katmandu":                    "Katmandu",
	"America/Kentucky/Louisville":      "Kentucky - Louisville",
	"America/Kentucky/Monticello":      "Kentucky - Monticello",
	"Indian/Kerguelen":                 "Kerguelen",
	"Africa/Khartoum":                  "Khartoum",
	"Europe/Kiev":                      "Kiev",
	"Africa/Kigali":                    "Kigali",
	"Africa/Kinshasa":                  "Kinshasa",
	"Pacific/Kiritimati":               "Kiritimati",
	"America/Knox_IN":                  "Knox IN",
	"Asia/Kolkata":                     "Kolkata",
	"ROK":                              "Korea Standard Time",
	"Pacific/Kosrae":                   "Kosrae",
	"Asia/Krasnoyarsk":                 "Krasnoyarsk",
	"Asia/Kuala_Lumpur":                "Kuala Lumpur",
	"Asia/Kuching":                     "Kuching",
	"Asia/Kuwait":                      "Kuwait",
	"Pacific/Kwajalein":                "Kwajalein",
	"Australia/LHI":                    "LHI",
	"America/La_Paz":                   "La Paz",
	"Africa/Lagos":                     "Lagos",
	"Africa/Libreville":                "Libreville",
	"America/Lima":                     "Lima",
	"Australia/Lindeman":               "Lindeman",
	"Europe/Lisbon":                    "Lisbon",
	"Europe/Ljubljana":                 "Ljubljana",
	"Africa/Lome":                      "Lome",
	"Europe/London":                    "London",
	"Arctic/Longyearbyen":              "Longyearbyen",
	"Australia/Lord_Howe":              "Lord Howe",
	"America/Los_Angeles":              "Los Angeles",
	"America/Louisville":               "Louisville",
	"Africa/Luanda":                    "Luanda",
	"Africa/Lubumbashi":                "Lubumbashi",
	"Africa/Lusaka":                    "Lusaka",
	"Europe/Luxembourg":                "Luxembourg",
	"Asia/Macao":                       "Macao",
	"Asia/Macau":                       "Macau",
	"America/Maceio":                   "Maceio",
	"Antarctica/Macquarie":             "Macquarie",
	"Atlantic/Madeira":                 "Madeira",
	"Europe/Madrid":                    "Madrid",
	"Asia/Magadan":                     "Magadan",
	"Indian/Mahe":                      "Mahe",
	"Pacific/Majuro":                   "Majuro",
	"Asia/Makassar":                    "Makassar",
	"Africa/Malabo":                    "Malabo",
	"Indian/Maldives":                  "Maldives",
	"Europe/Malta":                     "Malta",
	"America/Managua":                  "Managua",
	"America/Manaus":                   "Manaus",
	"Asia/Manila":                      "Manila",
	"Africa/Maputo":                    "Maputo",
	"Europe/Mariehamn":                 "Mariehamn",
	"America/Marigot":                  "Marigot",
	"Pacific/Marquesas":                "Marquesas",
	"America/Martinique":               "Martinique",
	"Africa/Maseru":                    "Maseru",
	"America/Matamoros":                "Matamoros",
	"Indian/Mauritius":                 "Mauritius",
	"Antarctica/Mawson":                "Mawson",
	"Indian/Mayotte":                   "Mayotte",
	"America/Mazatlan":                 "Mazatlan",
	"Africa/Mbabane":                   "Mbabane",
	"Antarctica/McMurdo":               "McMurdo",
	"Australia/Melbourne":              "Melbourne",
	"America/Mendoza":                  "Mendoza",
	"America/Menominee":                "Menominee",
	"America/Merida":                   "Merida",
	"America/Mexico_City":              "Mexico City",
	"MET":                              "Middle Europe Time",
	"Pacific/Midway":                   "Midway",
	"Europe/Minsk":                     "Minsk",
	"America/Miquelon":                 "Miquelon",
	"Africa/Mogadishu":                 "Mogadishu",
	"Europe/Monaco":                    "Monaco",
	"America/Moncton":                  "Moncton",
	"Africa/Monrovia":                  "Monrovia",
	"America/Monterrey":                "Monterrey",
	"America/Montevideo":               "Montevideo",
	"America/Montreal":                 "Montreal",
	"America/Montserrat":               "Montserrat",
	"Europe/Moscow":                    "Moscow",
	"W-SU":                             "Moscow Standard Time",
	"MST":                              "Mountain Standard Time",
	"PNT":                              "Mountain Standard Time",
	"Asia/Muscat":                      "Muscat",
	"Australia/NSW":                    "NSW",
	"Africa/Nairobi":                   "Nairobi",
	"America/Nassau":                   "Nassau",
	"Pacific/Nauru":                    "Nauru",
	"Africa/Ndjamena":                  "Ndjamena",
	"CNT":                              "Newfoundland Standard Time",
	"America/New_York":                 "New York",
	"NST":                              "New Zealand Standard Time",
	"NZ":                               "New Zealand Standard Time",
	"Africa/Niamey":                    "Niamey",
	"Asia/Nicosia":                     "Nicosia",
	"Europe/Nicosia":                   "Nicosia",
	"America/Nipigon":                  "Nipigon",
	"Pacific/Niue":                     "Niue",
	"America/Nome":                     "Nome",
	"Pacific/Norfolk":                  "Norfolk",
	"America/Noronha":                  "Noronha",
	"Australia/North":                  "North",
	"America/North_Dakota/Center":      "North Dakota - Center",
	"America/North_Dakota/New_Salem":   "North Dakota - New Salem",
	"Africa/Nouakchott":                "Nouakchott",
	"Pacific/Noumea":                   "Noumea",
	"Asia/Novokuznetsk":                "Novokuznetsk",
	"Asia/Novosibirsk":                 "Novosibirsk",
	"America/Ojinaga":                  "Ojinaga",
	"Asia/Omsk":                        "Omsk",
	"Asia/Oral":                        "Oral",
	"Europe/Oslo":                      "Oslo",
	"Africa/Ouagadougou":               "Ouagadougou",
	"PST":                              "Pacific Standard Time",
	"Pacific/Pago_Pago":                "Pago Pago",
	"PLT":                              "Pakistan Time",
	"Pacific/Palau":                    "Palau",
	"Antarctica/Palmer":                "Palmer",
	"America/Panama":                   "Panama",
	"America/Pangnirtung":              "Pangnirtung",
	"America/Paramaribo":               "Paramaribo",
	"Europe/Paris":                     "Paris",
	"Australia/Perth":                  "Perth",
	"Asia/Phnom_Penh":                  "Phnom Penh",
	"America/Phoenix":                  "Phoenix",
	"Pacific/Pitcairn":                 "Pitcairn",
	"Europe/Podgorica":                 "Podgorica",
	"Pacific/Ponape":                   "Ponape",
	"Asia/Pontianak":                   "Pontianak",
	"Pacific/Port_Moresby":             "Port Moresby",
	"America/Port_of_Spain":            "Port of Spain",
	"America/Port-au-Prince":           "Port-au-Prince",
	"America/Porto_Acre":               "Porto Acre",
	"America/Porto_Velho":              "Porto Velho",
	"Africa/Porto-Novo":                "Porto-Novo",
	"Europe/Prague":                    "Prague",
	"America/Puerto_Rico":              "Puerto Rico",
	"Asia/Pyongyang":                   "Pyongyang",
	"Asia/Qatar":                       "Qatar",
	"Australia/Queensland":             "Queensland",
	"Asia/Qyzylorda":                   "Qyzylorda",
	"America/Rainy_River":              "Rainy River",
	"Asia/Rangoon":                     "Rangoon",
	"America/Rankin_Inlet":             "Rankin Inlet",
	"Pacific/Rarotonga":                "Rarotonga",
	"America/Recife":                   "Recife",
	"America/Regina":                   "Regina",
	"America/Resolute":                 "Resolute",
	"Indian/Reunion":                   "Reunion",
	"Atlantic/Reykjavik":               "Reykjavik",
	"Europe/Riga":                      "Riga",
	"America/Rio_Branco":               "Rio Branco",
	"Asia/Riyadh":                      "Riyadh",
	"Europe/Rome":                      "Rome",
	"America/Rosario":                  "Rosario",
	"Antarctica/Rothera":               "Rothera",
	"Asia/Saigon":                      "Saigon",
	"Pacific/Saipan":                   "Saipan",
	"Asia/Sakhalin":                    "Sakhalin",
	"Europe/Samara":                    "Samara",
	"Asia/Samarkand":                   "Samarkand",
	"Pacific/Samoa":                    "Samoa",
	"Europe/San_Marino":                "San Marino",
	"America/Santa_Isabel":             "Santa Isabel",
	"America/Santarem":                 "Santarem",
	"America/Santiago":                 "Santiago",
	"America/Santo_Domingo":            "Santo Domingo",
	"America/Sao_Paulo":                "Sao Paulo",
	"Africa/Sao_Tome":                  "Sao Tome",
	"Europe/Sarajevo":                  "Sarajevo",
	"America/Scoresbysund":             "Scoresbysund",
	"Asia/Seoul":                       "Seoul",
	"Asia/Shanghai":                    "Shanghai",
	"America/Shiprock":                 "Shiprock",
	"Europe/Simferopol":                "Simferopol",
	"Asia/Singapore":                   "Singapore",
	"Singapore":                        "Singapore Time",
	"Europe/Skopje":                    "Skopje",
	"Europe/Sofia":                     "Sofia",
	"SST":                              "Solomon Is. Time",
	"Australia/South":                  "South",
	"Atlantic/South_Georgia":           "South Georgia",
	"Antarctica/South_Pole":            "South Pole",
	"America/St_Barthelemy":            "St Barthelemy",
	"Atlantic/St_Helena":               "St Helena",
	"America/St_Johns":                 "St Johns",
	"America/St_Kitts":                 "St Kitts",
	"America/St_Lucia":                 "St Lucia",
	"America/St_Thomas":                "St Thomas",
	"America/St_Vincent":               "St Vincent",
	"Atlantic/Stanley":                 "Stanley",
	"Europe/Stockholm":                 "Stockholm",
	"America/Swift_Current":            "Swift Current",
	"Australia/Sydney":                 "Sydney",
	"Antarctica/Syowa":                 "Syowa",
	"Pacific/Tahiti":                   "Tahiti",
	"Asia/Taipei":                      "Taipei",
	"Europe/Tallinn":                   "Tallinn",
	"Pacific/Tarawa":                   "Tarawa",
	"Asia/Tashkent":                    "Tashkent",
	"Australia/Tasmania":               "Tasmania",
	"Asia/Tbilisi":                     "Tbilisi",
	"America/Tegucigalpa":              "Tegucigalpa",
	"Asia/Tehran":                      "Tehran",
	"Asia/Tel_Aviv":                    "Tel Aviv",
	"Asia/Thimbu":                      "Thimbu",
	"Asia/Thimphu":                     "Thimphu",
	"America/Thule":                    "Thule",
	"America/Thunder_Bay":              "Thunder Bay",
	"America/Tijuana":                  "Tijuana",
	"Africa/Timbuktu":                  "Timbuktu",
	"Europe/Tirane":                    "Tirane",
	"Europe/Tiraspol":                  "Tiraspol",
	"Asia/Tokyo":                       "Tokyo",
	"Pacific/Tongatapu":                "Tongatapu",
	"America/Toronto":                  "Toronto",
	"America/Tortola":                  "Tortola",
	"Africa/Tripoli":                   "Tripoli",
	"Pacific/Truk":                     "Truk",
	"Africa/Tunis":                     "Tunis",
	"Asia/Ujung_Pandang":               "Ujung Pandang",
	"Asia/Ulaanbaatar":                 "Ulaanbaatar",
	"Asia/Ulan_Bator":                  "Ulan Bator",
	"Asia/Urumqi":                      "Urumqi",
	"Europe/Uzhgorod":                  "Uzhgorod",
	"Europe/Vaduz":                     "Vaduz",
	"America/Vancouver":                "Vancouver",
	"Europe/Vatican":                   "Vatican",
	"Australia/Victoria":               "Victoria",
	"Europe/Vienna":                    "Vienna",
	"Asia/Vientiane":                   "Vientiane",
	"Europe/Vilnius":                   "Vilnius",
	"America/Virgin":                   "Virgin",
	"Asia/Vladivostok":                 "Vladivostok",
	"Europe/Volgograd":                 "Volgograd",
	"Antarctica/Vostok":                "Vostok",
	"Pacific/Wake":                     "Wake",
	"Pacific/Wallis":                   "Wallis",
	"Europe/Warsaw":                    "Warsaw",
	"Australia/West":                   "West",
	"MIT":                              "West Samoa Standard Time",
	"Portugal":                         "Western European Time",
	"WET":                              "Western European Time",
	"America/Whitehorse":               "Whitehorse",
	"Africa/Windhoek":                  "Windhoek",
	"America/Winnipeg":                 "Winnipeg",
	"America/Yakutat":                  "Yakutat",
	"Asia/Yakutsk":                     "Yakutsk",
	"Australia/Yancowinna":             "Yancowinna",
	"Pacific/Yap":                      "Yap",
	"Asia/Yekaterinburg":               "Yekaterinburg",
	"America/Yellowknife":              "Yellowknife",
	"Asia/Yerevan":                     "Yerevan",
	"Europe/Zagreb":                    "Zagreb",
	"Europe/Zaporozhye":                "Zaporozhye",
	"Europe/Zurich":                    "Zurich",
	"UTC":                              "UTC",
	"UTC-12":                           "UTC-12",
	"UTC-11":                           "UTC-11",
	"UTC-10":                           "UTC-10",
	"UTC-9":                            "UTC-9",
	"UTC-8":                            "UTC-8",
	"UTC-7":                            "UTC-7",
	"UTC-6":                            "UTC-6",
	"UTC-5":                            "UTC-5",
	"UTC-4":                            "UTC-4",
	"UTC-3":                            "UTC-3",
	"UTC-2":                            "UTC-2",
	"UTC-1":                            "UTC-1",
	"UTC+0":                            "UTC+0",
	"UTC+1":                            "UTC+1",
	"UTC+2":                            "UTC+2",
	"UTC+3":                            "UTC+3",
	"UTC+4":                            "UTC+4",
	"UTC+5":                            "UTC+5",
	"UTC+6":                            "UTC+6",
	"UTC+7":                            "UTC+7",
	"UTC+8":                            "UTC+8",
	"UTC+9":                            "UTC+9",
	"UTC+10":                           "UTC+10",
	"UTC+11":                           "UTC+11",
	"UTC+12":                           "UTC+12",
	"UTC+13":                           "UTC+13",
	"UTC+14":                           "UTC+14",
}

// New returns a new Timezone.
func New(logger *jsonlog.Logger) *Timezone {
	return &Timezone{
		Logger:    logger,
		Timezones: timezone,
	}
}

// TimezoneToString returns the normalized tz.
func (t *Timezone) TimezoneToString(tz string) (string, error) {
	t.Logger.PrintDebug("Timezone.TimezoneToString()", map[string]string{
		"tz": tz,
	})

	if tz == "" {
		return "", ErrTimezoneIsEmpty
	}

	if _, ok := t.Timezones[tz]; !ok {
		return "", ErrNoMatchingTimezone
	}

	v := t.Timezones[tz]

	t.Logger.PrintDebug("Timezone.TimezoneToString()", map[string]string{
		"timezone": v,
	})

	return v, nil
}
