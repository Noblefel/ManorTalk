INSERT INTO public.users (email, password, created_at, updated_at, name, username) VALUES 
('test@example.com', '$2y$12$AI3YjckdxoalFHJQBGbQBu2aVNbaNpQO1wewFIaCrY5nMl4tnvYCq', '2020-01-02', '2020-01-02', 'Test User no.1', 'the-first-user');

INSERT INTO users ( 
    email, 
    password,  
    created_at, 
    updated_at,
    name,
    username,
    bio
) 
SELECT 
    concat('test', num, '@example.com'),
    '$2y$12$AI3YjckdxoalFHJQBGbQBu2aVNbaNpQO1wewFIaCrY5nMl4tnvYCq',
    '2020-01-02',
    '2020-01-02',
    concat('Test User no.', num),
    concat('test-user-',num),
    E'🌟 Lorem Ipsum 👋
    \n\n
    🌍 Fames Mauris | ☕ Nullam aliquam | 🎨 Curabitur convallis | 🎵 At porta nibh
    \n\n
    📚 Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla commodo nisl sed odio hendrerit, sit amet dignissim libero fringilla. Fusce vehicula enim eget mauris suscipit, at porta nibh fermentum. Interdum et malesuada fames ac ante ipsum primis in faucibus. Sed euismod turpis eget nisl molestie, id feugiat libero scelerisque. Curabitur convallis augue eu nisi fringilla, id vestibulum velit finibus. Pellentesque sit amet aliquet justo.
    \n\n
    🌱 Vivamus in libero varius, feugiat libero ut, pharetra lorem. Sed nec lacinia ante. Nam at mauris non libero bibendum lobortis. Nullam aliquam erat a tellus placerat, non ultricies libero ultrices. Integer ullamcorper, sem nec ultricies auctor, nisi mauris sollicitudin quam, eu eleifend nunc ipsum sed enim. 🖋️
    \n\n
    📖 Quisque ultricies nunc non urna vestibulum, a elementum justo sollicitudin. Nam commodo libero a sapien faucibus, nec egestas turpis eleifend. Integer ullamcorper consequat ligula, nec pharetra lectus ullamcorper a.
    \n\n
    🎶 Ut at leo eu libero posuere eleifend. Cras congue vestibulum magna, id euismod nisl lacinia non. Donec ac dapibus lectus, eu congue enim. ' 
FROM generate_series(2, 100, 1) as num;

