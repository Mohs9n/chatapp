-- +goose Up
CREATE TABLE friend_requests (
    Id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL REFERENCES users(Id),
    receiver_id INTEGER NOT NULL REFERENCES users(Id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending'
);

-- +goose Down
DROP TABLE friend_requests;