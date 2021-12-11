package repository_test

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
	"go-micro/tools"
	"io/ioutil"
	"testing"
)

func TestInvoice(t *testing.T) {
	tools.GetLogger().SetOutput(ioutil.Discard)
	database := (&database.Database{}).InitializeSqlite().Migrate()

	repo := repository.Invoice{DataBase: database}
	dto, err := repo.Add(&entity.Invoice{})
	if dto == nil || err != nil {
		t.Fatalf("Null return")
	}
	sel, err := repo.Get(dto.ID, []string{})

	if err != nil {
		t.Fatalf("Error")
	}

	if sel == nil {
		t.Fatalf("Not created")
	}

	sel, err = repo.Get(dto.ID, []string{})
	if err != nil {
		t.Fatalf("Error")
	}
	if sel == nil {
		t.Fatalf("Selecting deleted values")
	}

	all, err := repo.GetAll([]string{})
	if err != nil {
		t.Fatalf("Error")
	}
	if len(all) != 1 {
		t.Fatalf("No values")
	}

}
