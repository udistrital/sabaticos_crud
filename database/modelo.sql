CREATE SCHEMA IF NOT EXISTS sabatico;

-- =========================
-- TABLAS MAESTRAS
-- =========================

CREATE TABLE sabatico.estado_solicitud (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(20) NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    descripcion VARCHAR(250),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL
);

CREATE TABLE sabatico.tipo_solicitud (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(20) NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    descripcion VARCHAR(250),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL
);

CREATE TABLE sabatico.estado_soporte_solicitud (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(20) NOT NULL,
    nombre_estado VARCHAR(50) NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL
);

CREATE TABLE sabatico.estado_sabatico (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(20) NOT NULL,
    nombre_estado VARCHAR(50) NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL
);

CREATE TABLE sabatico.estado_soporte_sabatico (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(20) NOT NULL,
    nombre_estado VARCHAR(50) NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL
);

-- =========================
-- SABATICO
-- =========================

CREATE TABLE sabatico.sabatico (
    id SERIAL PRIMARY KEY,
    tercero_id INTEGER NOT NULL,
    observaciones VARCHAR(500),
    fecha_inicio TIMESTAMP NOT NULL,
    fecha_fin TIMESTAMP NOT NULL,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL
);

CREATE TABLE sabatico.historial_estado_sabatico (
    id SERIAL PRIMARY KEY,
    tercero_id INTEGER NOT NULL,
    justificacion VARCHAR(250),
    estado_sabatico_id INTEGER NOT NULL,
    sabatico_id INTEGER NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,

    CONSTRAINT fk_historial_estado_sabatico_estado
        FOREIGN KEY (estado_sabatico_id)
        REFERENCES sabatico.estado_sabatico(id),

    CONSTRAINT fk_historial_estado_sabatico_sabatico
        FOREIGN KEY (sabatico_id)
        REFERENCES sabatico.sabatico(id)
);

CREATE TABLE sabatico.soporte_sabatico (
    id SERIAL PRIMARY KEY,
    documento_id INTEGER NOT NULL,
    sabatico_id INTEGER,
    estado_soporte_sabatico_id INTEGER,
    rol_usuario VARCHAR(100),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,

    CONSTRAINT fk_soporte_sabatico_sabatico
        FOREIGN KEY (sabatico_id)
        REFERENCES sabatico.sabatico(id),

    CONSTRAINT fk_soporte_sabatico_estado
        FOREIGN KEY (estado_soporte_sabatico_id)
        REFERENCES sabatico.estado_soporte_sabatico(id)
);

-- =========================
-- SOLICITUD
-- =========================

CREATE TABLE sabatico.solicitud (
    id SERIAL PRIMARY KEY,
    tercero_id INTEGER NOT NULL,
    tipo_solicitud_id INTEGER,
    sabatico_id INTEGER,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,

    CONSTRAINT fk_solicitud_tipo
        FOREIGN KEY (tipo_solicitud_id)
        REFERENCES sabatico.tipo_solicitud(id),

    CONSTRAINT fk_solicitud_sabatico
        FOREIGN KEY (sabatico_id)
        REFERENCES sabatico.sabatico(id)
);

CREATE TABLE sabatico.historial_solicitud (
    id SERIAL PRIMARY KEY,
    tercero_id INTEGER NOT NULL,
    justificacion VARCHAR(250),
    estado_solicitud_id INTEGER,
    solicitud_id INTEGER,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,

    CONSTRAINT fk_historial_solicitud_estado
        FOREIGN KEY (estado_solicitud_id)
        REFERENCES sabatico.estado_solicitud(id),

    CONSTRAINT fk_historial_solicitud_solicitud
        FOREIGN KEY (solicitud_id)
        REFERENCES sabatico.solicitud(id)
);

CREATE TABLE sabatico.formulario_solicitud (
    id SERIAL PRIMARY KEY,
    contenido JSONB NOT NULL,
    solicitud_id INTEGER,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,

    CONSTRAINT fk_formulario_solicitud
        FOREIGN KEY (solicitud_id)
        REFERENCES sabatico.solicitud(id)
);

CREATE TABLE sabatico.soporte_solicitud (
    id SERIAL PRIMARY KEY,
    documento_id INTEGER NOT NULL,
    tercero_id INTEGER NOT NULL,
    solicitud_id INTEGER,
    estado_soporte_solicitud_id INTEGER,
    rol_usuario VARCHAR(100),
    tipo_documento_id NUMERIC,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,

    CONSTRAINT fk_soporte_solicitud_solicitud
        FOREIGN KEY (solicitud_id)
        REFERENCES sabatico.solicitud(id),

    CONSTRAINT fk_soporte_solicitud_estado
        FOREIGN KEY (estado_soporte_solicitud_id)
        REFERENCES sabatico.estado_soporte_solicitud(id)
);