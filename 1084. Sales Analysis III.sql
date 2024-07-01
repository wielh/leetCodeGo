select
    distinct a.product_id, a.product_name
from 
    Product a inner join Sales b
on 
    a.product_id = b.product_id
where EXISTS (
    SELECT 
        1 
    FROM 
        Sales
    WHERE 
        b.product_id = Sales.product_id and Sales.sale_date between '2019-01-01' AND '2019-03-31'
) and NOT EXISTS (
    SELECT 
        1 
    FROM 
        Sales
    WHERE 
        b.product_id = Sales.product_id and (Sales.sale_date < '2019-01-01' or Sales.sale_date > '2019-03-31') 
)
