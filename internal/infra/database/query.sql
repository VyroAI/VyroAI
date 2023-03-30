-- name: GetUserByID :one
SELECT id,
       username,
       email,
       avatar_id,
       permission,
       email_confirmed,
       is_banned,
       created_at
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
       is_banned,
       created_at
FROM users
WHERE email = ?;

-- name: GetUserByUsername :one
SELECT id,
       username,
       email,
       avatar_id,
       permission,
       email_confirmed,
       is_banned,
       created_at
FROM users
WHERE username = ?;

-- name: CreateUser :execlastid
INSERT users (username, email, password, subscription_id)
VALUES (?, ?, ?, ?);

-- name: CreateUserSubscription :execlastid
INSERT user_subscriptions (api_key)
VALUES (?);


-- name: AddEmailToNewsletter :exec
INSERT INTO newsletter_subscribed (email)
VALUES (?);