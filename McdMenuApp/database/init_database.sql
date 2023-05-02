CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    kind VARCHAR(255),
    energy NUMERIC(10, 2),
    protein NUMERIC(10, 2),
    fat NUMERIC(10, 2),
    carbohydrates NUMERIC(10, 2)
);
