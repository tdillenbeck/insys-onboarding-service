--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: insys_onboarding; Type: SCHEMA; Schema: -; Owner: zach.toolsongetweave.com
--

CREATE SCHEMA insys_onboarding;


ALTER SCHEMA insys_onboarding OWNER TO "zach.toolsongetweave.com";

SET search_path = insys_onboarding, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: goose_db_version; Type: TABLE; Schema: insys_onboarding; Owner: postgres
--

CREATE TABLE goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE goose_db_version OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: insys_onboarding; Owner: postgres
--

CREATE SEQUENCE goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE goose_db_version_id_seq OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: insys_onboarding; Owner: postgres
--

ALTER SEQUENCE goose_db_version_id_seq OWNED BY goose_db_version.id;


--
-- Name: onboarders; Type: TABLE; Schema: insys_onboarding; Owner: postgres
--

CREATE TABLE onboarders (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    schedule_custimization_link text,
    schedule_porting_link text,
    schedule_network_link text,
    schedule_software_install_link text,
    schedule_phone_install_link text,
    schedule_software_training_link text,
    schedule_phone_training_link text,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE onboarders OWNER TO postgres;

--
-- Name: onboarders_location; Type: TABLE; Schema: insys_onboarding; Owner: postgres
--

CREATE TABLE onboarders_location (
    id uuid NOT NULL,
    onboarder_id uuid NOT NULL,
    location_id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE onboarders_location OWNER TO postgres;

--
-- Name: onboarding_categories; Type: TABLE; Schema: insys_onboarding; Owner: postgres
--

CREATE TABLE onboarding_categories (
    id uuid NOT NULL,
    display_text text NOT NULL,
    display_order integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE onboarding_categories OWNER TO postgres;

--
-- Name: onboarding_task_instances; Type: TABLE; Schema: insys_onboarding; Owner: postgres
--

CREATE TABLE onboarding_task_instances (
    id uuid NOT NULL,
    location_id uuid NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    display_order integer NOT NULL,
    status integer NOT NULL,
    status_updated_at timestamp without time zone NOT NULL,
    status_updated_by text,
    completed_at timestamp without time zone,
    completed_by text,
    verified_at timestamp without time zone,
    verified_by text,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    onboarding_category_id uuid NOT NULL,
    onboarding_task_id uuid NOT NULL,
    button_content text,
    button_external_url text,
    explanation text
);


ALTER TABLE onboarding_task_instances OWNER TO postgres;

--
-- Name: onboarding_tasks; Type: TABLE; Schema: insys_onboarding; Owner: postgres
--

CREATE TABLE onboarding_tasks (
    id uuid NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    display_order integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    onboarding_category_id uuid NOT NULL,
    button_content text,
    button_external_url text
);


ALTER TABLE onboarding_tasks OWNER TO postgres;

--
-- Name: goose_db_version id; Type: DEFAULT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY goose_db_version ALTER COLUMN id SET DEFAULT nextval('goose_db_version_id_seq'::regclass);


--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: onboarders_location onboarders_location_pkey; Type: CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarders_location
    ADD CONSTRAINT onboarders_location_pkey PRIMARY KEY (id);


--
-- Name: onboarders onboarders_pkey; Type: CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarders
    ADD CONSTRAINT onboarders_pkey PRIMARY KEY (id);


--
-- Name: onboarding_categories onboarding_categories_pkey; Type: CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarding_categories
    ADD CONSTRAINT onboarding_categories_pkey PRIMARY KEY (id);


--
-- Name: onboarding_task_instances onboarding_task_instances_pkey; Type: CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarding_task_instances
    ADD CONSTRAINT onboarding_task_instances_pkey PRIMARY KEY (id);


--
-- Name: onboarding_tasks onboarding_tasks_pkey; Type: CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarding_tasks
    ADD CONSTRAINT onboarding_tasks_pkey PRIMARY KEY (id);


--
-- Name: index_onboarders_location_on_location_id; Type: INDEX; Schema: insys_onboarding; Owner: postgres
--

CREATE UNIQUE INDEX index_onboarders_location_on_location_id ON onboarders_location USING btree (location_id);


--
-- Name: index_onboarders_on_user_id; Type: INDEX; Schema: insys_onboarding; Owner: postgres
--

CREATE UNIQUE INDEX index_onboarders_on_user_id ON onboarders USING btree (user_id);


--
-- Name: onboarders_location onboarders_location_onboarder_id_fkey; Type: FK CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarders_location
    ADD CONSTRAINT onboarders_location_onboarder_id_fkey FOREIGN KEY (onboarder_id) REFERENCES onboarders(id);


--
-- Name: onboarding_task_instances onboarding_task_instances_onboarding_category_id_fkey; Type: FK CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarding_task_instances
    ADD CONSTRAINT onboarding_task_instances_onboarding_category_id_fkey FOREIGN KEY (onboarding_category_id) REFERENCES onboarding_categories(id);


--
-- Name: onboarding_task_instances onboarding_task_instances_onboarding_task_id_fkey; Type: FK CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarding_task_instances
    ADD CONSTRAINT onboarding_task_instances_onboarding_task_id_fkey FOREIGN KEY (onboarding_task_id) REFERENCES onboarding_tasks(id);


--
-- Name: onboarding_tasks onboarding_tasks_onboarding_category_id_fkey; Type: FK CONSTRAINT; Schema: insys_onboarding; Owner: postgres
--

ALTER TABLE ONLY onboarding_tasks
    ADD CONSTRAINT onboarding_tasks_onboarding_category_id_fkey FOREIGN KEY (onboarding_category_id) REFERENCES onboarding_categories(id);


--
-- PostgreSQL database dump complete
--

