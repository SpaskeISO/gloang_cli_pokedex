package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cashe
	data, ok := c.cashe.Get(fullURL)
	if ok {
		// cahse found
		//fmt.Println("cashe found!")
		locationAreasResponse := LocationAreaResponse{}
		err := json.Unmarshal(data, &locationAreasResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationAreasResponse, nil
	}
	//fmt.Println("cashe not found")
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code %v", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreasResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cashe.Add(fullURL, data)

	return locationAreasResponse, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cashe
	data, ok := c.cashe.Get(fullURL)
	if ok {
		// cahse found
		//fmt.Println("cashe found!")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	//fmt.Println("cashe not found")
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationArea{}, err
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %v", response.StatusCode)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cashe.Add(fullURL, data)

	return locationArea, nil
}
