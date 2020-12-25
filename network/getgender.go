package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jainSamkit/golang-design-patterns/enums"
)

type Gender struct {
	LikelyGender string `json:"likelyGender"`
}

func GetGender(firstName, lastName string) (string, error) {
	//break name into first and last name

	apiEndpoint := fmt.Sprintf("%v/%v/%v", enums.GENDER_API_ENDPOINT, firstName, lastName)
	net := Network{}
	net.SetHeader("X-API-KEY", enums.X_API_KEY_GENDER_API)

	res, err := net.Get(apiEndpoint)
	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(res.Body)
	var gender Gender
	_ = json.Unmarshal(body, &gender)

	if gender.LikelyGender == "male" {
		return enums.MALE, nil
	}
	return enums.FEMALE, nil
}
