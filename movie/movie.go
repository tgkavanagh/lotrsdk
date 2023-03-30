package movie

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
)

type Movie_API_Endpoint struct {
  bearer string
  movie_EP string
}

type Movie_Info struct {
  Id string `json:"_id"`
  Name string
  RuntimeInMinutes int
  BudgetInMillions float32
  BoxOfficeRevenueInMillions float32
  AcademyAwardNominations int
  AcademyAwardWins int
}

type Movie_Franchise_Listing struct {
  Docs []Movie_Info
  Total int
  Limit int
  Offset int
  page int
  pages int
}

/*******************************************************************************
 *
 ******************************************************************************/
func (mEP *Movie_API_Endpoint) SetBearerToken(token string) int {
  // Bearer token should not be an empty string
  if token == "" {
    return -1
  }

  // Set object bearer token even if it is already set
  mEP.bearer = "Bearer " + token
  return 0
}

/*******************************************************************************
 *
 ******************************************************************************/
func (mEP *Movie_API_Endpoint) SetMovieEP(url string) {
  mEP.movie_EP = url
}

/*******************************************************************************
 *
 ******************************************************************************/
func (mEP *Movie_API_Endpoint) GetMovieFranchiseDetailsRaw() []byte {
  log.Println("executing GET command")

  // Create a new request using http
  req, err := http.NewRequest("GET", mEP.movie_EP, nil)

  // add authorization header to the req
  req.Header.Add("Authorization", mEP.bearer)

  // Send req using http Client
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      log.Println("Error on response.\n[ERROR] -", err)
      return nil
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
      log.Println("Error while reading the response bytes:", err)
      return nil
  }

  return body
}

/*******************************************************************************
 *
 ******************************************************************************/
func (mfl *Movie_Franchise_Listing) UnmarshalMovieFranchiseDetails(data []byte) int {
  if len(data) > 0 {
    err := json.Unmarshal(data, &(mfl))
    if err != nil {
      log.Printf("Failed to unmarshall response: (%s)\n", err)
      return -1
    }
  }

  return 0
}
