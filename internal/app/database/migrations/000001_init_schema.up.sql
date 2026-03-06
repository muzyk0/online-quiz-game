-- Migration: 000001_init_schema
-- Purpose: Initialize base database schema for quiz game application

-- Enable pgcrypto extension for UUID generation
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- ============================================================================
-- Table: users
-- ============================================================================
CREATE TABLE users (
    id            UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    login         VARCHAR(10)  UNIQUE,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT,
    first_name    TEXT,
    last_name     TEXT,
    avatar_url    TEXT,
    is_verified   BOOLEAN      NOT NULL DEFAULT false,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
