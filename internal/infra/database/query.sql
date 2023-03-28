-- name: GetUserByID :one
SELECT id,
       username,
       email,
       avatar_id,
       permission,
       email_confirmed,
       status
FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT id,
       username,
       email,
       avatar_id,
       password,
       permission,
       email_confirmed,
       status
FROM users
WHERE email = ?;


-- name: AddEmailToNewsletter :exec
INSERT INTO newsletter_subscribed (email)
VALUES (?);