CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(255)            NOT NULL,
    email         VARCHAR(255)            NOT NULL UNIQUE,
    user_name     VARCHAR(255)            NOT NULL UNIQUE,
    password_hash VARCHAR(255)            NOT NULL,
    created_at    TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at    TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE todo_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_todo_lists
(
    id           serial                                           not null unique,
    user_id      int references users (id) on delete cascade      not null,
    todo_list_id int references todo_lists (id) on delete cascade not null
);

CREATE TABLE tasks
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);

CREATE TABLE todo_lists_tasks
(
    id           serial                                           not null unique,
    todo_list_id int references todo_lists (id) on delete cascade not null,
    task_id      int references tasks (id) on delete cascade      not null
)