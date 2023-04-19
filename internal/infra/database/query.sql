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
SELECT username,
       email,
       avatar_id,
       permission,
       email_confirmed,
       is_banned,
       users.created_at,
       chat_bot.title,
       chat_bot.id
FROM users
         LEFT JOIN
     chat_bot ON chat_bot.user_id = users.id
WHERE users.id = ?;


-- name: GetChatByChatID :one
SELECT *
from chat_bot
WHERE id = ?
  AND user_id = ?;

-- name: GetChatMessageByChatID :many
SELECT chat_message.id,
       chat_message.content,
       chat_message.bot,
       chat_message.created_by,
       chat_message.created_at,
       chat_bot.id as chatbot_id
from chat_message
         LEFT JOIN chat_bot on chat_bot.id = chat_message.chatbot_id
WHERE chatbot_id = ? && chat_bot.user_id = ?
LIMIT ? OFFSET ?;

-- name: CreateMessage :execlastid
INSERT INTO chat_message (content, created_by, chatbot_id)
VALUES (?, ?, ?);
