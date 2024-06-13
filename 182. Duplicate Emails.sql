select email as Email from (
    select 
        count(*) as cnt, email 
    from 
        Person 
    group by
        email
    having 
        cnt > 1
) as e
