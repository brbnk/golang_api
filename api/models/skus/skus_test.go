package skus

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

type Test struct {
	r    SkuRepositoryInterface
	id   uint
	mock func()
}

func TestGetAllSkus(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := NewSkuRepository(sqlxDB)
	time := time.Now()

	test := &Test{
		r: repository,
		mock: func() {
			rows := sqlmock.NewRows([]string{"id", "code", "name", "productid", "isactive", "isdeleted", "createdate", "lastupdate"}).
				AddRow(1, "1234", "SKU Teste 1", 1, true, false, time, time).
				AddRow(2, "5678", "SKU Teste 2", 1, true, false, time, time).
				AddRow(3, "1324", "SKU Teste 3", 1, true, false, time, time)

			mock.ExpectQuery(GETALL).WillReturnRows(rows)
		},
	}

	expect := []Sku{
		{
			Id:         1,
			Code:       "1234",
			ProductId:  1,
			Name:       "SKU Teste 1",
			IsActive:   true,
			IsDeleted:  false,
			CreateDate: time,
			LastUpdate: time,
		},
		{
			Id:         2,
			Code:       "5678",
			ProductId:  1,
			Name:       "SKU Teste 2",
			IsActive:   true,
			IsDeleted:  false,
			CreateDate: time,
			LastUpdate: time,
		},
		{
			Id:         3,
			Code:       "1324",
			ProductId:  1,
			Name:       "SKU Teste 3",
			IsActive:   true,
			IsDeleted:  false,
			CreateDate: time,
			LastUpdate: time,
		},
	}

	test.mock()

	got, err := test.r.GetSkus()
	if err != nil {
		t.Errorf("GetAllProducts() error new = %v", err)
		return
	}

	for index, item := range got {
		if !reflect.DeepEqual(item, &expect[index]) {
			t.Errorf("GetAllProducts() = %v, expect %v", item, &expect[index])
		}
	}
}

func TestGetProductById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := NewSkuRepository(sqlxDB)
	time := time.Now()

	test := &Test{
		r:  repository,
		id: 2,
		mock: func() {
			rows := sqlmock.NewRows([]string{"id", "code", "name", "productid", "isactive", "isdeleted", "createdate", "lastupdate"}).
				AddRow(1, "1234", "SKU Teste 1", 1, true, false, time, time).
				AddRow(2, "4567", "SKU Teste 2", 1, true, false, time, time)

			mock.ExpectQuery(GET).WithArgs(2).WillReturnRows(rows)
		},
	}

	expect := &Sku{
		Id:         1,
		Code:       "1234",
		Name:       "SKU Teste 1",
		ProductId:  1,
		IsActive:   true,
		IsDeleted:  false,
		CreateDate: time,
		LastUpdate: time,
	}

	test.mock()

	got, err := test.r.GetSkuById(&Sku{Id: test.id})

	if err != nil {
		t.Errorf("GetProductById() error new = %v", err)
		return
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("GetProductById() = %v, expect %v", err, expect)
	}

}

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := NewSkuRepository(sqlxDB)
	time := time.Now()

	test := &Test{
		r: repository,
		mock: func() {
			mock.ExpectPrepare(CREATE).ExpectExec().
				WithArgs("1234", "SKU Teste 1", 1, true, false, time, time).
				WillReturnResult(sqlmock.NewResult(1, 1))
		},
	}

	payload := &Sku{
		Code:       "1234",
		Name:       "SKU Teste 1",
		ProductId:  1,
		IsActive:   true,
		IsDeleted:  false,
		CreateDate: time,
		LastUpdate: time,
	}

	test.mock()

	err = test.r.CreateSku(payload)
	if err != nil {
		t.Errorf("CreateProduct() ERROR >> %v", err)
		return
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("CreateProduct() = %v, expect %v", err, nil)
	}
}

func TestUpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := NewSkuRepository(sqlxDB)
	time := time.Now()

	test := &Test{
		r: repository,
		mock: func() {
			mock.ExpectPrepare(UPDATE).ExpectExec().
				WithArgs("1234", "SKU EDITADO 1", true, false, time, 1).
				WillReturnResult(sqlmock.NewResult(0, 1))
		},
	}

	test.mock()

	payload := &Sku{
		Id:         1,
		Code:       "1234",
		Name:       "SKU EDITADO 1",
		IsActive:   true,
		IsDeleted:  false,
		CreateDate: time,
		LastUpdate: time,
	}

	err = test.r.UpdateSku(payload)
	if err != nil {
		t.Errorf("UpdateProduct() ERROR >> %v", err)
		return
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("UpdateProduct() = %v, expect %v", err, nil)
	}
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repository := NewSkuRepository(sqlxDB)

	time := time.Now()

	test := &Test{
		r: repository,
		mock: func() {
			mock.ExpectPrepare(DELETE).ExpectExec().
				WithArgs(time, 1).
				WillReturnResult(sqlmock.NewResult(0, 1))
		},
	}

	test.mock()

	err = test.r.DeleteSku(&Sku{Id: 1, LastUpdate: time})
	if err != nil {
		t.Errorf("DeleteProduct() ERROR >> %v", err)
		return
	}

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("DeleteProduct() = %v, expect %v", err, nil)
	}
}
