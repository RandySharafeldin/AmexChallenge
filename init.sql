create table if not exists Orders(
    id serial primary key,
    apples int not null,
    oranges int not null,
    cost float(32) not null
);