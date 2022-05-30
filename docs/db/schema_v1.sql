CREATE TABLE book (
    id      TEXT    PRIMARY KEY,
    title   TEXT    NOT NULL,
    author  TEXT,
    pages   INTEGER DEFAULT (0),
    deleted BOOLEAN DEFAULT (false) 
);

CREATE INDEX idx_author ON book (
    author
);

CREATE INDEX idx_title ON book (
    title
);
