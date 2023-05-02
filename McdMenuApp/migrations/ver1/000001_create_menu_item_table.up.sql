CREATE TABLE IF NOT EXISTS item(
    id SERIAL PRIMARY KEY,
    kind VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    energy FLOAT NOT NULL,
    protein FLOAT NOT NULL,
    fat FLOAT NOT NULL,
    carbohydrates FLOAT NOT NULL
);
