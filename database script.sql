CREATE TABLE public.tenant (
	tenant_id serial4 NOT NULL,
	tenant_name varchar(100) NOT NULL,
	tenant_details varchar(1024) NOT NULL,
	created_on timestamp NOT NULL,
	CONSTRAINT tenant_pkey PRIMARY KEY (tenant_id),
	CONSTRAINT tenant_tenantname_key UNIQUE (tenant_name)
);


CREATE TABLE public.users (
	username varchar(50) NOT NULL,
	"password" varchar(50) NOT NULL,
	tenant_id int4 NOT NULL,
	created_on timestamp NOT NULL,
	CONSTRAINT users_pkey1 PRIMARY KEY (username),
	CONSTRAINT fk_tenant
	FOREIGN KEY(tenant_id) 
	  REFERENCES tenant(tenant_id)
);

CREATE TABLE public.product (
	product_id serial4 NOT NULL,
	"name" varchar(100) NOT NULL,
	tenant_id int4 NOT NULL,
	price numeric NOT NULL,
	CONSTRAINT product_pkey1 PRIMARY KEY (product_id),
	CONSTRAINT fk_tenant
	FOREIGN KEY(tenant_id) 
	  REFERENCES tenant(tenant_id)
);

CREATE TABLE public."order" (
	order_id serial4 NOT NULL,
	user_name varchar(50) NOT NULL,
	datetime timestamp NOT NULL,
	CONSTRAINT order_pkey PRIMARY KEY (order_id),
	CONSTRAINT fk_user
      FOREIGN KEY(user_name) 
	  REFERENCES users(username)
);

CREATE TABLE public.order_product (
	id serial4 NOT NULL,
	order_id int4 NOT NULL,
	product_id int4 NOT NULL,
	product_quantity int4 NOT NULL,
	CONSTRAINT order_product_pkey1 PRIMARY KEY (id),
	CONSTRAINT fk_order
      FOREIGN KEY(order_id) 
	  REFERENCES "order"(order_id),
	CONSTRAINT fk_product
      FOREIGN KEY(product_id) 
	  REFERENCES product(product_id)
);

--Insert data into tenant table
INSERT INTO public.tenant (tenant_name, tenant_details, created_on)
VALUES('Amazon', 'Shop on the Amazon App. Fast, convenient and secure.Over 17 crore products in your pocket', current_timestamp),
('Flipkart', 'Flipkart: The One-stop Shopping Destination. E-commerce is revolutionizing the way we all shop in India.', current_timestamp);

--Insert data into users table
INSERT INTO public."users"  
(username, "password", tenant_id, created_on)
VALUES
  ('kmarke','abc123','1', Current_TimeStamp),
  ('asinha','abc123','2', Current_TimeStamp),
  ('mkotian','abc123','2', Current_TimeStamp);

--Insert data into product table 
INSERT INTO public.product 
("name", tenant_id, price) values
('PeterEnglandShirt', 1, 100),
('X TShirt', 1, 20),
('Woodland Jackets', 2, 56),
('T-Shirt', 1, 48),
('PeterEnglandShirt', 2, 100);

--Insert data into Order table
INSERT INTO public."order"
(user_name, datetime)
VALUES('kmarke', Current_TimeStamp),
('asinha', Current_TimeStamp),
('kmarke', Current_TimeStamp),
('asinha', Current_TimeStamp);

--Insert data into Order_product table
INSERT INTO public.order_product
(order_id, product_id, product_quantity)
VALUES(1, 1, 2),
(3, 2, 2),
(2, 3, 3),
(4, 4, 1);