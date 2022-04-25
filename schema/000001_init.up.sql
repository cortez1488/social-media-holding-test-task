CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL,
    chatID NUMERIC NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS ip_info(
    id SERIAL PRIMARY KEY,
    ip VARCHAR(255),
    continent_name VARCHAR(255) NOT NULL,
    country_name VARCHAR(255),
    region_name VARCHAR(255),
    city VARCHAR(255),
    zip VARCHAR(255),
    latitude DECIMAL,
    longitude DECIMAL
);

CREATE TABLE IF NOT EXISTS user_searched_ip(
    id SERIAL PRIMARY KEY,
    ip_id int REFERENCES ip_info(id),
    user_id int REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS search_date(
    id SERIAL PRIMARY KEY,
    user_searched_ip_id int REFERENCES user_searched_ip(id),
    timedate TIMESTAMP
);

















