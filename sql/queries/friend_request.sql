-- name: CreateFriendRequest :exec
INSERT INTO friend_requests (sender_id, receiver_id)
VALUES ($1, $2);

-- name: GetFriendRequests :many
SELECT u.firstname, u.lastname, u.username, f.status, f.Id
FROM friend_requests f
JOIN users u ON f.receiver_id = u.id
WHERE f.receiver_id = $1;

-- name: AcceptFriendRequest :exec
UPDATE friend_requests
SET status = 'accepted'
WHERE sender_id = $1 AND receiver_id = $2;