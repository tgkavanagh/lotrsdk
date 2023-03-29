package lotrsdk

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
)

const (
  MOVIE_API_URL = "https://the-one-api.dev/v2/movie"
)

type LOTR_Movie struct {
  Id string `json:"_id"`
  Name string
  RuntimeInMinutes int
  BudgetInMillions float32
  BoxOfficeRevenueInMillions float32
  AcademyAwardNominations int
  AcademyAwardWins int
  RottenTomatoesScore float32
}

type LOTR_Movies struct {
  Docs []LOTR_Movie
  Total int
  Limit int
  Offset int
  page int
  pages int
}

type LOTR_Object struct {
  bearer string
  Response []byte
  Movies LOTR_Movies
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o *LOTR_Object) SetBearerToken(token string) int {
  // Bearer token should not be an empty string
  if token == "" {
    return -1
  }

  // Set object bearer token even if it is already set
  o.bearer = "Bearer " + token
  return 0
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o *LOTR_Object) GetMovieDetails() int {
  // Get details only if the object does not contain the info
  if len(o.Response) == 0 {
    log.Println("executing GET command")

    // Create a new request using http
    req, err := http.NewRequest("GET", MOVIE_API_URL, nil)

    // add authorization header to the req
    req.Header.Add("Authorization", o.bearer)

    // Send req using http Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
        return -1
    }

    defer resp.Body.Close()

    o.Response, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Error while reading the response bytes:", err)
        return -1
    }
  }

  return 0
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o *LOTR_Object) UnmarshalMovieDetails() int {
  if o.GetMovieDetails() == 0 {
    err := json.Unmarshal(o.Response, &(o.Movies))
    if err != nil {
      log.Printf("Failed to unmarshall response: (%s)\n", err)
      return -1
    }
  }

  return 0
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o LOTR_Object) GetMovieId() map[string]string {
  if o.Movies.Total != 0 {
    results := make(map[string]string)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.Id
    }

    return results
  }

  return nil
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o LOTR_Object) GetMovieRunTimeInMinutes() map[string]int {
  if o.Movies.Total != 0 {
    results := make(map[string]int)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.RuntimeInMinutes
    }

    return results
  }

  return nil
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o LOTR_Object) GetMovieBudgetInMillions() map[string]float32 {
  if o.Movies.Total != 0 {
    results := make(map[string]float32)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.BudgetInMillions
    }

    return results
  }

  return nil
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o LOTR_Object) GetMovieRevenueInMillions() map[string]float32 {
  if o.Movies.Total != 0 {
    results := make(map[string]float32)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.BoxOfficeRevenueInMillions
    }

    return results
  }

  return nil
}


/*******************************************************************************
 *
 ******************************************************************************/
func (o LOTR_Object) GetMovieAcademyAwardNominations() map[string]int {
  if o.Movies.Total != 0 {
    results := make(map[string]int)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.AcademyAwardNominations
    }

    return results
  }

  return nil
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o LOTR_Object) GetMovieAcademyAwardWins() map[string]int {
  if o.Movies.Total != 0 {
    results := make(map[string]int)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.AcademyAwardWins
    }

    return results
  }

  return nil
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o LOTR_Object) GetMovieRottenTomatoesScore() map[string]float32 {
  if o.Movies.Total != 0 {
    results := make(map[string]float32)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.RottenTomatoesScore
    }

    return results
  }

  return nil
}
