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
-- Name: commission_works; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.commission_works (
    id integer NOT NULL,
    requester_nickname character varying(50) NOT NULL,
    plot_id integer NOT NULL,
    task_type character varying(50) NOT NULL,
    task_description text,
    status character varying(20) DEFAULT 'requested'::character varying NOT NULL,
    credit_cost integer NOT NULL,
    requested_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    completed_at timestamp with time zone
);


--
-- Name: commission_works_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.commission_works_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: commission_works_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.commission_works_id_seq OWNED BY public.commission_works.id;


--
-- Name: credit_transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.credit_transactions (
    id integer NOT NULL,
    nickname character varying(50) NOT NULL,
    transaction_type character varying(50) NOT NULL,
    amount integer NOT NULL,
    related_id integer,
    description text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: credit_transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.credit_transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: credit_transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.credit_transactions_id_seq OWNED BY public.credit_transactions.id;


--
-- Name: farm_plots; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.farm_plots (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    location character varying(200),
    size_sqm integer NOT NULL,
    monthly_rent integer NOT NULL,
    crop_type character varying(50),
    status character varying(20) DEFAULT 'available'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    persona_prompt text
);


--
-- Name: farm_plots_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.farm_plots_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: farm_plots_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.farm_plots_id_seq OWNED BY public.farm_plots.id;


--
-- Name: plant_cards; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.plant_cards (
    id integer NOT NULL,
    farm_plot_id integer NOT NULL,
    persona text NOT NULL,
    image_url character varying(255),
    video_url character varying(255),
    event_message text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: plant_cards_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.plant_cards_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: plant_cards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.plant_cards_id_seq OWNED BY public.plant_cards.id;


--
-- Name: raid_participations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.raid_participations (
    id integer NOT NULL,
    raid_id integer NOT NULL,
    participant_nickname character varying(50) NOT NULL,
    quantity integer NOT NULL,
    expected_revenue integer NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: raid_participations_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.raid_participations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: raid_participations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.raid_participations_id_seq OWNED BY public.raid_participations.id;


--
-- Name: raids; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.raids (
    id integer NOT NULL,
    title character varying(200) NOT NULL,
    description text,
    crop_type character varying(50) NOT NULL,
    target_quantity integer NOT NULL,
    min_participation integer NOT NULL,
    max_participation integer NOT NULL,
    price_per_kg integer NOT NULL,
    deadline timestamp without time zone NOT NULL,
    status character varying(20) DEFAULT 'open'::character varying,
    creator_nickname character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: raids_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.raids_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: raids_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.raids_id_seq OWNED BY public.raids.id;


--
-- Name: rentals; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.rentals (
    id integer NOT NULL,
    renter_nickname character varying(50) NOT NULL,
    plot_id integer NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    monthly_rent integer NOT NULL,
    status character varying(20) DEFAULT 'active'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: rentals_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.rentals_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: rentals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.rentals_id_seq OWNED BY public.rentals.id;


--
-- Name: revenue_records; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.revenue_records (
    id integer NOT NULL,
    nickname character varying(50) NOT NULL,
    type character varying(20) NOT NULL,
    amount integer NOT NULL,
    source_id integer,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: revenue_records_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.revenue_records_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: revenue_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.revenue_records_id_seq OWNED BY public.revenue_records.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying NOT NULL
);


--
-- Name: user_stats; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_stats (
    nickname character varying(50) NOT NULL,
    level integer DEFAULT 1,
    experience integer DEFAULT 0,
    credit integer DEFAULT 0,
    total_revenue integer DEFAULT 0,
    successful_raids integer DEFAULT 0,
    plots_rented integer DEFAULT 0,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


--
-- Name: commission_works id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.commission_works ALTER COLUMN id SET DEFAULT nextval('public.commission_works_id_seq'::regclass);


--
-- Name: credit_transactions id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.credit_transactions ALTER COLUMN id SET DEFAULT nextval('public.credit_transactions_id_seq'::regclass);


--
-- Name: farm_plots id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.farm_plots ALTER COLUMN id SET DEFAULT nextval('public.farm_plots_id_seq'::regclass);


--
-- Name: plant_cards id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.plant_cards ALTER COLUMN id SET DEFAULT nextval('public.plant_cards_id_seq'::regclass);


--
-- Name: raid_participations id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raid_participations ALTER COLUMN id SET DEFAULT nextval('public.raid_participations_id_seq'::regclass);


--
-- Name: raids id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raids ALTER COLUMN id SET DEFAULT nextval('public.raids_id_seq'::regclass);


--
-- Name: rentals id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rentals ALTER COLUMN id SET DEFAULT nextval('public.rentals_id_seq'::regclass);


--
-- Name: revenue_records id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.revenue_records ALTER COLUMN id SET DEFAULT nextval('public.revenue_records_id_seq'::regclass);


--
-- Name: commission_works commission_works_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.commission_works
    ADD CONSTRAINT commission_works_pkey PRIMARY KEY (id);


--
-- Name: credit_transactions credit_transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.credit_transactions
    ADD CONSTRAINT credit_transactions_pkey PRIMARY KEY (id);


--
-- Name: farm_plots farm_plots_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.farm_plots
    ADD CONSTRAINT farm_plots_pkey PRIMARY KEY (id);


--
-- Name: plant_cards plant_cards_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.plant_cards
    ADD CONSTRAINT plant_cards_pkey PRIMARY KEY (id);


--
-- Name: raid_participations raid_participations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raid_participations
    ADD CONSTRAINT raid_participations_pkey PRIMARY KEY (id);


--
-- Name: raid_participations raid_participations_raid_id_participant_nickname_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raid_participations
    ADD CONSTRAINT raid_participations_raid_id_participant_nickname_key UNIQUE (raid_id, participant_nickname);


--
-- Name: raids raids_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raids
    ADD CONSTRAINT raids_pkey PRIMARY KEY (id);


--
-- Name: rentals rentals_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rentals
    ADD CONSTRAINT rentals_pkey PRIMARY KEY (id);


--
-- Name: revenue_records revenue_records_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.revenue_records
    ADD CONSTRAINT revenue_records_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: user_stats user_stats_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_stats
    ADD CONSTRAINT user_stats_pkey PRIMARY KEY (nickname);


--
-- Name: idx_commission_works_plot_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_commission_works_plot_id ON public.commission_works USING btree (plot_id);


--
-- Name: idx_commission_works_requester; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_commission_works_requester ON public.commission_works USING btree (requester_nickname);


--
-- Name: idx_credit_transactions_nickname; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_credit_transactions_nickname ON public.credit_transactions USING btree (nickname);


--
-- Name: idx_plant_cards_created_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_plant_cards_created_at ON public.plant_cards USING btree (created_at DESC);


--
-- Name: idx_plant_cards_farm_plot_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_plant_cards_farm_plot_id ON public.plant_cards USING btree (farm_plot_id);


--
-- Name: idx_raid_participations_nickname; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_raid_participations_nickname ON public.raid_participations USING btree (participant_nickname);


--
-- Name: idx_raid_participations_raid_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_raid_participations_raid_id ON public.raid_participations USING btree (raid_id);


--
-- Name: idx_raids_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_raids_status ON public.raids USING btree (status);


--
-- Name: idx_rentals_nickname; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_rentals_nickname ON public.rentals USING btree (renter_nickname);


--
-- Name: idx_rentals_plot_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_rentals_plot_id ON public.rentals USING btree (plot_id);


--
-- Name: idx_revenue_records_nickname; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_revenue_records_nickname ON public.revenue_records USING btree (nickname);


--
-- Name: commission_works commission_works_plot_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.commission_works
    ADD CONSTRAINT commission_works_plot_id_fkey FOREIGN KEY (plot_id) REFERENCES public.farm_plots(id);


--
-- Name: commission_works commission_works_requester_nickname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.commission_works
    ADD CONSTRAINT commission_works_requester_nickname_fkey FOREIGN KEY (requester_nickname) REFERENCES public.user_stats(nickname);


--
-- Name: credit_transactions credit_transactions_nickname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.credit_transactions
    ADD CONSTRAINT credit_transactions_nickname_fkey FOREIGN KEY (nickname) REFERENCES public.user_stats(nickname);


--
-- Name: plant_cards plant_cards_farm_plot_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.plant_cards
    ADD CONSTRAINT plant_cards_farm_plot_id_fkey FOREIGN KEY (farm_plot_id) REFERENCES public.farm_plots(id) ON DELETE CASCADE;


--
-- Name: raid_participations raid_participations_participant_nickname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raid_participations
    ADD CONSTRAINT raid_participations_participant_nickname_fkey FOREIGN KEY (participant_nickname) REFERENCES public.user_stats(nickname);


--
-- Name: raid_participations raid_participations_raid_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raid_participations
    ADD CONSTRAINT raid_participations_raid_id_fkey FOREIGN KEY (raid_id) REFERENCES public.raids(id);


--
-- Name: raids raids_creator_nickname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.raids
    ADD CONSTRAINT raids_creator_nickname_fkey FOREIGN KEY (creator_nickname) REFERENCES public.user_stats(nickname);


--
-- Name: rentals rentals_plot_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rentals
    ADD CONSTRAINT rentals_plot_id_fkey FOREIGN KEY (plot_id) REFERENCES public.farm_plots(id);


--
-- Name: rentals rentals_renter_nickname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.rentals
    ADD CONSTRAINT rentals_renter_nickname_fkey FOREIGN KEY (renter_nickname) REFERENCES public.user_stats(nickname);


--
-- Name: revenue_records revenue_records_nickname_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.revenue_records
    ADD CONSTRAINT revenue_records_nickname_fkey FOREIGN KEY (nickname) REFERENCES public.user_stats(nickname);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20250628024600'),
    ('20250628024618'),
    ('20250628024626'),
    ('20250628024635'),
    ('20250628024642'),
    ('20250628024650'),
    ('20250628045344'),
    ('20250628045345'),
    ('20250628050000'),
    ('20250628134500');
