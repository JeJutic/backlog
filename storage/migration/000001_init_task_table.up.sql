CREATE TABLE task_status
(
    id INTEGER PRIMARY KEY,
    name varchar(20) NOT NULL
);

CREATE TABLE tasks
(
    id INTEGER PRIMARY KEY,
    text varchar NOT NULL,
    status_id INTEGER NOT NULL,
    FOREIGN KEY(status_id) REFERENCES task_status(id)
);

INSERT INTO task_status(name) VALUES
    ("To do"),
    ("In progress"),
    ("Done")