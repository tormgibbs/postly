-- name: CreateUser :one
INSERT INTO users (whatsapp_id)
VALUES ($1)
ON CONFLICT (whatsapp_id) DO UPDATE SET updated_at = now()
RETURNING *;

-- name: GetUserByWhatsAppID :one
SELECT * FROM users WHERE whatsapp_id = $1 LIMIT 1;

-- name: SetUserLoggedIn :exec
UPDATE users
SET is_logged_in = true, updated_at = now()
WHERE whatsapp_id = $1;
