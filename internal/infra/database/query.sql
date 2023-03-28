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

-- name: GetUserByUsername :one
SELECT id,
       username,
       email,
       avatar_id,
       permission,
       email_confirmed,
       status
FROM users
WHERE username = ?;

-- name: CreateUser :execlastid
INSERT users (username, email, password)
VALUES (?, ?, ?);


-- name: AddEmailToNewsletter :exec
INSERT INTO newsletter_subscribed (email)
VALUES (?);