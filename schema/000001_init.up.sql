CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    user_name     varchar(255) not null unique,
    password_hash varchar(255) not null
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