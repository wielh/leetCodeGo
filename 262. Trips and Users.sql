with temp as(
    select 
        a.request_at as Day, status, count(*) as total, 
        count(CASE WHEN status != 'completed' THEN 1 END) as cancelled
    from 
        Trips a inner join Users b inner join Users c
    on 
        a.client_id = b.users_id and a.driver_id= c.users_id
    where
        a.request_at between "2013-10-01" and "2013-10-03" and b.banned = 'No' and c.banned = 'No' 
    group by 
        a.request_at 
) 

select Day, CASE 
    WHEN total <= 0 THEN 0
    ELSE ROUND(cancelled / total, 2)
    END AS 'Cancellation Rate'
from 
    temp
