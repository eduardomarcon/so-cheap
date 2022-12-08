create table public.users
(
    id       integer generated always as identity constraint users_pk primary key,
    name     varchar(50) not null,
    email    varchar(50) not null,
    password varchar(60) not null
);
alter table public.users owner to admin;