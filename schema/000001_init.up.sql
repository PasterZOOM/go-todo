CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serial                                      not null unique,
    title       varchar(255)                                not null,
    description varchar(255),
    user_id     int references users (id) on delete cascade not null
);

CREATE TABLE tasks
(
    id           serial                                           not null unique,
    title        varchar(255)                                     not null,
    description  varchar(255),
    todo_list_id int references todo_lists (id) on delete cascade not null
);