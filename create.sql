create table blog(
	   id integer primary key,
	   title text,
	   body text,
	   tags_string text,
	   created_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime')),
	   updated_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
);
create table comment(
	   id integer primary key,
	   comment text,
	   blog_id integer,
	   username text,
	   created_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime')),
	   updated_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
);
create table user(
	   id integer primary key,
	   comment text,
	   blog_id integer,
	   username text,
	   created_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime')),
	   updated_at TIMESTAMP DEFAULT (datetime(CURRENT_TIMESTAMP,'localtime'))
);
