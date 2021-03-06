package skus

const (
	GET = `
		SELECT *
		FROM Skus
		WHERE Id = ?
	`

	GETALL = `
		SELECT *
		FROM Skus s
		WHERE s.IsDeleted = 0
		ORDER BY s.Code;
	`

	CREATE = `
		INSERT INTO
			Skus (code, name, productid, isactive, isdeleted, createdate, lastupdate)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	UPDATE = `
		UPDATE Skus
		SET
			Code = ?,
			Name = ?,
			IsActive = ?,
			IsDeleted = ?,
			LastUpdate = ?
		WHERE Id = ?
	`

	DELETE = `
		UPDATE Skus
		SET
			IsDeleted = 1,
			IsActive = 0,
			LastUpdate = ?
		WHERE Id = ?
	`
)
