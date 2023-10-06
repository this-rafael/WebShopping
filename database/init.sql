create table products
(
    id          serial primary key,
    name        varchar(100)  not null,
    price       decimal           not null,
    description varchar(1000) not null,
    quantity    int           not null
);