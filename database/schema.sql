--
-- PostgreSQL database dump
--

\restrict af0SbR3Bqk8fqeuO8A6cdVSCrr43cvY0NbcFuEZYM5tdsmR6stueFOSn78a30zO

-- Dumped from database version 18.4
-- Dumped by pg_dump version 18.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: alerts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alerts (
    id integer NOT NULL,
    alert_time timestamp without time zone DEFAULT now(),
    file_name text,
    risk_level text,
    verdict text,
    message text,
    campaign_ioc text
);


ALTER TABLE public.alerts OWNER TO postgres;

--
-- Name: alerts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.alerts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.alerts_id_seq OWNER TO postgres;

--
-- Name: alerts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.alerts_id_seq OWNED BY public.alerts.id;


--
-- Name: analysis_results; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.analysis_results (
    id integer NOT NULL,
    file_name text,
    risk_score integer,
    risk_level text,
    verdict text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    findings text,
    sha256 text,
    urls text,
    mitre text
);


ALTER TABLE public.analysis_results OWNER TO postgres;

--
-- Name: analysis_results_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.analysis_results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.analysis_results_id_seq OWNER TO postgres;

--
-- Name: analysis_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.analysis_results_id_seq OWNED BY public.analysis_results.id;


--
-- Name: analyst_notes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.analyst_notes (
    id integer NOT NULL,
    ioc text,
    analyst text,
    notes text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.analyst_notes OWNER TO postgres;

--
-- Name: analyst_notes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.analyst_notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.analyst_notes_id_seq OWNER TO postgres;

--
-- Name: analyst_notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.analyst_notes_id_seq OWNED BY public.analyst_notes.id;


--
-- Name: campaigns; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.campaigns (
    id integer NOT NULL,
    campaign_name text,
    ioc text,
    occurrence_count integer,
    sources text,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.campaigns OWNER TO postgres;

--
-- Name: campaigns_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.campaigns_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.campaigns_id_seq OWNER TO postgres;

--
-- Name: campaigns_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.campaigns_id_seq OWNED BY public.campaigns.id;


--
-- Name: case_notes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.case_notes (
    id integer NOT NULL,
    case_id integer,
    analyst text,
    note text,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.case_notes OWNER TO postgres;

--
-- Name: case_notes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.case_notes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.case_notes_id_seq OWNER TO postgres;

--
-- Name: case_notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.case_notes_id_seq OWNED BY public.case_notes.id;


--
-- Name: cases; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cases (
    id integer NOT NULL,
    file_name text,
    analyst text,
    status text,
    notes text,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.cases OWNER TO postgres;

--
-- Name: cases_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cases_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cases_id_seq OWNER TO postgres;

--
-- Name: cases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cases_id_seq OWNED BY public.cases.id;


--
-- Name: docker_reports; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.docker_reports (
    id integer NOT NULL,
    sandbox_job_id integer,
    container_status character varying(50),
    execution_status character varying(50),
    duration_seconds integer,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.docker_reports OWNER TO postgres;

--
-- Name: docker_reports_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.docker_reports_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.docker_reports_id_seq OWNER TO postgres;

--
-- Name: docker_reports_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.docker_reports_id_seq OWNED BY public.docker_reports.id;


--
-- Name: emails; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.emails (
    id integer NOT NULL,
    sender text,
    subject text,
    body text,
    risk_score integer,
    decision text,
    findings text,
    attachments text,
    analyst_note text,
    mitre_technique text
);


ALTER TABLE public.emails OWNER TO postgres;

--
-- Name: emails_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.emails_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.emails_id_seq OWNER TO postgres;

--
-- Name: emails_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.emails_id_seq OWNED BY public.emails.id;


--
-- Name: ioc_correlation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ioc_correlation (
    id integer NOT NULL,
    ioc text,
    source_type text,
    file_name text,
    created_at timestamp without time zone DEFAULT now(),
    first_seen timestamp without time zone DEFAULT now(),
    last_seen timestamp without time zone DEFAULT now(),
    hit_count integer DEFAULT 1
);


ALTER TABLE public.ioc_correlation OWNER TO postgres;

--
-- Name: ioc_correlation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ioc_correlation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ioc_correlation_id_seq OWNER TO postgres;

--
-- Name: ioc_correlation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ioc_correlation_id_seq OWNED BY public.ioc_correlation.id;


--
-- Name: mitre_events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mitre_events (
    id integer NOT NULL,
    file_name text,
    technique text
);


ALTER TABLE public.mitre_events OWNER TO postgres;

--
-- Name: mitre_events_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mitre_events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.mitre_events_id_seq OWNER TO postgres;

--
-- Name: mitre_events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mitre_events_id_seq OWNED BY public.mitre_events.id;


--
-- Name: sandbox_jobs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sandbox_jobs (
    id integer NOT NULL,
    file_name text,
    file_path text,
    status text,
    submitted_at timestamp without time zone DEFAULT now(),
    completed_at timestamp without time zone
);


ALTER TABLE public.sandbox_jobs OWNER TO postgres;

--
-- Name: sandbox_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sandbox_jobs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sandbox_jobs_id_seq OWNER TO postgres;

--
-- Name: sandbox_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sandbox_jobs_id_seq OWNED BY public.sandbox_jobs.id;


--
-- Name: sandbox_reports; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sandbox_reports (
    id integer NOT NULL,
    job_id integer,
    findings text,
    risk_score integer,
    risk_level text,
    verdict text,
    mitre text,
    created_at timestamp without time zone DEFAULT now(),
    file_name text,
    file_size BIGINT NOT NULL CHECK (file_size > 0),
    extension text,
    mime_type text,
    md5 text,
    sha256 text
);


ALTER TABLE public.sandbox_reports OWNER TO postgres;

--
-- Name: sandbox_reports_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sandbox_reports_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sandbox_reports_id_seq OWNER TO postgres;

--
-- Name: sandbox_reports_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sandbox_reports_id_seq OWNED BY public.sandbox_reports.id;


--
-- Name: threat_intel_cache; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.threat_intel_cache (
    id integer NOT NULL,
    ioc text,
    reputation text,
    source text,
    last_checked timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.threat_intel_cache OWNER TO postgres;

--
-- Name: threat_intel_cache_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.threat_intel_cache_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.threat_intel_cache_id_seq OWNER TO postgres;

--
-- Name: threat_intel_cache_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.threat_intel_cache_id_seq OWNED BY public.threat_intel_cache.id;


--
-- Name: url_reputation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.url_reputation (
    id integer NOT NULL,
    domain character varying(255),
    reputation character varying(50)
);


ALTER TABLE public.url_reputation OWNER TO postgres;

--
-- Name: url_reputation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.url_reputation_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.url_reputation_id_seq OWNER TO postgres;

--
-- Name: url_reputation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.url_reputation_id_seq OWNED BY public.url_reputation.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username text,
    password text,
    role character varying(50) DEFAULT 'ANALYST'::character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: alerts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alerts ALTER COLUMN id SET DEFAULT nextval('public.alerts_id_seq'::regclass);


--
-- Name: analysis_results id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.analysis_results ALTER COLUMN id SET DEFAULT nextval('public.analysis_results_id_seq'::regclass);


--
-- Name: analyst_notes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.analyst_notes ALTER COLUMN id SET DEFAULT nextval('public.analyst_notes_id_seq'::regclass);


--
-- Name: campaigns id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.campaigns ALTER COLUMN id SET DEFAULT nextval('public.campaigns_id_seq'::regclass);


--
-- Name: case_notes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.case_notes ALTER COLUMN id SET DEFAULT nextval('public.case_notes_id_seq'::regclass);


--
-- Name: cases id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cases ALTER COLUMN id SET DEFAULT nextval('public.cases_id_seq'::regclass);


--
-- Name: docker_reports id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.docker_reports ALTER COLUMN id SET DEFAULT nextval('public.docker_reports_id_seq'::regclass);


--
-- Name: emails id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.emails ALTER COLUMN id SET DEFAULT nextval('public.emails_id_seq'::regclass);


--
-- Name: ioc_correlation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ioc_correlation ALTER COLUMN id SET DEFAULT nextval('public.ioc_correlation_id_seq'::regclass);


--
-- Name: mitre_events id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mitre_events ALTER COLUMN id SET DEFAULT nextval('public.mitre_events_id_seq'::regclass);


--
-- Name: sandbox_jobs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sandbox_jobs ALTER COLUMN id SET DEFAULT nextval('public.sandbox_jobs_id_seq'::regclass);


--
-- Name: sandbox_reports id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sandbox_reports ALTER COLUMN id SET DEFAULT nextval('public.sandbox_reports_id_seq'::regclass);


--
-- Name: threat_intel_cache id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.threat_intel_cache ALTER COLUMN id SET DEFAULT nextval('public.threat_intel_cache_id_seq'::regclass);


--
-- Name: url_reputation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url_reputation ALTER COLUMN id SET DEFAULT nextval('public.url_reputation_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: alerts alerts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alerts
    ADD CONSTRAINT alerts_pkey PRIMARY KEY (id);


--
-- Name: analysis_results analysis_results_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.analysis_results
    ADD CONSTRAINT analysis_results_pkey PRIMARY KEY (id);


--
-- Name: analyst_notes analyst_notes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.analyst_notes
    ADD CONSTRAINT analyst_notes_pkey PRIMARY KEY (id);


--
-- Name: campaigns campaigns_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_pkey PRIMARY KEY (id);


--
-- Name: case_notes case_notes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.case_notes
    ADD CONSTRAINT case_notes_pkey PRIMARY KEY (id);


--
-- Name: cases cases_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cases
    ADD CONSTRAINT cases_pkey PRIMARY KEY (id);


--
-- Name: docker_reports docker_reports_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.docker_reports
    ADD CONSTRAINT docker_reports_pkey PRIMARY KEY (id);


--
-- Name: emails emails_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.emails
    ADD CONSTRAINT emails_pkey PRIMARY KEY (id);


--
-- Name: ioc_correlation ioc_correlation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ioc_correlation
    ADD CONSTRAINT ioc_correlation_pkey PRIMARY KEY (id);


--
-- Name: mitre_events mitre_events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mitre_events
    ADD CONSTRAINT mitre_events_pkey PRIMARY KEY (id);


--
-- Name: sandbox_jobs sandbox_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sandbox_jobs
    ADD CONSTRAINT sandbox_jobs_pkey PRIMARY KEY (id);


--
-- Name: sandbox_reports sandbox_reports_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sandbox_reports
    ADD CONSTRAINT sandbox_reports_pkey PRIMARY KEY (id);


--
-- Name: threat_intel_cache threat_intel_cache_ioc_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.threat_intel_cache
    ADD CONSTRAINT threat_intel_cache_ioc_key UNIQUE (ioc);


--
-- Name: threat_intel_cache threat_intel_cache_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.threat_intel_cache
    ADD CONSTRAINT threat_intel_cache_pkey PRIMARY KEY (id);


--
-- Name: url_reputation url_reputation_domain_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url_reputation
    ADD CONSTRAINT url_reputation_domain_key UNIQUE (domain);


--
-- Name: url_reputation url_reputation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.url_reputation
    ADD CONSTRAINT url_reputation_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- PostgreSQL database dump complete
--

\unrestrict af0SbR3Bqk8fqeuO8A6cdVSCrr43cvY0NbcFuEZYM5tdsmR6stueFOSn78a30zO

