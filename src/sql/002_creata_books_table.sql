CREATE TABLE public.books (
	id serial NOT NULL,
	book_name varchar NOT NULL,
	edition int4 NOT NULL,
	publication_year int4 NOT NULL,
	authors _int4 NOT NULL,
	CONSTRAINT books_pk PRIMARY KEY (id)
);
