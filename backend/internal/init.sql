CREATE TABLE IF NOT EXISTS container_status (
                                                id SERIAL PRIMARY KEY,
                                                ip VARCHAR(15) unique NOT NULL,
                                                alive BOOLEAN NOT NULL,
                                                checked TIMESTAMP NOT NULL,
                                                lastSuccess TIMESTAMP NOT NULL
);