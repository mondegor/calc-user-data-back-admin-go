-- --------------------------------------------------------------------------------------------------

CREATE TABLE printshop_dictionaries.material_types (
    type_id int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY CONSTRAINT pk_material_types PRIMARY KEY,
    tag_version int4 NOT NULL DEFAULT 1 CHECK(tag_version > 0),
    parent_id int8 NULL,
    type_caption character varying(64) NOT NULL,
    type_status int2 NOT NULL, -- 1=DRAFT, 2=ENABLED, 3=DISABLED
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp with time zone NULL,
    FOREIGN KEY (parent_id) REFERENCES printshop_dictionaries.material_types(type_id)
);

-- --------------------------------------------------------------------------------------------------

INSERT INTO printshop_dictionaries.material_types (type_id, tag_version, parent_id, type_caption, type_status, created_at, updated_at, deleted_at)
VALUES
    (1, 1, NULL, 'Бумага', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL),
    (2, 1, NULL, 'Ламинат', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL),
    (3, 1, 1, 'меловка матовая', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL),
    (4, 1, 1, 'меловка глянцевая', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL),
    (5, 1, 1, 'офсетная', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL),
    (6, 1, 1, 'картон', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL),
    (7, 1, 2, 'матовый', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL),
    (8, 1, 2, 'глянцевый', 2/*ENABLED*/, '2023-07-30 12:30:52.651613', '2023-07-30 12:30:52.651613', NULL);

ALTER SEQUENCE printshop_dictionaries.material_types_type_id_seq RESTART WITH 9;