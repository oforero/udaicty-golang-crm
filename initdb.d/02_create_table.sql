\c udacity_crm;

CREATE TABLE IF NOT EXISTS customers (
    ID          SERIAL  PRIMARY KEY,
    NAME        TEXT    NOT NULL,
    ROLE        TEXT    NOT NULL,
    EMAIL       TEXT    NOT NULL,
    PHONE       TEXT    NOT NULL,
    CONTACTED   BOOLEAN
);
