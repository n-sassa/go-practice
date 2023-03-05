CREATE TABLE todos (
	id serial4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	title text NOT NULL DEFAULT ''::text,
	user_id int4 NOT NULL,
	CONSTRAINT todos_pk PRIMARY KEY (id),
	CONSTRAINT todos_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO todos (title, user_id) VALUES('test', 1);
INSERT INTO todos (title, user_id) VALUES('test', 2);
INSERT INTO todos (title, user_id) VALUES('test', 3);
