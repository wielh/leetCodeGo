select
    a.id, 
    CASE
        WHEN a.p_id IS NULL THEN "Root" 
        WHEN exists (select 1 from Tree b where a.id = b.p_id) THEN "Inner" 
        ELSE "Leaf" 
    END as type
from
    Tree a
