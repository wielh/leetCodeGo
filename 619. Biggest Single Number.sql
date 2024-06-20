with result as (
    select 
        num
    from 
        MyNumbers 
    group by 
        num 
    having 
        count(num) = 1 
    order by 
        num desc 
    limit 1
)

SELECT num FROM result RIGHT JOIN (SELECT 1) AS dummy ON 1=1;