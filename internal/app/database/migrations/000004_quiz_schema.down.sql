-- Migration: 000004_quiz_schema (rollback)
DROP TABLE IF EXISTS quiz_game_answers;
DROP TABLE IF EXISTS quiz_game_questions;
DROP TABLE IF EXISTS quiz_games;
DROP TABLE IF EXISTS quiz_questions;
