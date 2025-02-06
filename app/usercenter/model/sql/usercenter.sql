create database if not exists `usercenter`;
use `usercenter`;

drop table if exists `users`;
create table `users` (
    id          bigint not null comment 'snowflake ID',
    email       varchar(64) not null comment 'user email',
    password    varchar(255) not null comment 'argon2 hash',
    status      tinyint default 0 comment '0:unvarified, 1:varified',
    nickname    varchar(64) comment 'user nickname',
    sex         tinyint default 0 comment '0:unknown, 1:male, 2:female',
    avatar_url  varchar(255) comment 'avatar url',
    info        varchar(255) comment 'user info',
    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    unique key idx_email (email),
    unique key idx_nickname (nickname)
);

drop table `user_tokens`;
create table `user_tokens` (
    id          bigint not null auto_increment comment 'token id',
    user_id     bigint not null comment 'user id',
    token       varchar(512) not null comment 'user token',
    type        tinyint not null comment '0:register, 1:reset, 2:refresh',
    status      tinyint not null default 0 comment '0:active, 1:used, 2:expired, 3:revoked',
    expire      datetime not null comment 'token expire',
    created_at  timestamp default current_timestamp,
    updated_at  timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    unique key idx_user_type (user_id, type)
)
