package repository

func GetArticleType(id uint, fetchs []string) *ArticleType {
	articleType := ArticleType{}

	db := getDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	rows := db.Find(&articleType, id).RowsAffected
	if rows == 1 {
		return &articleType
	} else {
		return nil
	}
}

func GetArticleTypes(fetchs []string) []ArticleType {
	articleTypes := []ArticleType{}

	db := getDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articleTypes)
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
