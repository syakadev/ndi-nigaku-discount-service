package auth_db

func GetPhoneNumber() string {
	return `
	SELECT
		phone_number,
		otp,
		expired_at
	FROM ndi_otp
	WHERE phone_number = $1
	LIMIT 1
	`
}

func InsertOTP() string {
	return `
	INSERT INTO ndi_otp (phone_number, otp, expired_at, created_at, updated_at)
	VALUES ($1, $2, $3, NOW(), NOW())
	`
}

func UpdateOTP() string {
	return `
	UPDATE ndi_otp
	SET otp = $2,
		expired_at = $3,
		updated_at = NOW()
	WHERE phone_number = $1
	`
}

func DeleteOTPByPhoneNumber() string {
	return `
	DELETE FROM ndi_otp
	WHERE phone_number = $1
	`
}
