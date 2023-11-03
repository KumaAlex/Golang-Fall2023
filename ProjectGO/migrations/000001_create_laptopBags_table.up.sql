CREATE TABLE IF NOT EXISTS laptopBags
(
    id           bigserial PRIMARY KEY,
    created_at   timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    brand        text NOT NULL DEFAULT 'Urban Lifestyle',
    model        text NOT NULL DEFAULT 'Wanderer',
    color        VARCHAR(64),
    material     text DEFAULT 'fabric',
    compartments integer DEFAULT 3,
    weight       numeric(8,2) NOT NULL,
    dimensions   numeric(8,2)[] NOT NULL,
    version integer NOT NULL DEFAULT 1
);