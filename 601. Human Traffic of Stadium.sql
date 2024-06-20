with 
helper as (
    SELECT *, id - ROW_NUMBER() OVER (PARTITION BY people >= 100 ORDER BY id) as consecutive_group FROM Stadium
),
helper_cnt as (
    select 
        consecutive_group, count(case when helper.people>=100 THEN 1 END) as cnt
    from 
        helper 
    group by 
        helper.consecutive_group 
    having 
        cnt>=3
)

select 
    a.id, a.visit_date, a.people
from 
    helper a inner join helper_cnt b 
on 
    a.consecutive_group = b.consecutive_group
where 
    a.people >= 100
