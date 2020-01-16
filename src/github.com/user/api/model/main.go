package model

// Article model
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles :: let's declare a global Article array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

// LoadArticles will load all articles
func LoadArticles() []Article {
	articles := []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	return articles
}
