CREATE TABLE IF NOT EXISTS orders (
	id bigserial,
    id_user int8 NULL,
	name_product varchar(100) NOT NULL,
	total_order int8 NULL,
    CONSTRAINT orders_pkey PRIMARY KEY (id),
    CONSTRAINT fk_orders_user FOREIGN KEY (id_user) REFERENCES users(id)
);