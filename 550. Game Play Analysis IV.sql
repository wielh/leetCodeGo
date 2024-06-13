with tmp as (
    select 
        distinct a.player_id as player_id
    from 
        Activity a inner join (
            select 
                player_id, min(event_date) as event_date
            from 
                Activity
            group by
                player_id
        ) b
    on
        a.player_id = b.player_id and DATEDIFF(a.event_date, b.event_date) = 1 
)

SELECT 
    round((SELECT COUNT(player_id) FROM tmp) / (SELECT COUNT(distinct player_id) FROM Activity),2) 
AS fraction;