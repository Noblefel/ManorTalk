CREATE TABLE IF NOT EXISTS public.posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    excerpt VARCHAR(255),
    content TEXT NOT NULL,
    category_id INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO public.posts (user_id, title, slug, excerpt, content, category_id, created_at, updated_at) VALUES 
(1,'The first post title', 'the-first-post-title', 'sample excerpt', 'content', 1, '2020-01-02', '2020-01-02');