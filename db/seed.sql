DROP TABLE IF EXISTS "users";

CREATE TABLE IF NOT EXISTS "users" (
	"user_id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"first_name"	TEXT,
	"last_name"	TEXT,
	"email"	TEXT UNIQUE
);

INSERT INTO "users"
("first_name", "last_name", "email") VALUES
("Jane", "Doe", "jane.doe@example.com"),
("Spencer", "White", "spencer.white@example.com"),
("Edward", "Riverstone", "ed.riverstone@example.com");
