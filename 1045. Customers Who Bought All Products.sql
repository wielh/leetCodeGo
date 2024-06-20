select 
    a.customer_id 
from
    (
        select 
            distinct product_key,customer_id 
        from
            Customer
    ) as a
group by
    a.customer_id
having
    count(a.product_key) = (SELECT COUNT(*) as cnt FROM Product)
