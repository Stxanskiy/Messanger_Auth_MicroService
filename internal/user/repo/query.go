package repo

const (
	QuerySaveToken           = `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3);`
	QueryDeleteToken         = `DELETE FROM refresh_tokens WHERE token = $1;`
	QueryIsTokenValid        = `SELECT COUNT(*) > 0 FROM refresh_tokens WHERE token = $1 AND expires_at > NOW();`
	QueryDeleteTokenByUserID = `Delete from refresh_tokens WHERE user_id = $1;`

	QueryCreateUser           = `INSERT INTO users (nickname, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id;`
	QueryIsNicknameTaken      = `SELECT 1 FROM users WHERE nickname = $1 LIMIT 1;`
	QueryGetUserByNickname    = `SELECT id, nickname, password_hash FROM users WHERE nickname = $1;`
	QuerySearchUserByNickname = `SELECT id,  nickname, created_at  from users where nickname ILIKE $1`
)
