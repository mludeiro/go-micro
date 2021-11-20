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
	database.CreateSampleData()

	repo := repository.ArticleRepository{DB: database.GetDB()}
	article, err := repo.AddArticle(&entity.Article{Name: "Prueba"})
	if article == nil || err != nil {
		t.Fatalf("Que paso")
	}
	articlesel := repo.GetArticle(article.ID, []string{})

	if articlesel == nil {
		t.Fatalf("No se creo")
	}

	articlesel = repo.DeleteArticle(article.ID)

	if articlesel == nil {
		t.Fatalf("No lo borro")
	}
	if !articlesel.DeletedAt.Valid {
		t.Fatalf("No lo borro")
	}
	articlesel = repo.GetArticle(article.ID, []string{})
	if articlesel != nil {
		t.Fatalf("No lo borro")
	}
}
