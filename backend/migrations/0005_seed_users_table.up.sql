INSERT INTO public.users (email, password, created_at, updated_at, name, username) VALUES 
('test@example.com', '$2y$12$AI3YjckdxoalFHJQBGbQBu2aVNbaNpQO1wewFIaCrY5nMl4tnvYCq', '2020-01-02', '2020-01-02', 'Test User no.1', 'the-first-user');

INSERT INTO users ( 
    email, 
    password,  
    created_at, 
    updated_at,
    name,
    username
) 
SELECT 
    concat('test', num, '@example.com'),
    '$2y$12$AI3YjckdxoalFHJQBGbQBu2aVNbaNpQO1wewFIaCrY5nMl4tnvYCq',
    '2020-01-02',
    '2020-01-02',
    concat('Test User no.', num),
    concat('test-user-',num) 
FROM generate_series(2, 100, 1) as num;

