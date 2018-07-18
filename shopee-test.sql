--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.13
-- Dumped by pg_dump version 9.5.13

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: currency; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.currency (
    date date,
    id_exchange integer,
    value double precision
);


ALTER TABLE public.currency OWNER TO postgres;

--
-- Name: currency_exchange; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.currency_exchange (
    id_unit integer,
    id_unit_target integer,
    status smallint,
    id_exchange integer NOT NULL
);


ALTER TABLE public.currency_exchange OWNER TO postgres;

--
-- Name: currency_exchange_id_exchange_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.currency_exchange_id_exchange_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.currency_exchange_id_exchange_seq OWNER TO postgres;

--
-- Name: currency_exchange_id_exchange_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.currency_exchange_id_exchange_seq OWNED BY public.currency_exchange.id_exchange;


--
-- Name: currency_unit; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.currency_unit (
    unit character(3),
    description character varying(100),
    id_unit integer NOT NULL
);


ALTER TABLE public.currency_unit OWNER TO postgres;

--
-- Name: currency_unit_id_unit_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.currency_unit_id_unit_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.currency_unit_id_unit_seq OWNER TO postgres;

--
-- Name: currency_unit_id_unit_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.currency_unit_id_unit_seq OWNED BY public.currency_unit.id_unit;


--
-- Name: id_exchange; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.currency_exchange ALTER COLUMN id_exchange SET DEFAULT nextval('public.currency_exchange_id_exchange_seq'::regclass);


--
-- Name: id_unit; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.currency_unit ALTER COLUMN id_unit SET DEFAULT nextval('public.currency_unit_id_unit_seq'::regclass);


--
-- Data for Name: currency; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.currency (date, id_exchange, value) FROM stdin;
2018-07-10	1	0.564230000000000009
\.


--
-- Data for Name: currency_exchange; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.currency_exchange (id_unit, id_unit_target, status, id_exchange) FROM stdin;
1	2	1	1
2	1	0	2
1	4	1	3
\.


--
-- Name: currency_exchange_id_exchange_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.currency_exchange_id_exchange_seq', 3, true);


--
-- Data for Name: currency_unit; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.currency_unit (unit, description, id_unit) FROM stdin;
USD	Dollar Amerika	1
SGD	Dollar Singapura	2
IDR	Indonesia	3
IDR	United States	4
SGD	Singapore	5
JPY	Japan	6
\.


--
-- Name: currency_unit_id_unit_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.currency_unit_id_unit_seq', 6, true);


--
-- Name: currency_exchange_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.currency_exchange
    ADD CONSTRAINT currency_exchange_pkey PRIMARY KEY (id_exchange);


--
-- Name: currency_unit_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.currency_unit
    ADD CONSTRAINT currency_unit_pkey PRIMARY KEY (id_unit);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

