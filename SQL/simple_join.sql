/* For this challenge you need to create a simple SELECT statement that will return all columns from the products table, 
and join to the companies table so that you can return the company name.

products table schema
id
name
isbn
company_id
price
companies table schema
id
name
You should return all product fields as well as the company name as "company_name".

NOTE: Your solution should use pure SQL. Ruby is used within the test cases to do the actual testing. */

CREATE INDEX idx_products_company_id ON products(company_id);
CREATE INDEX idx_company_id ON companies(id);
SELECT 
  prod.id,
  prod.name,
  prod.isbn,
  prod.company_id,
  prod.price,
  comp.name AS company_name 
    FROM products AS prod
    JOIN companies AS comp ON comp.id = prod.company_id;
