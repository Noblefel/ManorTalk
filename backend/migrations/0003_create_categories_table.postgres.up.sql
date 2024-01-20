CREATE TABLE IF NOT EXISTS public.categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO public.categories (name, slug, created_at, updated_at) VALUES 
('Other', 'other', '2020-01-02', '2020-01-02'),
('Business', 'business', '2020-01-02', '2020-01-02'),
('Culinary', 'culinary', '2020-01-02', '2020-01-02'),
('Education', 'education', '2020-01-02', '2020-01-02'),
('Entertainment', 'entertainment', '2020-01-02', '2020-01-02'),
('Health', 'health', '2020-01-02', '2020-01-02'),
('Nature', 'nature', '2020-01-02', '2020-01-02'),
('Sports', 'sports', '2020-01-02', '2020-01-02'),
('Technology', 'technology', '2020-01-02', '2020-01-02'); 

ALTER TABLE public.posts
    ALTER COLUMN category_id SET DEFAULT 1,
    ADD CONSTRAINT fk_categories
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET DEFAULT;