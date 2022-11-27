package utils

import (
	"encoding/json"
	"net/http"
)

func GetCall(url string)(interface{}, error)  {
	client := http.Client{}
	req , err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var respData interface{}
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"X-API-KEY": {"6317ca665f618f804f5f58bf"},
	}
	
	resp , err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&respData)

		if err != nil {
			return nil, err
		}
	}

	return respData, nil
}