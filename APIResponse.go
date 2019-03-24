package gomdb

type APIResponseRating struct {
	Source string
	Value  string
}

type APIResponse struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []APIResponseRating
	Metascore  string
	IMDbRating string
	IMDbVotes  string
	IMDbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}
