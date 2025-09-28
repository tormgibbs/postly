-- name: CreateUser :one
INSERT INTO users (whatsapp_id)
VALUES ($1)
ON CONFLICT (whatsapp_id) DO UPDATE SET updated_at = now()
RETURNING *;

-- name: GetUserByWhatsAppID :one
SELECT * FROM users WHERE whatsapp_id = $1 LIMIT 1;

-- name: SetUserLoggedInWithToken :exec
UPDATE users
SET 
  is_logged_in = true,
  access_token = $2,
  refresh_token = $3,
  token_expiry = $4,
  updated_at = now()
WHERE whatsapp_id = $1;

-- name: UpdateUserTokens :exec
UPDATE users
SET 
  access_token = $2,
  refresh_token = $3,
  token_expiry = $4,
  updated_at = now()
WHERE whatsapp_id = $1;

-- name: GetUserTokens :one
SELECT access_token, refresh_token, token_expiry 
FROM users 
WHERE whatsapp_id = $1;