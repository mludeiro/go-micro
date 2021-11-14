package repository

func GetArticles() []Article {
	articles := []Article{}
	getDB().Debug().Find(&articles)
	return articles
}

func AddArticle(a *Article) {
	getDB().Create(a)
}
