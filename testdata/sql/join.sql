--
-- Test SELECT ... JOIN
--

DROP TABLE IF EXISTS customers;

CREATE TABLE customers (customer_id int primary key, customer varchar(128));

DROP TABLE IF EXISTS orders;

CREATE TABLE orders (order_id int primary key, customer_id int, price int, item varchar(128));

DROP TABLE IF EXISTS empty_customers;

CREATE TABLE empty_customers (customer_id int primary key, customer varchar(128));

DROP TABLE IF EXISTS empty_orders;

CREATE TABLE empty_orders (order_id int primary key, customer_id int, price int, item varchar(128));

INSERT INTO customers VALUES
    (1, 'Andrew'),
    (2, 'Barry'),
    (3, 'Cindy'),
    (4, 'David'),
    (5, 'Edward'),
    (6, 'Frank'),
    (7, 'Greg'),
    (8, 'Harry');

INSERT INTO orders VALUES
    (1, 1, 10, 'Groceries'),
    (2, 1, 1, 'Candy'),
    (3, 3, 2, 'Snacks'),
    (4, 6, 20, 'Gasoline'),
    (5, 7, 18, 'Gasoline'),
    (6, 9, 3, 'Taxes');

SELECT * FROM customers CROSS JOIN orders;

SELECT orders.order_id, customers.customer_id, customers.customer, orders.item
    FROM customers CROSS JOIN orders;

SELECT * FROM customers JOIN orders ON customers.customer_id = orders.customer_id;

SELECT orders.order_id, customers.customer_id, customers.customer, orders.item
    FROM customers JOIN orders ON customers.customer_id = orders.customer_id;

SELECT * FROM customers JOIN orders USING (customer_id);

SELECT orders.order_id, customers.customer_id, customers.customer, orders.item
    FROM customers JOIN orders USING (customer_id);

SELECT * FROM customers LEFT JOIN orders ON customers.customer_id = orders.customer_id;

SELECT orders.order_id, customers.customer_id, customers.customer, orders.item
    FROM customers LEFT JOIN orders ON customers.customer_id = orders.customer_id;

SELECT * FROM customers LEFT JOIN orders USING (customer_id);

SELECT orders.order_id, customers.customer_id, customers.customer, orders.item
    FROM customers LEFT JOIN orders USING (customer_id);

SELECT * FROM customers CROSS JOIN empty_orders;

SELECT * FROM customers JOIN empty_orders ON customers.customer_id = empty_orders.customer_id;

SELECT * FROM customers LEFT JOIN empty_orders ON customers.customer_id = empty_orders.customer_id;

SELECT * FROM empty_customers CROSS JOIN orders;

SELECT * FROM empty_customers JOIN orders ON empty_customers.customer_id = orders.customer_id;

SELECT * FROM empty_customers LEFT JOIN orders ON empty_customers.customer_id = orders.customer_id;

SELECT * FROM empty_customers CROSS JOIN empty_orders;

SELECT * FROM empty_customers JOIN empty_orders
    ON empty_customers.customer_id = empty_orders.customer_id;

SELECT * FROM empty_customers LEFT JOIN empty_orders
    ON empty_customers.customer_id = empty_orders.customer_id;
