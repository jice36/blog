--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2
-- Dumped by pg_dump version 14.0

-- Started on 2022-03-24 23:07:23 MSK

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 24614)
-- Name: notes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notes (
    id_note uuid NOT NULL,
    header text,
    text text,
    tags text[] DEFAULT '{}'::text[] NOT NULL,
    time_create timestamp without time zone,
    id_user uuid
);


ALTER TABLE public.notes OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 24609)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3580 (class 0 OID 24614)
-- Dependencies: 210
-- Data for Name: notes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.notes (id_note, header, text, tags, time_create, id_user) FROM stdin;
\.


--
-- TOC entry 3579 (class 0 OID 24609)
-- Dependencies: 209
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id) FROM stdin;
b03d13da-ab8a-11ec-90e5-acde48001122
b03d1916-ab8a-11ec-90e5-acde48001122
b03d1934-ab8a-11ec-90e5-acde48001122
b03d1948-ab8a-11ec-90e5-acde48001122
b03d1952-ab8a-11ec-90e5-acde48001122
b03d1966-ab8a-11ec-90e5-acde48001122
b03d1970-ab8a-11ec-90e5-acde48001122
b03d1984-ab8a-11ec-90e5-acde48001122
b03d198e-ab8a-11ec-90e5-acde48001122
b03d19a2-ab8a-11ec-90e5-acde48001122
\.


--
-- TOC entry 3438 (class 2606 OID 24621)
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id_note);


--
-- TOC entry 3436 (class 2606 OID 24613)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3439 (class 2606 OID 24622)
-- Name: notes notes_id_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_id_user_fkey FOREIGN KEY (id_user) REFERENCES public.users(id);


-- Completed on 2022-03-24 23:07:23 MSK

--
-- PostgreSQL database dump complete
--

