with tmp as (
    select 
        managerId, count(*) as cnt
    from 
        Employee
    where
        managerId IS NOT NULL
    group by 
        managerId
    having
        cnt >= 5
)

select 
    a.name
from 
    Employee a inner join tmp b
on 
    a.id = b.managerId
