package repository

func GetArticle(id uint) *Article {
	articles := []Article{}
	getDB().Debug().Find(&articles, id)
	if len(articles) == 1 {
		return &articles[0]
	} else {
		return nil
	}
}

func GetArticles() []Article {
	articles := []Article{}
	getDB().Debug().Find(&articles)
	return articles
}

func AddArticle(a *Article) {
	getDB().Create(a)
}
