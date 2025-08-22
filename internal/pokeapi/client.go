package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

func GetLocations(url string)([]string,string,string,error) {
	res,err:=http.Get(url)
	if err!=nil{
		return []string{},"","",err
	}
	
	var raw LocationsResponse

	defer res.Body.Close()

	body,err:=io.ReadAll(res.Body)

	if err!=nil{
		return []string{},"","",err
	}

	if err:=json.Unmarshal(body,&raw);err!=nil{
		return []string{},"","",err
	}
	locations:=[]string{}

	for _,name := range raw.Results{
		locations=append(locations, name.Name)
	}

	var prevUrl string

	if raw.Previous != nil{
		prevUrl=*raw.Previous
	}

	return locations,prevUrl,raw.Next,nil
}
