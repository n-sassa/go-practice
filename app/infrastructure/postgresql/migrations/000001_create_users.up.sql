CREATE TABLE users (
	id serial4 NOT NULL,
	"name" text NOT NULL,
	"password" text NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);

INSERT INTO users ("name", "password") VALUES('test_user_1', 'password');
INSERT INTO users ("name", "password") VALUES('test_user_2', 'password');
INSERT INTO users ("name", "password") VALUES('test_user_3', 'password');