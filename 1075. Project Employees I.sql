
select 
    a.project_id, ROUND(avg(experience_years),2) as average_years 
from 
    Project a inner join Employee b
on
    a.employee_id = b.employee_id 
group by
    a.project_id