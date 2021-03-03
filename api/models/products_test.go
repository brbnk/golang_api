package models

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type Test struct {
	r       ProductRepositoryInterface
	id      uint
	mock    func()
	expect  *Product
	wantErr bool
}

func TestGetAllProducts(t *testing.T) {
	t.Error()
}

func TestGetProductById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	repository := NewProductRepository(db)
	time := time.Now()

	test := &Test{
		r:  repository,
		id: 2,
		mock: func() {
			rows := sqlmock.NewRows([]string{"id", "code", "name", "isactive", "isdeleted", "createdaet", "lastupdate"}).
				AddRow(1, "1234", "Produto Teste 1", true, false, time, time).
				AddRow(2, "4567", "Produto Teste 2", true, false, time, time)

			mock.ExpectQuery(`
				SELECT p.Id, p.Code, p.IsActive, p.IsDeleted, p.Name, p.CreateDate, p.LastUpdate
				FROM Products p
				WHERE Id = ?`).WithArgs(2).WillReturnRows(rows)
		},
		expect: &Product{
			Id:         1,
			Code:       "1234",
			Name:       "Produto Teste 1",
			IsActive:   true,
			IsDeleted:  false,
			CreateDate: time,
			LastUpdate: time,
		},
	}

	test.mock()

	got, err := test.r.GetProductById(&Product{Id: test.id})

	if (err != nil) != test.wantErr {
		t.Errorf("GetProductById() error new = %v, wantErr %v", err, test.wantErr)
		return
	}

	if !reflect.DeepEqual(got, test.expect) {
		t.Errorf("GetProductById() = %v, expect %v", err, test.expect)
	}

}

func TestCreateProduct(t *testing.T) {
	t.Error()
}

func TestUpdateProduct(t *testing.T) {
	t.Error()
}

func TestDeleteProduct(t *testing.T) {
	t.Error()
}
