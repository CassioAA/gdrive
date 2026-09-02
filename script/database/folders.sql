CREATE TABLE folders (
    id SERIAL,
    parent_id INT,
    name VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    modified_at TIMESTAMP NOT NULL,
    deleted BOOL NOT NULL DEFAULT false,
    PRIMARY KEY (id),
    /*
    The two lines below the CONSTRAINT line form a rule and the CONSTRAINT line name the rule as fk_parent.
    FOREIGN KEY sets the parent_id column as a column of foreign keys values.
    REFERENCES sets folders table id column values as parent_id foreign keys.
    */
    CONSTRAINT fk_parent
        FOREIGN KEY (parent_id)
            REFERENCES folders (id)
)
