// Copyright 2014 Brian J. Downs
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openweathermap

var (
	baseURL      = "http://api.openweathermap.org/data/2.5/weather?%s"
	iconURL      = "http://openweathermap.org/img/w/%s"
	stationURL   = "http://api.openweathermap.org/data/2.5/station?id=%d"
	forecastBase = "http://api.openweathermap.org/data/2.5/forecast/daily?%s=%s&mode=json&units=%s&cnt=%d"
	dataPostURL  = "http://openweathermap.org/data/post"
	DataUnits    = map[string]string{"C": "metric", "F": "imperial", "K": "internal"}
	LangCodes    = map[string]string{
		"EN":    "English",
		"RU":    "Russian",
		"IT":    "Italian",
		"ES":    "Spanish",
		"SP":    "Spanish",
		"UK":    "Ukrainian",
		"UA":    "Ukrainian",
		"DE":    "German",
		"PT":    "Portuguese",
		"RO":    "Romanian",
		"PL":    "Polish",
		"FI":    "Finnish",
		"NL":    "Dutch",
		"FR":    "French",
		"BG":    "Bulgarian",
		"SV":    "Swedish",
		"SE":    "Swedish",
		"TR":    "Turkish",
		"HR":    "Croatian",
		"CA":    "Catalan",
		"ZH_TW": "Chinese Traditional",
		"ZH":    "Chinese Simplified",
		"ZH_CN": "Chinese Simplified",
	}
)

// Config will hold default settings to be passed into the
// "New{current, forecast, etc}" function.
type Config struct {
	Mode     string // User choice of JSON or XML
	Unit     string // F, C, or K
	APIKey   string // API Key for connecting to the OWM
	Username string // Username for posting data
	Password string // Pasword for posting data
}

// APIError returned on failed API calls.
type APIError struct {
	Message string `json:"message"`
	COD     string `json:"cod"`
}

// Coordinates struct holds longitude and latitude data
// in returned JSON or as parameter data for requests
// using longitude and latitude.
type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

// Sys struct contains general information about the request
// and the surrounding area for where the request was made.
type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// Wind struct contains the speed and degree of the wind.
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

// Weather struct holds high-level, basic info on the returned
// data.
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Main struct contains the temperates, humidity, pressure for
// the request.
type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressure float64 `json:"pressure"`
	Humidity int     `json:"humidity"`
}

// Clouds struct holds data regarding cloud cover.
type Clouds struct {
	All int `json:"all"`
}

// ValidDataUnit makes sure the string passed in is an accepted
// unit of measure to be used for the return data.
func ValidDataUnit(u string) bool {
	for d := range DataUnits {
		if u == d {
			return true
		}
	}
	return false
}

// ValidLangValue makes sure the string passed in is an
// acceptable lang code.
func ValidLangCode(c string) bool {
	for d := range LangCodes {
		if c == d {
			return true
		}
	}
	return false
}

// CheckAPIKeyExists will see if an API key has been set.
func (c *Config) CheckAPIKeyExists() bool { return len(c.APIKey) > 1 }
