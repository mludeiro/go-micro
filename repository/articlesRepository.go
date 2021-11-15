package repository

func GetArticle(id uint) *Article {
	article := Article{}
	rows := getDB().Debug().Find(&article, id).RowsAffected
	if rows == 1 {
		return &article
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

func DeleteArticle(id uint) *Article {
	article := Article{}
	rows := getDB().Debug().Where("deleted_at is NULL").Delete(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}
