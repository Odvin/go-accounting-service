-- name: CreateClientProfile :one
INSERT INTO client.profile (
  id,
  adm,
  kyc,
  name,
  surname,
  password,
  email
) VALUES (
  $1, $2, $3, $4, $5, $6, $7 
) RETURNING id;

-- name: GetClientProfile :one
SELECT * FROM client.profile
WHERE id = $1 LIMIT 1;


-- name: DeleteClientProfile :exec
DELETE FROM client.profile
WHERE id = $1;

-- name: GetClientPasswordByEmail :one
SELECT password, password_updated FROM client.profile
WHERE email = $1 LIMIT 1;

-- name: GetClientPasswordById :one
SELECT password, password_updated FROM client.profile
WHERE id = $1 LIMIT 1;

-- name: UpdateClientProfile :exec
UPDATE client.profile
SET
  name = COALESCE(sqlc.narg(name), name),
  surname = COALESCE(sqlc.narg(surname), surname),
  email = COALESCE(sqlc.narg(email), email),
  updated = now()
WHERE
  id = $1;

-- name: UpdateClientPassword :exec
UPDATE client.profile
SET
  password = $2,
  password_updated = now()
WHERE
  id = $1
RETURNING *;

-- name: UpdateClientStatus :exec
UPDATE client.profile
SET
  adm = COALESCE(sqlc.narg(adm), adm),
  kyc = COALESCE(sqlc.narg(kyc), kyc),
  updated = now()
WHERE
  id = $1;
