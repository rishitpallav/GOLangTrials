package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAgeUsingName(name string) (string, error) {

	name = strings.TrimSuffix(name, "\r\n")
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Name cannot be empty")
		return "", errors.New(fmt.Sprintln("The name entered is empty"))
	}

	apiURL := "https://api.agify.io/?name=" + name
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
