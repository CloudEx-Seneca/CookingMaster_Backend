create database if not exists `recipe`;
use `recipe`;

drop table if exists `recipes`;
create table `recipes` (
    id          bigint not null auto_increment,
    user_id     bigint not null,
    name        varchar(64) not null,
    description varchar(512),
    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    unique key idx_name (name)
);

drop table if exists `ingredients`;
create table `ingredients` (
    id          bigint not null auto_increment,
    recipe_id   bigint not null,
    name        varchar(64) not null,
    quantity    varchar(64),
    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    unique key index_recipe (recipe_id)
);

