CREATE TABLE users (
	id serial4 NOT NULL,
	"name" text NOT NULL,
	"password" text NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

INSERT INTO users (id, "name", "password") VALUES(1, 'test_user_1', 'password');
INSERT INTO users (id, "name", "password") VALUES(2, 'test_user_2', 'password');
INSERT INTO users (id, "name", "password") VALUES(3, 'test_user_3', 'password');