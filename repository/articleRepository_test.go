package repository_test

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
	"testing"
)

func TestAdd(t *testing.T) {
	database.Initialize(true)
	database.Migrate()

	repo := repository.ArticleRepository{DB: database.GetDB()}
	article, err := repo.AddArticle(&entity.Article{Name: "test"})
	if article == nil || err != nil {
		t.Fatalf("Null return")
	}
	articlesel := repo.GetArticle(article.ID, []string{})

	if articlesel == nil {
		t.Fatalf("Not created")
	}

	articlesel = repo.DeleteArticle(article.ID)

	if articlesel == nil {
		t.Fatalf("Delete not working")
	}
	if !articlesel.DeletedAt.Valid {
		t.Fatalf("Not marked as deleted")
	}
	articlesel = repo.GetArticle(article.ID, []string{})
	if articlesel != nil {
		t.Fatalf("Selecting deleted values")
	}
}
