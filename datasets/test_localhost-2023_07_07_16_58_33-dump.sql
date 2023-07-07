--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.3 (Ubuntu 15.3-0ubuntu0.23.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: credit_limits; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.credit_limits (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    student_id uuid NOT NULL,
    credit_limit integer DEFAULT 1000 NOT NULL,
    create_at timestamp with time zone DEFAULT now(),
    update_at timestamp with time zone
);


ALTER TABLE public.credit_limits OWNER TO root;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO root;

--
-- Name: students; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.students (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    full_name text NOT NULL,
    group_num text NOT NULL,
    email text NOT NULL,
    username text NOT NULL,
    verify_email boolean DEFAULT false,
    create_at timestamp with time zone DEFAULT now(),
    update_at timestamp with time zone
);


ALTER TABLE public.students OWNER TO root;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.tasks (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    description text NOT NULL,
    cost integer NOT NULL,
    create_at timestamp with time zone DEFAULT now(),
    update_at timestamp with time zone
);


ALTER TABLE public.tasks OWNER TO root;

--
-- Data for Name: credit_limits; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.credit_limits (id, student_id, credit_limit, create_at, update_at) FROM stdin;
6fdc901b-5820-4c8b-ad6c-b5b1b9e42df4	5218a4ac-f8f2-47bd-b2dd-11b3041766c8	1000	2023-07-07 10:46:26.917752+00	\N
6a311194-6067-4a2e-9570-f2fab941dfff	1621d82b-ba2b-413d-929c-2a6014557c2e	1013	2023-07-07 13:17:20.894847+00	\N
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.schema_migrations (version, dirty) FROM stdin;
20230707110520	f
\.


--
-- Data for Name: students; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.students (id, full_name, group_num, email, username, verify_email, create_at, update_at) FROM stdin;
5218a4ac-f8f2-47bd-b2dd-11b3041766c8	asdsad Dmitrij asdasd	ANTM-23m	qweqw@qwe.com	asdqw	f	2023-07-07 10:46:26.912952+00	\N
1621d82b-ba2b-413d-929c-2a6014557c2e	Dmitrij Semenkin 	ANTM-23m	dmitrijsemenkin@gmail.com	SDA	f	2023-07-07 13:17:20.88372+00	\N
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.tasks (id, description, cost, create_at, update_at) FROM stdin;
48acfe1a-9bcb-41fd-b5d6-e34f14c10f87	Дан неотсортированный массив из N чисел от 1 до N,\nпри этом несколько чисел из диапазона [1, N] пропущено,\nа некоторые присутствуют дважды.\n\nНайти все пропущенные числа.	100	2023-07-07 11:06:48.940309+00	\N
\.


--
-- Name: credit_limits credit_limits_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.credit_limits
    ADD CONSTRAINT credit_limits_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: students students_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_pkey PRIMARY KEY (id);


--
-- Name: students students_username_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.students
    ADD CONSTRAINT students_username_key UNIQUE (username);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: idx_credit_limits_student_id; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_credit_limits_student_id ON public.credit_limits USING btree (student_id);


--
-- Name: idx_tasks_id; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_tasks_id ON public.tasks USING btree (id);


--
-- Name: credit_limits credit_limits_student_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.credit_limits
    ADD CONSTRAINT credit_limits_student_id_fkey FOREIGN KEY (student_id) REFERENCES public.students(id);


--
-- PostgreSQL database dump complete
--

