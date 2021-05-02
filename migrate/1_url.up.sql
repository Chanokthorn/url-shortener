create table if not exists url
(
    short_code text constraint table_name_pk primary key,
    full_url text not null,
    has_expire_date bool not null,
    expire_date timestamp,
    deleted bool not null,
    number_of_hits int not null
);
