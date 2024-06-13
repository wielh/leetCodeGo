with tmp as (
    select
        empId, sum(bonus) as total_bonus 
    from 
        Bonus
    group by 
        empId
)

select 
    a.name, b.total_bonus as bonus 
from 
    Employee a left join tmp b
on 
    a.empId = b.empId
where
    b.total_bonus is null or b.total_bonus < 1000