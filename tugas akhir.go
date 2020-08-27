package main

import (
  "fmt" ;
  "net/http" ;
  "encoding/json" ;
  "bytes" ;
  "net/url"

)

var baseUrl = "http://localhost:8080"

type Rental struct {
	Id          string `json:"id"`
	Brand       string `json:"brand"`
	Year        int    `json:"year"`
	OwnerId     string `json:"owner_id"`
	RentPrice   int    `json:"rent_per_hour"`
	IsAvailable int    `json:"availability"`
}

type response struct {
  Success bool `json:"success"`
  Message string `json:"message"`
  Data []Rental `json:"data"`
}

func get()(response, error){
   var err error
   var client = &http.Client{}
   var data response

   request, err := http.NewRequest("GET", baseUrl+"/mobil", nil)
   if err != nil {
     return data, err
   }

   request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   response, err := client.Do(request)
   if err != nil {
     return data, err
   }
   defer response.Body.Close()

   err = json.NewDecoder(response.Body).Decode(&data)
   if err != nil {
     return data, err
   }

   return data, nil
}

func post(key string, input string)(response, error){
  var err error
  var client = &http.Client{}
  var data response

  var param =  url.Values{}
  param.Set(key, input)
  var payload = bytes.NewBufferString(param.Encode())
  request, err := http.NewRequest("POST", baseUrl+"/mobil", payload)
  if err != nil {
    return data, err
  }
  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  response, err := client.Do(request)
  if err != nil {
    return data, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err != nil {
    return data, err
  }

  return data, nil
}

func patch(id string, input string)(response, error){
  var err error
  var client = &http.Client{}
  var data response

  var param =  url.Values{}
  param.Set("Brand", input)
  var payload = bytes.NewBufferString(param.Encode())
  request, err := http.NewRequest("PATCH", baseUrl+"/mobil/"+id, payload)
  if err != nil {
    return data, err
  }

  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  response, err := client.Do(request)
  if err != nil {
    return data, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err != nil {
    return data, err
  }

  return data, nil
}

func delete(id string)(response, error){
  var err error
  var client = &http.Client{}
  var data response

  request, err := http.NewRequest("DELETE", baseUrl+"/mobil/"+id, nil)
  if err != nil {
    return data, err
  }

  request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  response, err := client.Do(request)
  if err != nil {
    return data, err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err != nil {
    return data, err
  }

  return data, nil
}

func main(){
  // var response, err = get()
  // var response, err = post("ownerId", "O11")
  // var response, err = patch("1", "Toyotaaa")
  var response, err = delete("4")
  if err != nil {
    fmt.Println("Error! ", err.Error())
    return
  }
  fmt.Println(response.Message)
  mobil := response.Data
  for _, each := range mobil {
    fmt.Print("ID : ",each.Id, ", Brand : ",each.Brand,", Tahun : ", each.Year, " Owner: ", each.OwnerId )
    switch each.IsAvailable {
    case 1 :
      fmt.Println(", Status: Occupied")
    default :
      fmt.Println(", Status: Available")
    }
  }
  // fmt.Println(menu)
  //   fmt.Println("ID :",menu.Id, ", Menu :", menu.Nama,", Harga :", menu.Harga )

}
