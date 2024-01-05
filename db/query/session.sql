-- name: CreateSession :one
INSERT INTO client.session (
  id,
  sub,
  refresh,
  agent,
  ip,
  blocked,
  expires
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
SELECT * FROM client.session
WHERE id = $1 LIMIT 1;