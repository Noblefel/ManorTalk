ALTER TABLE public.posts
    DROP CONSTRAINT IF EXISTS fk_categories;

DROP TABLE IF EXISTS categories;