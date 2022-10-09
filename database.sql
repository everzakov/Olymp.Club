--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5
-- Dumped by pg_dump version 14.5

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
-- Name: AdminModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."AdminModel" (
    id integer NOT NULL,
    user_id integer,
    priority integer NOT NULL
);


ALTER TABLE public."AdminModel" OWNER TO efimverzakov;

--
-- Name: AdminModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."AdminModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."AdminModel_id_seq" OWNER TO efimverzakov;

--
-- Name: AdminModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."AdminModel_id_seq" OWNED BY public."AdminModel".id;


--
-- Name: BigOlympiadModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."BigOlympiadModel" (
    big_olympiad_id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    short text DEFAULT ''::text NOT NULL,
    logo text DEFAULT ''::text,
    description text DEFAULT ''::text,
    status text DEFAULT ''::text
);


ALTER TABLE public."BigOlympiadModel" OWNER TO efimverzakov;

--
-- Name: BigOlympiadModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."BigOlympiadModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."BigOlympiadModel_id_seq" OWNER TO efimverzakov;

--
-- Name: BigOlympiadModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."BigOlympiadModel_id_seq" OWNED BY public."BigOlympiadModel".big_olympiad_id;


--
-- Name: EventModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."EventModel" (
    id integer NOT NULL,
    "Name" text DEFAULT ''::text NOT NULL,
    "Description" text DEFAULT ''::text NOT NULL,
    "Short" text DEFAULT ''::text NOT NULL,
    "Img" text DEFAULT ''::text NOT NULL,
    "Status" text DEFAULT ''::text NOT NULL,
    "HolderId" integer,
    "Website" text DEFAULT ''::text NOT NULL
);


ALTER TABLE public."EventModel" OWNER TO efimverzakov;

--
-- Name: EventModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."EventModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."EventModel_id_seq" OWNER TO efimverzakov;

--
-- Name: EventModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."EventModel_id_seq" OWNED BY public."EventModel".id;


