#!/usr/bin/env python3
# country names and codes in python3 format
# executing this script outputs the list in json and html table format.
# source: https://www.worldatlas.com/aatlas/ctycodes.htm

import os

countryDetails = (
  "COUNTRY", "A2 (ISO)", "A3 (UN)", "NUM (UN)", "DIALING CODE")

countries = [
  ("Afghanistan", "AF", "AFG", "4", "93"),
  ("Albania", "AL", "ALB", "8", "355"),
  ("Algeria", "DZ", "DZA", "12", "213"),
  ("American Samoa", "AS", "ASM", "16", "1-684"),
  ("Andorra", "AD", "AND", "20", "376"),
  ("Angola", "AO", "AGO", "24", "244"),
  ("Anguilla", "AI", "AIA", "660", "1-264"),
  ("Antarctica", "AQ", "ATA", "10", "672"),
  ("Antigua and Barbuda", "AG", "ATG", "28", "1-268"),
  ("Argentina", "AR", "ARG", "32", "54"),
  ("Armenia", "AM", "ARM", "51", "374"),

  ("Aruba", "AW", "ABW", "533", "297"),
  ("Australia", "AU", "AUS", "36", "61"),
  ("Austria", "AT", "AUT", "40", "43"),
  ("Azerbaijan", "AZ", "AZE", "31", "994"),
  ("Bahamas", "BS", "BHS", "44", "1-242"),
  ("Bahrain", "BH", "BHR", "48", "973"),
  ("Bangladesh", "BD", "BGD", "50", "880"),
  ("Barbados", "BB", "BRB", "52", "1-246"),
  ("Belarus", "BY", "BLR", "112", "375"),
  ("Belgium", "BE", "BEL", "56", "32"),

  ("Belize", "BZ", "BLZ", "84", "501"),
  ("Benin", "BJ", "BEN", "204", "229"),
  ("Bermuda", "BM", "BMU", "60", "1-441"),
  ("Bhutan", "BT", "BTN", "64", "975"),
  ("Bolivia", "BO", "BOL", "68", "591"),
  ("Bonaire", "BQ", "BES", "535", "599"),
  ("Bosnia and Herzegovina", "BA", "BIH", "70", "387"),
  ("Botswana", "BW", "BWA", "72", "267"),
  ("Bouvet Island", "BV", "BVT", "74", "47"),
  ("Brazil", "BR", "BRA", "76", "55"),

  ("British Indian Ocean Territory", "IO", "IOT", "86", "246"),
  ("Brunei Darussalam", "BN", "BRN", "96", "673"),
  ("Bulgaria", "BG", "BGR", "100", "359"),
  ("Burkina Faso", "BF", "BFA", "854", "226"),
  ("Burundi", "BI", "BDI", "108", "257"),
  ("Cambodia", "KH", "KHM", "116", "855"),
  ("Cameroon", "CM", "CMR", "120", "237"),
  ("Canada", "CA", "CAN", "124", "1"),
  ("Cape Verde", "CV", "CPV", "132", "238"),
  ("Cayman Islands", "KY", "CYM", "136", "1-345"),

  ("Central African Republic", "CF", "CAF", "140", "236"),
  ("Chad", "TD", "TCD", "148", "235"),
  ("Chile", "CL", "CHL", "152", "56"),
  ("China", "CN", "CHN", "156", "86"),
  ("Christmas Island", "CX", "CXR", "162", "61"),
  ("Cocos (Keeling) Islands", "CC", "CCK", "166", "61"),
  ("Colombia", "CO", "COL", "170", "57"),
  ("Comoros", "KM", "COM", "174", "269"),
  ("Congo", "CG", "COG", "178", "242"),
  ("Democratic Republic of the Congo", "CD", "COD", "180", "243"),

  ("Cook Islands", "CK", "COK", "184", "682"),
  ("Costa Rica", "CR", "CRI", "188", "506"),
  ("Croatia", "HR", "HRV", "191", "385"),
  ("Cuba", "CU", "CUB", "192", "53"),
  ("Curacao", "CW", "CUW", "531", "599"),
  ("Cyprus", "CY", "CYP", "196", "357"),
  ("Czech Republic", "CZ", "CZE", "203", "420"),
  ("Cote d'Ivoire", "CI", "CIV", "384", "225"),
  ("Denmark", "DK", "DNK", "208", "45"),
  ("Djibouti", "DJ", "DJI", "262", "253"),

  ("Dominica", "DM", "DMA", "212", "1-767"),
  ("Dominican Republic", "DO", "DOM", "214", "1-809,1-829,1-849"),
  ("Ecuador", "EC", "ECU", "218", "593"),
  ("Egypt", "EG", "EGY", "818", "20"),
  ("El Salvador", "SV", "SLV", "222", "503"),
  ("Equatorial Guinea", "GQ", "GNQ", "226", "240"),
  ("Eritrea", "ER", "ERI", "232", "291"),
  ("Estonia", "EE", "EST", "233", "372"),
  ("Ethiopia", "ET", "ETH", "231", "251"),
  ("Falkland Islands (Malvinas)", "FK", "FLK", "238", "500"),

  ("Faroe Islands", "FO", "FRO", "234", "298"),
  ("Fiji", "FJ", "FJI", "242", "679"),
  ("Finland", "FI", "FIN", "246", "358"),
  ("France", "FR", "FRA", "250", "33"),
  ("French Guiana", "GF", "GUF", "254", "594"),
  ("French Polynesia", "PF", "PYF", "258", "689"),
  ("French Southern Territories", "TF", "ATF", "260", "262"),
  ("Gabon", "GA", "GAB", "266", "241"),
  ("Gambia", "GM", "GMB", "270", "220"),
  ("Georgia", "GE", "GEO", "268", "995"),

  ("Germany", "DE", "DEU", "276", "49"),
  ("Ghana", "GH", "GHA", "288", "233"),
  ("Gibraltar", "GI", "GIB", "292", "350"),
  ("Greece", "GR", "GRC", "300", "30"),
  ("Greenland", "GL", "GRL", "304", "299"),
  ("Grenada", "GD", "GRD", "308", "1-473"),
  ("Guadeloupe", "GP", "GLP", "312", "590"),
  ("Guam", "GU", "GUM", "316", "1-671"),
  ("Guatemala", "GT", "GTM", "320", "502"),
  ("Guernsey", "GG", "GGY", "831", "44"),

  ("Guinea", "GN", "GIN", "324", "224"),
  ("Guinea-Bissau", "GW", "GNB", "624", "245"),
  ("Guyana", "GY", "GUY", "328", "592"),
  ("Haiti", "HT", "HTI", "332", "509"),
  ("Heard Island and McDonald Islands", "HM", "HMD", "334", "672"),
  ("Holy See (Vatican City State)", "VA", "VAT", "336", "379"),
  ("Honduras", "HN", "HND", "340", "504"),
  ("Hong Kong", "HK", "HKG", "344", "852"),
  ("Hungary", "HU", "HUN", "348", "36"),
  ("Iceland", "IS", "ISL", "352", "354"),

  ("India", "IN", "IND", "356", "91"),
  ("Indonesia", "ID", "IDN", "360", "62"),
  ("Iran, Islamic Republic of", "IR", "IRN", "364", "98"),
  ("Iraq", "IQ", "IRQ", "368", "964"),
  ("Ireland", "IE", "IRL", "372", "353"),
  ("Isle of Man", "IM", "IMN", "833", "44"),
  ("Israel", "IL", "ISR", "376", "972"),
  ("Italy", "IT", "ITA", "380", "39"),
  ("Jamaica", "JM", "JAM", "388", "1-876"),
  ("Japan", "JP", "JPN", "392", "81"),

  ("Jersey", "JE", "JEY", "832", "44"),
  ("Jordan", "JO", "JOR", "400", "962"),
  ("Kazakhstan", "KZ", "KAZ", "398", "7"),
  ("Kenya", "KE", "KEN", "404", "254"),
  ("Kiribati", "KI", "KIR", "296", "686"),
  ("Korea, Democratic People's Republic of", "KP", "PRK", "408", "850"),
  ("Korea, Republic of", "KR", "KOR", "410", "82"),
  ("Kuwait", "KW", "KWT", "414", "965"),
  ("Kyrgyzstan", "KG", "KGZ", "417", "996"),
  ("Lao People's Democratic Republic", "LA", "LAO", "418", "856"),

  ("Latvia", "LV", "LVA", "428", "371"),
  ("Lebanon", "LB", "LBN", "422", "961"),
  ("Lesotho", "LS", "LSO", "426", "266"),
  ("Liberia", "LR", "LBR", "430", "231"),
  ("Libya", "LY", "LBY", "434", "218"),
  ("Liechtenstein", "LI", "LIE", "438", "423"),
  ("Lithuania", "LT", "LTU", "440", "370"),
  ("Luxembourg", "LU", "LUX", "442", "352"),
  ("Macao", "MO", "MAC", "446", "853"),
  ("Macedonia, the Former Yugoslav Republic of", "MK", "MKD", "807", "389"),

  ("Madagascar", "MG", "MDG", "450", "261"),
  ("Malawi", "MW", "MWI", "454", "265"),
  ("Malaysia", "MY", "MYS", "458", "60"),
  ("Maldives", "MV", "MDV", "462", "960"),
  ("Mali", "ML", "MLI", "466", "223"),
  ("Malta", "MT", "MLT", "470", "356"),
  ("Marshall Islands", "MH", "MHL", "584", "692"),
  ("Martinique", "MQ", "MTQ", "474", "596"),
  ("Mauritania", "MR", "MRT", "478", "222"),
  ("Mauritius", "MU", "MUS", "480", "230"),

  ("Mayotte", "YT", "MYT", "175", "262"),
  ("Mexico", "MX", "MEX", "484", "52"),
  ("Micronesia, Federated States of", "FM", "FSM", "583", "691"),
  ("Moldova, Republic of", "MD", "MDA", "498", "373"),
  ("Monaco", "MC", "MCO", "492", "377"),
  ("Mongolia", "MN", "MNG", "496", "976"),
  ("Montenegro", "ME", "MNE", "499", "382"),
  ("Montserrat", "MS", "MSR", "500", "1-664"),
  ("Morocco", "MA", "MAR", "504", "212"),
  ("Mozambique", "MZ", "MOZ", "508", "258"),

  ("Myanmar", "MM", "MMR", "104", "95"),
  ("Namibia", "NA", "NAM", "516", "264"),
  ("Nauru", "NR", "NRU", "520", "674"),
  ("Nepal", "NP", "NPL", "524", "977"),
  ("Netherlands", "NL", "NLD", "528", "31"),
  ("New Caledonia", "NC", "NCL", "540", "687"),
  ("New Zealand", "NZ", "NZL", "554", "64"),
  ("Nicaragua", "NI", "NIC", "558", "505"),
  ("Niger", "NE", "NER", "562", "227"),
  ("Nigeria", "NG", "NGA", "566", "234"),

  ("Niue", "NU", "NIU", "570", "683"),
  ("Norfolk Island", "NF", "NFK", "574", "672"),
  ("Northern Mariana Islands", "MP", "MNP", "580", "1-670"),
  ("Norway", "NO", "NOR", "578", "47"),
  ("Oman", "OM", "OMN", "512", "968"),
  ("Pakistan", "PK", "PAK", "586", "92"),
  ("Palau", "PW", "PLW", "585", "680"),
  ("Palestine, State of", "PS", "PSE", "275", "970"),
  ("Panama", "PA", "PAN", "591", "507"),
  ("Papua New Guinea", "PG", "PNG", "598", "675"),

  ("Paraguay", "PY", "PRY", "600", "595"),
  ("Peru", "PE", "PER", "604", "51"),
  ("Philippines", "PH", "PHL", "608", "63"),
  ("Pitcairn", "PN", "PCN", "612", "870"),
  ("Poland", "PL", "POL", "616", "48"),
  ("Portugal", "PT", "PRT", "620", "351"),
  ("Puerto Rico", "PR", "PRI", "630", "1"),
  ("Qatar", "QA", "QAT", "634", "974"),
  ("Romania", "RO", "ROU", "642", "40"),
  ("Russian Federation", "RU", "RUS", "643", "7"),

  ("Rwanda", "RW", "RWA", "646", "250"),
  ("Reunion", "RE", "REU", "638", "262"),
  ("Saint Barthelemy", "BL", "BLM", "652", "590"),
  ("Saint Helena", "SH", "SHN", "654", "290"),
  ("Saint Kitts and Nevis", "KN", "KNA", "659", "1-869"),
  ("Saint Lucia", "LC", "LCA", "662", "1-758"),
  ("Saint Martin (French part)", "MF", "MAF", "663", "590"),
  ("Saint Pierre and Miquelon", "PM", "SPM", "666", "508"),
  ("Saint Vincent and the Grenadines", "VC", "VCT", "670", "1-784"),
  ("Samoa", "WS", "WSM", "882", "685"),

  ("San Marino", "SM", "SMR", "674", "378"),
  ("Sao Tome and Principe", "ST", "STP", "678", "239"),
  ("Saudi Arabia", "SA", "SAU", "682", "966"),
  ("Senegal", "SN", "SEN", "686", "221"),
  ("Serbia", "RS", "SRB", "688", "381"),
  ("Seychelles", "SC", "SYC", "690", "248"),
  ("Sierra Leone", "SL", "SLE", "694", "232"),
  ("Singapore", "SG", "SGP", "702", "65"),
  ("Sint Maarten (Dutch part)", "SX", "SXM", "534", "1-721"),
  ("Slovakia", "SK", "SVK", "703", "421"),

  ("Slovenia", "SI", "SVN", "705", "386"),
  ("Solomon Islands", "SB", "SLB", "90", "677"),
  ("Somalia", "SO", "SOM", "706", "252"),
  ("South Africa", "ZA", "ZAF", "710", "27"),
  ("South Georgia and the South Sandwich Islands", "GS", "SGS", "239", "500"),
  ("South Sudan", "SS", "SSD", "728", "211"),
  ("Spain", "ES", "ESP", "724", "34"),
  ("Sri Lanka", "LK", "LKA", "144", "94"),
  ("Sudan", "SD", "SDN", "729", "249"),
  ("Suriname", "SR", "SUR", "740", "597"),

  ("Svalbard and Jan Mayen", "SJ", "SJM", "744", "47"),
  ("Swaziland", "SZ", "SWZ", "748", "268"),
  ("Sweden", "SE", "SWE", "752", "46"),
  ("Switzerland", "CH", "CHE", "756", "41"),
  ("Syrian Arab Republic", "SY", "SYR", "760", "963"),
  ("Taiwan", "TW", "TWN", "158", "886"),
  ("Tajikistan", "TJ", "TJK", "762", "992"),
  ("United Republic of Tanzania", "TZ", "TZA", "834", "255"),
  ("Thailand", "TH", "THA", "764", "66"),
  ("Timor-Leste", "TL", "TLS", "626", "670"),
  ("Togo", "TG", "TGO", "768", "228"),

  ("Tokelau", "TK", "TKL", "772", "690"),
  ("Tonga", "TO", "TON", "776", "676"),
  ("Trinidad and Tobago", "TT", "TTO", "780", "1-868"),
  ("Tunisia", "TN", "TUN", "788", "216"),
  ("Turkey", "TR", "TUR", "792", "90"),
  ("Turkmenistan", "TM", "TKM", "795", "993"),
  ("Turks and Caicos Islands", "TC", "TCA", "796", "1-649"),
  ("Tuvalu", "TV", "TUV", "798", "688"),
  ("Uganda", "UG", "UGA", "800", "256"),
  ("Ukraine", "UA", "UKR", "804", "380"),

  ("United Arab Emirates", "AE", "ARE", "784", "971"),
  ("United Kingdom", "GB", "GBR", "826", "44"),
  ("United States", "US", "USA", "840", "1"),
  ("United States Minor Outlying Islands", "UM", "UMI", "581", "1"),
  ("Uruguay", "UY", "URY", "858", "598"),
  ("Uzbekistan", "UZ", "UZB", "860", "998"),
  ("Vanuatu", "VU", "VUT", "548", "678"),
  ("Venezuela", "VE", "VEN", "862", "58"),
  ("Viet Nam", "VN", "VNM", "704", "84"),
  ("British Virgin Islands", "VG", "VGB", "92", "1-284"),

  ("US Virgin Islands", "VI", "VIR", "850", "1-340"),
  ("Wallis and Futuna", "WF", "WLF", "876", "681"),
  ("Western Sahara", "EH", "ESH", "732", "212"),
  ("Yemen", "YE", "YEM", "887", "967"),
  ("Zambia", "ZM", "ZMB", "894", "260"),
  ("Zimbabwe", "ZW", "ZWE", "716", "263"),
]

def printJsonObject():
    print("// json")
    print("countries = {")
    print("// {}: {}".format(countryDetails[3], countryDetails))
    n = len(countries)
    count = 0
    for tup in countries:
        count += 1
        if count == n:
            print('"{}":{}'.format(tup[3], list(tup)))
        else:
            print('"{}":{},'.format(tup[3], list(tup)))

    print("};")

def printHtmlTable():
    count = 0
    print("<!-- html table for country names and its codes -->")
    print('<table border="1">')
    print("<tr>")
    print("  <th>S.No.</th>")
    for d in countryDetails:
        print("  <th>{}</th>".format(d))
    print("</tr>", os.linesep)
    for tup in countries:
        count += 1
        print("<tr>")
        print("  <td>{}</td>".format(count))
        for t in tup:
            print("  <td>{}</td>".format(t))
        print("</tr>", os.linesep)
    print("</table>")


if __name__ == "__main__":
    printJsonObject()
    print("\n\n")
    printHtmlTable()


