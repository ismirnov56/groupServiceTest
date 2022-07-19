CREATE TABLE groups
(
    id        serial       not null unique,
    name      varchar(255) not null,
    parent_id int references groups (id) on delete cascade
);

CREATE TABLE users
(
    id         serial       not null unique,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    birth_year int          not null,
    group_id   int          references groups (id) on delete set null
);