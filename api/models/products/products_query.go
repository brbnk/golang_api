package products

const (
	GET = `
		SELECT
			p.Id, p.Code, p.Name, p.IsActive, p.IsDeleted,
			p.CreateDate, p.LastUpdate
		FROM Products p
		WHERE Id = ?
	`

	GETALL = `
		SELECT
			p.Id, p.Code, p.Name, p.IsActive,
			p.IsDeleted, p.CreateDate, p.LastUpdate
		FROM Products p
		WHERE IsDeleted = 0
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
