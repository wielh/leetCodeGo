select 
    a.name as Department, b.name as Employee, b.salary as Salary
from 
    Department a inner join Employee b inner join (
        select 
            distinct DENSE_RANK() OVER (PARTITION BY departmentId ORDER BY salary DESC) 
            as department_salary_order, salary, departmentId
        from 
            Employee
    ) c
on 
    a.id = b.departmentId and b.departmentId = c.departmentId and b.salary = c.salary
where 
    c.department_salary_order <= 3