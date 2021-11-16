package repository

func GetArticleType(id uint) *ArticleType {
	articleType := ArticleType{}
	rows := getDB().Find(&articleType, id).RowsAffected
	if rows == 1 {
		return &articleType
	} else {
		return nil
	}
}

func GetArticleTypes() []ArticleType {
	articleTypes := []ArticleType{}
	getDB().Preload("Article").Find(&articleTypes)
	return articleTypes
}

func AddArticleType(a *ArticleType) {
	getDB().Create(a)
}

func DeleteArticleType(id uint) *ArticleType {
	articleType := ArticleType{}
	rows := getDB().Where("deleted_at is NULL").Delete(&articleType, id).RowsAffected
	if rows == 1 {
		return &articleType
	} else {
		return nil
	}
}
