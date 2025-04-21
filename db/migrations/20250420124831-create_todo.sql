
-- +migrate Up
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    priority INT DEFAULT 0,
    done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parent_id INT,
    FOREIGN KEY (parent_id) REFERENCES todos(id) ON DELETE CASCADE
);

CREATE INDEX idx_todos_parent_id ON todos(parent_id);

-- +migrate Down
DROP TABLE todos;