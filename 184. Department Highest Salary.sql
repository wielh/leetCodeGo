select 
    a.name as Department, b.name as Employee, b.salary as Salary
from 
    Department a inner join Employee b inner join (
        select 
            Max(salary) as max_salary, departmentId
        from 
            Employee
        group by
            departmentId
    ) c
on 
    a.id = b.departmentId and b.departmentId = c.departmentId and b.salary = c.max_salary 