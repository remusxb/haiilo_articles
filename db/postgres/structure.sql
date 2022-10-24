CREATE TABLE IF NOT EXISTS articles
(
    id uuid NOT NULL UNIQUE,
    title character varying(500) NOT NULL UNIQUE,
    link character varying(1000) NOT NULL UNIQUE,
    PRIMARY KEY (id)
)

