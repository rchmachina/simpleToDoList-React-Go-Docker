--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

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
-- Name: to_do_list; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA to_do_list;


ALTER SCHEMA to_do_list OWNER TO postgres;

--
-- Name: create_subtask(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.create_subtask(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO to_do_list.subtasks (task_id, subtask_name, is_complete)
    VALUES (
        (params->>'taskId')::INT,
        (params->>'subTaskName')::VARCHAR(100),
        (params->>'isComplete')::BOOlEAN
    );
END;
$$;


ALTER FUNCTION to_do_list.create_subtask(params jsonb) OWNER TO postgres;

--
-- Name: create_task(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.create_task(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO to_do_list.tasks (task_name, deadline, is_complete)
    VALUES (
        (params->>'taskName')::VARCHAR(100),
        (params->>'deadline')::TIMESTAMP,
        
        (params->>'isComplete')::BOOLEAN
    );
END;
$$;


ALTER FUNCTION to_do_list.create_task(params jsonb) OWNER TO postgres;

--
-- Name: delete_subtask(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.delete_subtask(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$
	begin
		 Delete from to_do_list.subtasks where subtask_id  = (params->>'subTaskId')::int;
	
	END;
$$;


ALTER FUNCTION to_do_list.delete_subtask(params jsonb) OWNER TO postgres;

--
-- Name: delete_task(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.delete_task(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$
	begin
		 Delete from to_do_list.tasks where task_id  = (params->>'taskId')::int;
	
	END;
$$;


ALTER FUNCTION to_do_list.delete_task(params jsonb) OWNER TO postgres;

--
-- Name: get_tasks_with_subtasks(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.get_tasks_with_subtasks(params jsonb) RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
DECLARE
    jsonb_result JSONB;
BEGIN
    SELECT jsonb_agg(jsonb_build_object(
                'taskId', t.task_id,
                'taskName', t.task_name,
                'deadline', TO_CHAR( t.deadline,'YYYY-MM-DD HH24:MI"'),
                'isComplete', t.is_complete,
                'subtasks', (
                    SELECT jsonb_agg(jsonb_build_object(
                                    'subtaskId', st.subtask_id,
                                    'subtaskName', st.subtask_name,
                                    'isComplete', st.is_complete
                                ))
                    FROM to_do_list.subtasks st
                    WHERE st.task_id = t.task_id
                )
            ))
    INTO jsonb_result
    FROM (
        SELECT t.*
        FROM to_do_list.tasks t
        ORDER BY t.created
    ) t;

    RETURN jsonb_result;
END;
$$;


ALTER FUNCTION to_do_list.get_tasks_with_subtasks(params jsonb) OWNER TO postgres;

--
-- Name: update_subtask(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.update_subtask(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$

BEGIN
    -- Extract values from the JSON parameter


    UPDATE to_do_list.subtasks
    set
    subtask_name  = (params ->> 'subTaskName')::varchar
    WHERE subtask_id  = (params->>'subTaskId')::int;
END;
$$;


ALTER FUNCTION to_do_list.update_subtask(params jsonb) OWNER TO postgres;

--
-- Name: update_subtask_is_complete(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.update_subtask_is_complete(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$

BEGIN
    -- Extract values from the JSON parameter


    UPDATE to_do_list.subtasks
	SET is_complete = (params ->> 'isComplete')::BOOLEAN

    WHERE subtask_id  = (params->>'subTaskId')::int;
END;
$$;


ALTER FUNCTION to_do_list.update_subtask_is_complete(params jsonb) OWNER TO postgres;

--
-- Name: update_task(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.update_task(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$

BEGIN
    -- Extract values from the JSON parameter


    UPDATE to_do_list.tasks
    set
    task_name = (params ->> 'taskName'),
    deadline = (params ->> 'deadline')::TIMESTAMP,
    is_complete =(params ->> 'isComplete')::boolean
    WHERE task_id = (params->>'taskId')::int;
END;
$$;


ALTER FUNCTION to_do_list.update_task(params jsonb) OWNER TO postgres;

--
-- Name: update_task_is_complete(jsonb); Type: FUNCTION; Schema: to_do_list; Owner: postgres
--

CREATE FUNCTION to_do_list.update_task_is_complete(params jsonb) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Extract values from the JSON parameter
    UPDATE to_do_list.tasks
    SET is_complete = (params ->> 'isComplete')::BOOLEAN
    WHERE task_id = (params ->> 'taskId')::INT;

    -- Update subtask status based on the task's completion status
    IF (params ->> 'isComplete')::BOOLEAN = TRUE THEN
        UPDATE to_do_list.subtasks
        SET is_complete = TRUE
        WHERE task_id = (params ->> 'taskId')::INT;
    END IF;

END;
$$;


ALTER FUNCTION to_do_list.update_task_is_complete(params jsonb) OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: subtasks; Type: TABLE; Schema: to_do_list; Owner: postgres
--

CREATE TABLE to_do_list.subtasks (
    subtask_id integer NOT NULL,
    task_id integer,
    subtask_name character varying(100),
    is_complete boolean DEFAULT false,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone
);


ALTER TABLE to_do_list.subtasks OWNER TO postgres;

--
-- Name: subtasks_subtask_id_seq; Type: SEQUENCE; Schema: to_do_list; Owner: postgres
--

CREATE SEQUENCE to_do_list.subtasks_subtask_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE to_do_list.subtasks_subtask_id_seq OWNER TO postgres;

--
-- Name: subtasks_subtask_id_seq; Type: SEQUENCE OWNED BY; Schema: to_do_list; Owner: postgres
--

ALTER SEQUENCE to_do_list.subtasks_subtask_id_seq OWNED BY to_do_list.subtasks.subtask_id;


--
-- Name: tasks; Type: TABLE; Schema: to_do_list; Owner: postgres
--

CREATE TABLE to_do_list.tasks (
    task_id integer NOT NULL,
    task_name character varying(100),
    deadline timestamp without time zone,
    is_complete boolean DEFAULT false,
    created timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated timestamp without time zone
);


ALTER TABLE to_do_list.tasks OWNER TO postgres;

--
-- Name: tasks_task_id_seq; Type: SEQUENCE; Schema: to_do_list; Owner: postgres
--

CREATE SEQUENCE to_do_list.tasks_task_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE to_do_list.tasks_task_id_seq OWNER TO postgres;

--
-- Name: tasks_task_id_seq; Type: SEQUENCE OWNED BY; Schema: to_do_list; Owner: postgres
--

ALTER SEQUENCE to_do_list.tasks_task_id_seq OWNED BY to_do_list.tasks.task_id;


--
-- Name: subtasks subtask_id; Type: DEFAULT; Schema: to_do_list; Owner: postgres
--

ALTER TABLE ONLY to_do_list.subtasks ALTER COLUMN subtask_id SET DEFAULT nextval('to_do_list.subtasks_subtask_id_seq'::regclass);


--
-- Name: tasks task_id; Type: DEFAULT; Schema: to_do_list; Owner: postgres
--

ALTER TABLE ONLY to_do_list.tasks ALTER COLUMN task_id SET DEFAULT nextval('to_do_list.tasks_task_id_seq'::regclass);


--
-- Data for Name: subtasks; Type: TABLE DATA; Schema: to_do_list; Owner: postgres
--

COPY to_do_list.subtasks (subtask_id, task_id, subtask_name, is_complete, created, updated) FROM stdin;
1	1	Subtask 1.6	f	2024-02-03 23:47:58.955307	\N
2	1	Subtask 1.5	f	2024-02-03 23:47:58.955307	\N
3	1	Subtask 1.3	f	2024-02-03 23:47:58.955307	\N
4	1	Subtask 1.4	f	2024-02-03 23:47:58.955307	\N
5	2	Subtask 2.1	f	2024-02-03 23:47:58.955307	\N
6	3	Subtask 3.1	t	2024-02-03 23:47:58.955307	\N
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: to_do_list; Owner: postgres
--

COPY to_do_list.tasks (task_id, task_name, deadline, is_complete, created, updated) FROM stdin;
1	Task 1	2024-02-10 15:30:00.22	f	2024-02-03 23:47:58.955307	\N
2	Task 2	2024-02-15 15:30:00.22	f	2024-02-03 23:47:58.955307	\N
3	Task 3	2024-02-20 15:30:00.22	t	2024-02-03 23:47:58.955307	\N
\.


--
-- Name: subtasks_subtask_id_seq; Type: SEQUENCE SET; Schema: to_do_list; Owner: postgres
--

SELECT pg_catalog.setval('to_do_list.subtasks_subtask_id_seq', 6, true);


--
-- Name: tasks_task_id_seq; Type: SEQUENCE SET; Schema: to_do_list; Owner: postgres
--

SELECT pg_catalog.setval('to_do_list.tasks_task_id_seq', 3, true);


--
-- Name: subtasks subtasks_pkey; Type: CONSTRAINT; Schema: to_do_list; Owner: postgres
--

ALTER TABLE ONLY to_do_list.subtasks
    ADD CONSTRAINT subtasks_pkey PRIMARY KEY (subtask_id);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: to_do_list; Owner: postgres
--

ALTER TABLE ONLY to_do_list.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (task_id);


--
-- Name: subtasks subtasks_task_id_fkey; Type: FK CONSTRAINT; Schema: to_do_list; Owner: postgres
--

ALTER TABLE ONLY to_do_list.subtasks
    ADD CONSTRAINT subtasks_task_id_fkey FOREIGN KEY (task_id) REFERENCES to_do_list.tasks(task_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

