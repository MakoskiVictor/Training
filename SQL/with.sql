/* For this challenge you need to create a SELECT statement, this SELECT statement will use an 
IN to check whether a department has had a sale with a price over 90.00 dollars BUT the sql MUST use 
the WITH statement which will be used to select all columns from sales where the price is greater than 90.00, 
you must call this sub-query special_sales.

departments table schema
id
name

sales table schema
id
department_id (department foreign key) */
WITH special_sales AS (
    SELECT * 
    FROM sales
    WHERE sales.price > 90.00
)
SELECT 
    departments.*
FROM departments
WHERE departments.id IN (
    SELECT sales.department_id 
    FROM special_sales
)

-- MEJOR  QUERY 
SELECT DISTINCT * FROM departments
JOIN sales ON sales.department_id = departments.id
WHERE sales.price > 90.00
