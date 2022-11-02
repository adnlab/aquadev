-- Create table based on specified ERD

create table if not exists customers (
	id serial not null,
	customer_name char(50) not null,
	primary key (id)
);

create table if not exists products (
	id serial not null,
	name char(50) not null,
	primary key (id)
);

create table if not exists orders (
	order_id serial not null,
	customer_id int not null,
	product_id int not null,
	order_date date default current_date,
	total float not null,
	primary key (order_id),
	foreign key (customer_id) references customers(id),
	foreign key (product_id) references products(id)
);

-- Insert value into the table

insert into public.customers (customer_name) values ('andi'), ('candi'), ('gandi'), ('randi');

insert into public.products (name) values ('bakwan'), ('tahu'), ('tempe'), ('combro');

insert into public.orders (customer_id, product_id, total) values (2, 3, 10), (1, 2, 12), (3, 1, 2), (3, 2, 24);

-- Update values from the table

update public.customers set customer_name = 'gusti' where id = 3;

update public.products set name = 'kroket' where id = 4;

update public.orders set customer_id = 2 where order_id = 4;

-- Delete value from the table

delete from public.customers where id = 4;

delete from public.products where id = 4;

delete from public.orders where order_id =1;

-- Join data from orders table to customers and products table 

SELECT orders.order_id, customers.customer_name, products.name, orders.order_date, orders.total
FROM orders
INNER JOIN customers ON orders.customer_id = customers.id
INNER JOIN products ON orders.product_id = products.id;


create table if not exists Users (
	ID serial not null,
	Name char(50) not null,
	Email char(50) not null,
	Phone char(50) not null,
	Address char(50) not null,
	primary key (ID)
);