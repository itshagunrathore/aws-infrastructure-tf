create type backup_types as enum ('FULL', 'DELTA', 'CUMULATIVE')
;

create type site_target_type as enum ('AWS', 'AZURE', 'GCP')
;

create type job_type as enum ('BACKUP', 'RESTORE')
;

create type retention_source as enum ('S3', 'DATA_DOMAIN')
;



create table customer_site
(
	customer_name text,
	site_id text not null,
	customer_site_id serial not null
		constraint customer_site_pkey
			primary key,
	system_name text,
	server_uuid uuid
		constraint not null,
	site_target_type site_target_type not null,
	site_region text,
	opt_in_email_report boolean default false,
	opt_in_retry_flag boolean default true,
	is_deleted boolean default false,
	created_at timestamp,
	updated_at timestamp,
	onboarded_date timestamp,
    offering_type text,
	constraint customer_site_unique_key
		unique (customer_site_id, site_id)
)
;

create table job_definition
(
	job_id serial not null
		constraint job_definition_pkey
			primary key,
	time_updated timestamp,
	job_name text,
	is_active boolean default false,
	time_created timestamp,
	description text,
	status text,
	is_deleted boolean default false,
	is_hidden boolean default false,
	backup_type backup_types not null,
	job_type job_type not null,
	customer_site_id integer
		constraint job_definition_customer_site_id_fkey
			references customer_site,
	retention_copies_count integer
		constraint null_or_non_negative_retention_count
			check ((retention_copies_count IS NULL) OR (retention_copies_count >= 0)),
	retention_source retention_source not null
)
;