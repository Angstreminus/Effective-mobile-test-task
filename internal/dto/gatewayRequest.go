package dto

type AgeRequest struct {
	Count int    `json:"count"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
}

type GenderRequest struct {
	Count       int    `json:"count"`
	Gender      string `json:"gender"`
	Name        string `json:"name"`
	Probability int    `json:"probability"`
}

type NationalityRequest struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}
