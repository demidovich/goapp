
-- +goose Up
-- +goose StatementBegin
create sequence if not exists profile_id_seq start 1;
alter sequence profile_id_seq restart with 1;

create table profile (
	id bigint default nextval('profile_id_seq'::regclass) not null,
	email text,
	password_hash text,
	name text,
	created_at timestamp without time zone,
	updated_at timestamp without time zone,
	constraint profile_pkey primary key (id)
);

create unique index profile_email_idx on profile (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists profile;
-- +goose StatementEnd
