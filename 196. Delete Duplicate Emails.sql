with unique_email as (
    select min(id) as id, email from person group by email
)

delete p from 
    person p, unique_email u 
where 
    p.email=u.email and p.id>u.id;