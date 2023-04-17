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


-- name: GetUserByOAuthID :one
SELECT users.id,
       username,
       email,
       avatar_id,
       permission,
       email_confirmed,
       is_banned,
       created_at,
       account_id
FROM users
         INNER JOIN oauth_account
                    ON oauth_account.user_id = users.id

WHERE account_id = ?;

-- name: CreateOAuthAccount :exec
INSERT oauth_account (user_id, oauth_provider, account_id)
VALUES (?, ?, ?);

-- name: CreateUser :execlastid
INSERT users (username, email, password)
VALUES (?, ?, ?);

-- name: CreateUserSubscription :exec
INSERT user_subscriptions (user_id, api_key)
VALUES (?, ?);


-- name: AddEmailToNewsletter :exec
INSERT INTO newsletter_subscribed (email)
VALUES (?);


-- name: GetProfileAndChats :many
SELECT  username,
        email,
        avatar_id,
        permission,
        email_confirmed,
        is_banned,
        users.created_at,
        chat_bot.title,
        chat_bot.chatbot_id
FROM users LEFT JOIN
     chat_bot ON chat_bot.user_id=users.id
WHERE users.id=?;
