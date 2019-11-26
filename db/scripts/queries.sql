-- name: drop-users-table
DROP TABLE IF EXISTS "users";

-- name: create-users-table
CREATE TABLE IF NOT EXISTS users (
	user_id	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	first_name	TEXT,
	last_name	TEXT,
	email	TEXT UNIQUE
);

-- name: seed-users-table
INSERT INTO users
(first_name, last_name, email) VALUES
("Jane", "Doe", "jane.doe@example.com"),
("Spencer", "White", "spencer.white@example.com"),
("Edward", "Riverstone", "ed.riverstone@example.com");

-- name: get-user-by-id
SELECT * FROM users WHERE user_id = ?;

-- name: get-users
SELECT * FROM users ORDER BY user_id;

-- name: create-user
INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?);
