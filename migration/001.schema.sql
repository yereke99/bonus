CREATE TABLE IF NOT EXISTS customer(
    id BIGINT PRIMARY KEY,
    user_name VARCHAR(255),
    user_last_name VARCHAR(255),
    email VARCHAR(255),
    locations VARCHAR(255), 
    city VARCHAR(255),
    qr VARCHAR(800), 
    bonus INT,
    isDeleted BOOLEAN 
)

CREATE TABLE IF NOT EXISTS code(
    id BIGINT PRIMARY KEY,
    code INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE TABLE IF NOT EXISTS company(
    id BIGINT PRIMARY KEY,
    company VARCHAR(255),
    company_name VARCHAR(255),
    email VARCHAR(255),
    city VARCHAR(255),
    company_addres VARCHAR(255),
    company_iin INT,
    bonus INT,
    isDeleted BOOLEAN
)

CREATE TABLE IF NOT EXISTS busines_types(
    id BIGINT PRIMARY KEY,
    company_id BIGINT,
    business_type VARCHAR(255),
    city VARCHAR(255),
    email VARCHAR(255),
    working_time VARCHAR(255),
    trc VARCHAR(255),
    business_address VARCHAR(255),
    floor INT,
    business_line VARCHAR(255),
    business_number INT
)




