select
    b.id as id
from 
    Weather a inner join Weather b 
on 
    DATEDIFF(b.recordDate, a.recordDate) = 1
where 
    b.temperature > a.temperature