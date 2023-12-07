-- name: CreateClientProfile :one
INSERT INTO client.profile (
  id,
  adm,
  kyc,
  name,
  surname
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id;

-- name: GetClientProfile :one
SELECT * FROM client.profile
WHERE id = $1 LIMIT 1;


-- name: DeleteCreateClientProfile :exec
DELETE FROM client.profile
WHERE id = $1;