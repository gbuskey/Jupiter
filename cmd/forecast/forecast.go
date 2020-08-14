/*
Copyright © 2020 Grant Buskey gbuskey@ncsu.edu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package forecast

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Used to report information about current weather data",
	Long: `Forcast is used to tell the user the current weather report,
or the predicted weather report in a certain area`,
	RunE: runE,
}

// NewForecastCmd returnns a new command line instancce
func NewForecastCmd() *cobra.Command {
	return forecastCmd
}

// set up all flagas
func init() {
	// set any default values for our flags
	viper.SetDefault("now", true)
	viper.SetDefault("five-day", false)

	forecastCmd.Flags().BoolP("now", "n", true, "Gets the current weather.")
	forecastCmd.Flags().BoolP("five-day", "f", false, "Gets the weather for the next five days.")
	forecastCmd.Flags().String("city", "", "The city you are interested in. Must provide city at a minimum.")
	forecastCmd.Flags().String("state-code", "", "ONLY VALID FOR US CITIES. The state your desired city resides in.")
	forecastCmd.Flags().String("country-code", "", "The country your state and city resides in.")
	forecastCmd.Flags().String("apikey", "", "REQUIRED. Must provide APIKey to fulfil request")

	// read in flags
	err := viper.BindPFlags(forecastCmd.Flags())
	if err != nil {
		fmt.Print(fmt.Errorf("error loading flags. error: %s", err))
	}
}
func runE(cmd *cobra.Command, args []string) error {
	// get a hold of all of our flags / env vars
	apikey := viper.GetString("APIKey")
	currentWeather := viper.GetBool("now")
	fiveDayForecast := viper.GetBool("five-day")
	city := viper.GetString("city")
	stateCode := viper.GetString("state-code")
	countryCode := viper.GetString("country-code")

	// build the parameter query string for the api request
	requestParams, err := getRequestParams(city, stateCode, countryCode)
	if err != nil {
		return err
	}

	if apikey == "" {
		return fmt.Errorf("apikey cannot be nil")
	}

	// Gets current weather by default, five day forecast if desired
	if currentWeather {
		err = getCurrentWeather(requestParams, apikey)
		if err != nil {
			return err
		}
	}

	if fiveDayForecast {
		err = getFiveDayForecast(requestParams, apikey)
		if err != nil {
			return err
		}
	}

	return nil
}

// getRequestParams returns the parameters minus an APIKey to execute a get request
// possible returns:
//	city
//  city,state,US
//	city,countryCode
func getRequestParams(city, stateCode, countryCode string) (string, error) {
	paramStr := ""

	// return if city isn't set
	if city == "" {
		return "", fmt.Errorf("city flag cannot be empty")
	}
	paramStr = city

	// If no country code provided just return the city
	if countryCode == "" {
		return paramStr, nil
	}

	// if the country is the us we try and add a state
	if countryCode == "US" && stateCode != "" {
		return paramStr + "," + stateCode + "," + countryCode, nil
	}
	return paramStr + "," + countryCode, nil

}
func makeRequest(requestUrl string) (string, error) {
	// make the request
	res, err := http.Get("https://" + requestUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// print out the response
	fmt.Println(res)
	fmt.Println(string(body))
	return "", nil
}

func getCurrentWeather(requestParams, apikey string) error {
	requestUrl := fmt.Sprintf("api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", requestParams, apikey)
	resp, err := makeRequest(requestUrl)

	if err != nil {
		return fmt.Errorf("invalid current weather request. error: %s", err)
	}
	fmt.Print(resp)
	return nil
}

func getFiveDayForecast(requestParams, apikey string) error {
	requestUrl := fmt.Sprintf("api.openweathermap.org/data/2.5/forecast?q=%sappid=%s", requestParams, apikey)
	resp, err := makeRequest(requestUrl)

	if err != nil {
		return fmt.Errorf("invalid five day request. error: %s", err)
	}
	fmt.Print(resp)
	return nil
}
