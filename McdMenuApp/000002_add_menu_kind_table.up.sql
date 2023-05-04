CREATE TABLE IF NOT EXISTS kind (
    kind_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS item(
    id SERIAL PRIMARY KEY,
    kind_id INTEGER REFERENCES kind(kind_id),
    name VARCHAR(255) NOT NULL,
    energy FLOAT NOT NULL,
    protein FLOAT NOT NULL,
    fat FLOAT NOT NULL,
    carbohydrates FLOAT NOT NULL
);

INSERT INTO kind (name)
VALUES ('ドリンク'), ('バーガー'), ('サイド'), ('バリスタ');
