-- +goose Up
CREATE TABLE chats (
	id SERIAL 	PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now() 	
);

CREATE TABLE messages (
	id SERIAL PRIMARY KEY,
	chat_id INT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
	text VARCHAR(5000) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now()
);
-- +goose Down
DROP TABLE messages;
DROP TABLE chats;