WITH users_count AS (
    SELECT COUNT(*) AS total FROM users 
)

INSERT INTO posts ( 
    user_id, 
    title, 
    slug,  
    content, 
    category_id, 
    created_at, 
    updated_at
) 
SELECT 
    ceil(random() * (SELECT total FROM users_count)),
    concat('A Sample Post no.', num),
    concat('a-sample-post-',num), 
    'A long text: Lorem ipsum dolor sit amet, consectetur adipiscing 
    elit. Nulla posuere neque id magna pretium rutrum. Sed ornare nunc arcu. 
    Cras pharetra, nibh ac ultricies blandit, purus sapien mattis turpis, et 
    congue felis ligula sit amet mi',
    ceil(random() * 8),
    now() - interval '1 hour 14 minute' * (10000 - num),
    now() - interval '1 hour 14 minute' * (10000 - num)
FROM generate_series(1, 10000, 1) as num;

