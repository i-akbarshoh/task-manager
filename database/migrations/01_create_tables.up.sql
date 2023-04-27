CREATE TYPE role_type AS ENUM ('teamlead', 'programmer', 'intern');
CREATE TYPE position_type AS ENUM ('director', 'admin', 'programmer');
CREATE TYPE status_type AS ENUM ('new', 'in_process', 'finished');
CREATE TYPE attendance_type AS ENUM ('come', 'gone');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    avatar TEXT,
    role role_type NOT NULL,
    birth_date DATE NOT NULL,
    phone TEXT NOT NULL,
    position position_type NOT NULL
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    started_date DATE NOT NULL DEFAULT CURRENT_DATE,
    finished_date DATE,
    status status_type NOT NULL DEFAULT 'new',
    teamlead_id UUID NOT NULL REFERENCES users(id),
    attachment TEXT
);

CREATE TABLE attendance (
    type attendance_type NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    date DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    start_at DATE NOT NULL,
    finish_at DATE NOT NULL,
    started_at DATE NOT NULL DEFAULT CURRENT_DATE,
    finished_at DATE,
    status status_type NOT NULL DEFAULT 'new',
    programmer_id UUID NOT NULL REFERENCES users(id),
    project_id INT NOT NULL REFERENCES projects(id),
    attachment TEXT
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    programmer_id UUID NOT NULL REFERENCES users(id),
    task_id INT NOT NULL REFERENCES tasks(id)
);