--
-- Name: EventUserModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."EventUserModel" (
    id integer NOT NULL,
    event_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public."EventUserModel" OWNER TO efimverzakov;

--
-- Name: HolderModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."HolderModel" (
    holder_id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    logo text DEFAULT ''::text NOT NULL
);


ALTER TABLE public."HolderModel" OWNER TO efimverzakov;

--
-- Name: HolderModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."HolderModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."HolderModel_id_seq" OWNER TO efimverzakov;

--
-- Name: HolderModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."HolderModel_id_seq" OWNED BY public."HolderModel".holder_id;


--
-- Name: NewsModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."NewsModel" (
    id integer NOT NULL,
    "Title" text DEFAULT ''::text NOT NULL,
    "Description" text DEFAULT ''::text NOT NULL,
    "TableStruct" text DEFAULT ''::text NOT NULL,
    "Key" integer NOT NULL
);


ALTER TABLE public."NewsModel" OWNER TO efimverzakov;

--
-- Name: NewsModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."NewsModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."NewsModel_id_seq" OWNER TO efimverzakov;

--
-- Name: NewsModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."NewsModel_id_seq" OWNED BY public."NewsModel".id;


--
-- Name: OlympiadModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."OlympiadModel" (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    subject text DEFAULT ''::text NOT NULL,
    level text DEFAULT ''::text NOT NULL,
    img text DEFAULT ''::text NOT NULL,
    short text DEFAULT ''::text NOT NULL,
    big_olympiad_id integer NOT NULL,
    status text NOT NULL,
    grade text DEFAULT ''::text NOT NULL,
    holder_id integer,
    website text DEFAULT ''::text NOT NULL
);


ALTER TABLE public."OlympiadModel" OWNER TO efimverzakov;

--
-- Name: OlympiadModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."OlympiadModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."OlympiadModel_id_seq" OWNER TO efimverzakov;

--
-- Name: OlympiadModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."OlympiadModel_id_seq" OWNED BY public."OlympiadModel".id;


--
-- Name: OlympiadUserModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."OlympiadUserModel" (
    id integer NOT NULL,
    olympiad_id integer,
    user_id integer
);


ALTER TABLE public."OlympiadUserModel" OWNER TO efimverzakov;

--
-- Name: OlympiadUserModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."OlympiadUserModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."OlympiadUserModel_id_seq" OWNER TO efimverzakov;

--
-- Name: OlympiadUserModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."OlympiadUserModel_id_seq" OWNED BY public."OlympiadUserModel".id;


--
-- Name: SessionModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."SessionModel" (
    id integer NOT NULL,
    user_id integer NOT NULL,
    token text DEFAULT ''::text NOT NULL,
    expiry timestamp without time zone DEFAULT (now() + '01:00:00'::interval)
);


ALTER TABLE public."SessionModel" OWNER TO efimverzakov;

--
-- Name: SessiomModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."SessiomModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SessiomModel_id_seq" OWNER TO efimverzakov;

--
-- Name: SessiomModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."SessiomModel_id_seq" OWNED BY public."SessionModel".id;


--
-- Name: UnConfirmedUsers; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."UnConfirmedUsers" (
    id integer NOT NULL,
    email text DEFAULT ''::text NOT NULL,
    token1 text DEFAULT ''::text NOT NULL,
    token2 text DEFAULT ''::text NOT NULL,
    pass_hash text DEFAULT '''''::text'::text NOT NULL,
    confirmed boolean
);


ALTER TABLE public."UnConfirmedUsers" OWNER TO efimverzakov;

--
-- Name: UnConfirmedUsers_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."UnConfirmedUsers_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."UnConfirmedUsers_id_seq" OWNER TO efimverzakov;

--
-- Name: UnConfirmedUsers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."UnConfirmedUsers_id_seq" OWNED BY public."UnConfirmedUsers".id;


--
-- Name: UserModel; Type: TABLE; Schema: public; Owner: efimverzakov
--

CREATE TABLE public."UserModel" (
    id integer NOT NULL,
    email text DEFAULT ''::text NOT NULL,
    pass_hash text DEFAULT ''::text NOT NULL,
    token1 text DEFAULT ''::text,
    token2 text DEFAULT ''::text
);


ALTER TABLE public."UserModel" OWNER TO efimverzakov;

--
-- Name: UserModel_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public."UserModel_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."UserModel_id_seq" OWNER TO efimverzakov;

--
-- Name: UserModel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public."UserModel_id_seq" OWNED BY public."UserModel".id;


--
-- Name: untitled_table_213_id_seq; Type: SEQUENCE; Schema: public; Owner: efimverzakov
--

CREATE SEQUENCE public.untitled_table_213_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.untitled_table_213_id_seq OWNER TO efimverzakov;

--
-- Name: untitled_table_213_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: efimverzakov
--

ALTER SEQUENCE public.untitled_table_213_id_seq OWNED BY public."EventUserModel".id;


--
-- Name: AdminModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."AdminModel" ALTER COLUMN id SET DEFAULT nextval('public."AdminModel_id_seq"'::regclass);


--
-- Name: BigOlympiadModel big_olympiad_id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."BigOlympiadModel" ALTER COLUMN big_olympiad_id SET DEFAULT nextval('public."BigOlympiadModel_id_seq"'::regclass);


--
-- Name: EventModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."EventModel" ALTER COLUMN id SET DEFAULT nextval('public."EventModel_id_seq"'::regclass);


--
-- Name: EventUserModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."EventUserModel" ALTER COLUMN id SET DEFAULT nextval('public.untitled_table_213_id_seq'::regclass);


--
-- Name: HolderModel holder_id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."HolderModel" ALTER COLUMN holder_id SET DEFAULT nextval('public."HolderModel_id_seq"'::regclass);


--
-- Name: NewsModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."NewsModel" ALTER COLUMN id SET DEFAULT nextval('public."NewsModel_id_seq"'::regclass);


--
-- Name: OlympiadModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadModel" ALTER COLUMN id SET DEFAULT nextval('public."OlympiadModel_id_seq"'::regclass);


--
-- Name: OlympiadUserModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadUserModel" ALTER COLUMN id SET DEFAULT nextval('public."OlympiadUserModel_id_seq"'::regclass);


--
-- Name: SessionModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."SessionModel" ALTER COLUMN id SET DEFAULT nextval('public."SessiomModel_id_seq"'::regclass);


--
-- Name: UnConfirmedUsers id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."UnConfirmedUsers" ALTER COLUMN id SET DEFAULT nextval('public."UnConfirmedUsers_id_seq"'::regclass);


--
-- Name: UserModel id; Type: DEFAULT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."UserModel" ALTER COLUMN id SET DEFAULT nextval('public."UserModel_id_seq"'::regclass);


--
-- Name: AdminModel AdminModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."AdminModel"
    ADD CONSTRAINT "AdminModel_pkey" PRIMARY KEY (id);


--
-- Name: BigOlympiadModel BigOlympiadModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."BigOlympiadModel"
    ADD CONSTRAINT "BigOlympiadModel_pkey" PRIMARY KEY (big_olympiad_id);


--
-- Name: EventModel EventModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."EventModel"
    ADD CONSTRAINT "EventModel_pkey" PRIMARY KEY (id);


--
-- Name: HolderModel HolderModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."HolderModel"
    ADD CONSTRAINT "HolderModel_pkey" PRIMARY KEY (holder_id);


--
-- Name: NewsModel NewsModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."NewsModel"
    ADD CONSTRAINT "NewsModel_pkey" PRIMARY KEY (id);


--
-- Name: OlympiadModel OlympiadModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadModel"
    ADD CONSTRAINT "OlympiadModel_pkey" PRIMARY KEY (id);


--
-- Name: OlympiadUserModel OlympiadUserModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadUserModel"
    ADD CONSTRAINT "OlympiadUserModel_pkey" PRIMARY KEY (id);


--
-- Name: SessionModel SessiomModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."SessionModel"
    ADD CONSTRAINT "SessiomModel_pkey" PRIMARY KEY (id);


--
-- Name: UnConfirmedUsers UnConfirmedUsers_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."UnConfirmedUsers"
    ADD CONSTRAINT "UnConfirmedUsers_pkey" PRIMARY KEY (id);


--
-- Name: UserModel UserModel_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."UserModel"
    ADD CONSTRAINT "UserModel_pkey" PRIMARY KEY (id);


--
-- Name: EventUserModel untitled_table_213_pkey; Type: CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."EventUserModel"
    ADD CONSTRAINT untitled_table_213_pkey PRIMARY KEY (id);


--
-- Name: AdminModel AdminModel_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."AdminModel"
    ADD CONSTRAINT "AdminModel_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."UserModel"(id);


--
-- Name: EventModel EventModel_HolderId_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."EventModel"
    ADD CONSTRAINT "EventModel_HolderId_fkey" FOREIGN KEY ("HolderId") REFERENCES public."HolderModel"(holder_id);


--
-- Name: EventUserModel EventUserModel_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."EventUserModel"
    ADD CONSTRAINT "EventUserModel_event_id_fkey" FOREIGN KEY (event_id) REFERENCES public."EventModel"(id);


--
-- Name: EventUserModel EventUserModel_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."EventUserModel"
    ADD CONSTRAINT "EventUserModel_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."UserModel"(id);


--
-- Name: OlympiadModel OlympiadModel_holder_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadModel"
    ADD CONSTRAINT "OlympiadModel_holder_id_fkey" FOREIGN KEY (holder_id) REFERENCES public."HolderModel"(holder_id);


--
-- Name: OlympiadUserModel OlympiadUserModel_olympiad_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadUserModel"
    ADD CONSTRAINT "OlympiadUserModel_olympiad_id_fkey" FOREIGN KEY (olympiad_id) REFERENCES public."OlympiadModel"(id);


--
-- Name: OlympiadUserModel OlympiadUserModel_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadUserModel"
    ADD CONSTRAINT "OlympiadUserModel_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."UserModel"(id);


--
-- Name: SessionModel SessiomModel_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."SessionModel"
    ADD CONSTRAINT "SessiomModel_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."UserModel"(id);


--
-- Name: OlympiadModel big_olympiad_id; Type: FK CONSTRAINT; Schema: public; Owner: efimverzakov
--

ALTER TABLE ONLY public."OlympiadModel"
    ADD CONSTRAINT big_olympiad_id FOREIGN KEY (big_olympiad_id) REFERENCES public."BigOlympiadModel"(big_olympiad_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

