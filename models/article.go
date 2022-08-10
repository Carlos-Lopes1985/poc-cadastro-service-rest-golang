package models

// Article
type Article struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

// Articles
var Articles []Article

func ReturnAllArticles() []Article {

	return []Article{
		Article{Title: "Python Intermediate and Advanced 101",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089KVK23P"},
		Article{Title: "R programming Advanced",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089WH12CR"},
		Article{Title: "R programming Fundamentals",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089S58WWG"},
	}
}
