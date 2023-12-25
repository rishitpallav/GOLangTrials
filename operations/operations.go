package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAgeUsingName(name string) (string, error) {

	// EDA I guess
	name = strings.TrimSuffix(name, "\r\n")
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Name cannot be empty")
		return "", errors.New(fmt.Sprintln("The name entered is empty"))
	}

	// API call
	// We get all the data from https://api.agify.io/?name=YOUR_NAME
	apiURL := "https://api.agify.io/?name=" + name
	response, err := http.Get(apiURL)

	if err != nil {
		return "", errors.New(fmt.Sprintf("There was an error: %v", err))
	}
	// defer is used to close the response.Body after the function returns
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("API request failed with the status code: %d", response.StatusCode))
	}

	// We use an interface to store the data from the API
	// The interface is a map with string keys and interface values
	var result map[string]interface{}
	decoder := json.NewDecoder(response.Body)
	// We decode the response.Body into the interface
	// if there is an error, we return it
	if err := decoder.Decode(&result); err != nil {
		return "", errors.New(fmt.Sprintf("There was an error decoding the file: %v", err))
	}

	age := result["age"].(float64)

	return fmt.Sprintf("Predicted age for %s is %.2f\n", name, age), nil
}

func GetGenderUsingName(name string) (string, error) {

	name = strings.TrimSuffix(name, "\r\n")
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Name cannot be empty")
		return "", errors.New(fmt.Sprintln("The name entered is empty"))
	}

	name = strings.TrimSuffix(name, "\r\n")
	apiURL := "https://api.genderize.io/?name=" + name
	response, err := http.Get(apiURL)

	if err != nil {
		return "", errors.New(fmt.Sprintf("There was an error: %v", err))
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("API request failed with the status code: %d", response.StatusCode))
	}

	var result map[string]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		return "", errors.New(fmt.Sprintf("There was an error decoding the file: %v \n", err))
	}

	gender := result["gender"]
	return fmt.Sprintf("Predicted gender for %s is %s\n", name, gender), nil
}
