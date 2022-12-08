create table public.orders_itens
(
    id_order integer        not null constraint orders_itens_orders_id_fk references public.orders,
    id_item  integer        not null constraint orders_itens_itens_id_fk references public.itens,
    quantity smallint       not null,
    total    numeric(19, 4) not null,
    constraint orders_itens_pk primary key (id_order, id_item)
);
alter table public.orders_itens owner to admin;