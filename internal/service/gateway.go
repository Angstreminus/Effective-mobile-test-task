package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Angstreminus/Effective-mobile-test-task/internal/dto"
)

func GetAge(name, url string) (int, error) {
	var (
		sb   strings.Builder
		data dto.AgeRequest
	)
	sb.WriteString(url)
	sb.WriteString(name)
	resUrl := sb.String()

	resp, err := http.Get(resUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}
	return data.Age, nil
}

func GetGender(name, url string) (string, error) {
	var (
		sb   strings.Builder
		data dto.GenderRequest
	)
	sb.WriteString(url)
	sb.WriteString(name)
	resUrl := sb.String()

	resp, err := http.Get(resUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	return data.Gender, nil
}

func GetNationality(name, url string) (string, error) {
	var (
		sb   strings.Builder
		data dto.NationalityRequest
	)
	sb.WriteString(url)
	sb.WriteString(name)
	resUrl := sb.String()

	resp, err := http.Get(resUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	return data.Country[0].CountryID, nil
}
