CREATE TABLE IF NOT EXISTS public.users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO public.users (email, password, created_at, updated_at) VALUES 
('test@example.com', '$2y$12$AI3YjckdxoalFHJQBGbQBu2aVNbaNpQO1wewFIaCrY5nMl4tnvYCq', '2020-01-02', '2020-01-02');