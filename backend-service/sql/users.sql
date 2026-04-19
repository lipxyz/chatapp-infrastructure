-- DDL for users table
-- Use ddl migrations when working with Kubernetes module
-- Set MIGRATE_IN_CODE=false in .env (disable auto migrate app)

create table if not exists users
(
    id         bigserial
        primary key,
    username   text not null
        constraint uni_users_username
            unique,
    email      text not null
        constraint uni_users_email
            unique,
    password   text not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);