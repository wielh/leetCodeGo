with tmp as (
    select 
        class, count(class) as cnt
    from
        Courses
    group by
        class
    having
        cnt >= 5
)

select class from tmp