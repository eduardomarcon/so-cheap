create table public.itens
(
    id          integer generated always as identity constraint itens_pk primary key,
    description varchar(50) not null,
    amount      integer     not null,
    price       numeric(19, 4)
);
alter table public.itens owner to admin;