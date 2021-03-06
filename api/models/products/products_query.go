package products

// CRUD Operations
const (
	GET = `
		SELECT *
		FROM Products
		WHERE Id = ?
	`

	GETALL = `
		SELECT *
		FROM Products p
		WHERE p.IsDeleted = 0
		ORDER BY p.Code;
	`

	CREATE = `
		INSERT INTO
			Products (code, name, isactive, isdeleted, createdate, lastupdate)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	UPDATE = `
		UPDATE Products
		SET
			Code = ?,
			Name = ?,
			IsActive = ?,
			IsDeleted = ?,
			LastUpdate = ?
		WHERE Id = ?
	`

	DELETE = `
		UPDATE Products
		SET
			IsDeleted = 1,
			IsActive = 0,
			LastUpdate = ?
		WHERE Id = ?
	`
)

const (
	GET_SKUS_BY_PRODUCTID = `
		SELECT *
		FROM Skus s
		WHERE s.IsDeleted = 0 and s.ProductId = ?;
	`
)
