package repository_test

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
	"testing"
)

func TestArticleType(t *testing.T) {
	database.Initialize(true)
	database.Migrate()

	repo := repository.ArticleTypeRepository{DB: database.GetDB()}
	dto, err := repo.Add(&entity.ArticleType{Name: "test"})
	if dto == nil || err != nil {
		t.Fatalf("Null return")
	}
	sel := repo.Get(dto.ID, []string{})

	if sel == nil {
		t.Fatalf("Not created")
	}

	sel = repo.Delete(dto.ID)

	if sel == nil {
		t.Fatalf("Delete not working")
	}
	sel = repo.Get(dto.ID, []string{})
	if sel != nil {
		t.Fatalf("Selecting deleted values")
	}

	if len(repo.GetAll([]string{})) != 0 {
		t.Fatalf("Too many values")
	}

}
