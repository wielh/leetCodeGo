with tmp as (
    select  a.product_id, min(a.year) as first_year from Sales a group by a.product_id
)


select
    a.product_id, a.year as first_year, a.quantity, a.price 
from 
    Sales a 
WHERE EXISTS (
    SELECT 
        1 
    FROM 
        tmp 
    WHERE 
        a.product_id = tmp.product_id and a.first_year = tmp.first_year
);