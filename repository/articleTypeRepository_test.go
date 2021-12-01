package repository_test

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
	"go-micro/tools"
	"io"
	"testing"
)

func TestArticleType(t *testing.T) {
	tools.GetLogger().SetOutput(io.Discard)

	database := (&database.Database{}).InitializeSqlite().Migrate()

	repo := repository.ArticleType{DataBase: database}
	dto, err := repo.Add(&entity.ArticleType{Name: "test"})
	if dto == nil || err != nil {
		t.Fatalf("Null return")
	}
	sel, _ := repo.Get(dto.ID, []string{})

	if sel == nil {
		t.Fatalf("Not created")
	}

	sel, _ = repo.Delete(dto.ID)

	if sel == nil {
		t.Fatalf("Delete not working")
	}
	sel, _ = repo.Get(dto.ID, []string{})
	if sel != nil {
		t.Fatalf("Selecting deleted values")
	}

	list, err := repo.GetAll([]string{})
	if err != nil {
		t.Fatalf("Too many values")
	}
	if len(list) != 0 {
		t.Fatalf("Too many values")
	}

}
