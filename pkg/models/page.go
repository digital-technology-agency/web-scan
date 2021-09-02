package models

/*Page type of page*/
type Page struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Robots      string `json:"robots"`
}
