package repository_test

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
	"go-micro/tools"
	"io"
	"testing"
)

func TestArticle(t *testing.T) {
	tools.GetLogger().SetOutput(io.Discard)
	database := (&database.Database{}).InitializeSqlite().Migrate()

	repo := repository.Article{DataBase: database}
	dto, err := repo.Add(&entity.Article{Name: "test"})
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
	if !sel.DeletedAt.Valid {
		t.Fatalf("Not marked as deleted")
	}
	sel, _ = repo.Get(dto.ID, []string{})
	if sel != nil {
		t.Fatalf("Selecting deleted values")
	}

	lista, _ := repo.GetAll([]string{})
	if len(lista) != 0 {
		t.Fatalf("Too many values")
	}

}
func TestInsertWrongArticle(t *testing.T) {
	database := (&database.Database{}).InitializeSqlite().Migrate()

	repo := repository.Article{DataBase: database}
	dto, err := repo.Add(&entity.Article{})
	if dto != nil {
		t.Fatalf("Not Null return")
	}
	if err == nil {
		t.Fatalf("Nil err")
	}
}
