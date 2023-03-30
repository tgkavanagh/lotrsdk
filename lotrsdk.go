package lotrsdk

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  . "../lotrsdk/movie"
)

const (
  LOTR_MOVIE_ENDPOINT_URL = "https://the-one-api.dev/v2/movie"
)

type LOTR_Data struct {
  movieEP Movie_API_Endpoint
  response []byte
  movies Movie_Franchise_Listing
}

var lotr_movid_ids map[string]string

/*******************************************************************************
 *
 ******************************************************************************/
func CreateLOTRData(token string) LOTR_Data {
  object := new(LOTR_Data)
  object.movieEP.SetBearerToken(token)
  object.movieEP.SetMovieEP = LOTR_MOVIE_ENDPOINT_URL

  return object
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o *LOTR_Data) GetLOTRMovieData() int {
  o.response = o.movieEP.GetMovieFranchiseDetailsRaw()

  return o.movies.UnmarshalMovieFranchiseDetails(o.response)
}

/*******************************************************************************
 *
 ******************************************************************************/
func (o *LOTR_Data) GetMovieId() map[string]string {
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
func (o LOTR_Data) GetMovieRunTimeInMinutes() map[string]int {
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
func (o LOTR_Data) GetMovieBudgetInMillions() map[string]float32 {
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
func (o LOTR_Data) GetMovieRevenueInMillions() map[string]float32 {
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
func (o LOTR_Data) GetMovieAcademyAwardNominations() map[string]int {
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
func (o LOTR_Data) GetMovieAcademyAwardWins() map[string]int {
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
func (o LOTR_Data) GetMovieRottenTomatoesScore() map[string]float32 {
  if o.Movies.Total != 0 {
    results := make(map[string]float32)

    for _, movie := range o.Movies.Docs {
      results[movie.Name] = movie.RottenTomatoesScore
    }

    return results
  }

  return nil
}
