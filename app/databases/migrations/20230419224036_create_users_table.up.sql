CREATE TABLE users (
	id bigserial,
	name varchar(256) NOT NULL,
	email varchar (200) NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

-- ALTER TABLE users 
-- ADD email varchar(200) NULL;