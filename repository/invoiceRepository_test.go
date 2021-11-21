package repository_test

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
	"testing"
)

func TestInvoice(t *testing.T) {
	database.InitializeSqlite()
	database.Migrate()

	repo := repository.Invoice{DB: database.GetDB()}
	dto, err := repo.Add(&entity.Invoice{})
	if dto == nil || err != nil {
		t.Fatalf("Null return")
	}
	sel := repo.Get(dto.ID, []string{})

	if sel == nil {
		t.Fatalf("Not created")
	}

	sel = repo.Get(dto.ID, []string{})
	if sel == nil {
		t.Fatalf("Selecting deleted values")
	}

	if len(repo.GetAll([]string{})) != 1 {
		t.Fatalf("No values")
	}

}
