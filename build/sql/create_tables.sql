CREATE TABLE cities 
(
    id BIGSERIAL PRIMARY KEY,
    city_name TEXT NOT NULL CHECK (city_name <> ''),
    code TEXT NOT NULL CHECK (city_name <> ''),
    country_code TEXT NOT NULL CHECK (city_name <> '')
);