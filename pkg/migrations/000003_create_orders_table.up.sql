create table public.orders
(
    id      integer generated always as identity constraint orders_pk primary key,
    id_user integer        not null constraint orders_users_id_fk references public.users,
    status  smallint       not null,
    total   numeric(19, 4) not null
);
alter table public.orders owner to admin;