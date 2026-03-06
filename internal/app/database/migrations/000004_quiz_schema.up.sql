-- Migration: 000004_quiz_schema
-- Purpose: Quiz game tables - questions, games, game questions, answers

-- ============================================================================
-- Table: quiz_questions
-- Purpose: Pool of questions managed by super admin
-- ============================================================================
CREATE TABLE quiz_questions (
    id              UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    body            TEXT        NOT NULL CHECK (length(body) >= 10 AND length(body) <= 500),
    correct_answers JSONB       NOT NULL DEFAULT '[]',
    published       BOOLEAN     NOT NULL DEFAULT false,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ
);

CREATE INDEX idx_quiz_questions_published ON quiz_questions (published);

-- ============================================================================
-- Table: quiz_games
-- Purpose: Pair game sessions between two registered users
-- ============================================================================
CREATE TABLE quiz_games (
    id                          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    first_player_id             UUID        NOT NULL REFERENCES users(id),
    second_player_id            UUID        REFERENCES users(id),
    first_player_score          INTEGER     NOT NULL DEFAULT 0,
    second_player_score         INTEGER     NOT NULL DEFAULT 0,
    first_player_finished_at    TIMESTAMPTZ,
    second_player_finished_at   TIMESTAMPTZ,
    status                      VARCHAR(50) NOT NULL DEFAULT 'PendingSecondPlayer'
                                    CHECK (status IN ('PendingSecondPlayer', 'Active', 'Finished')),
    created_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    started_at                  TIMESTAMPTZ,
    finished_at                 TIMESTAMPTZ
);

CREATE INDEX idx_quiz_games_first_player  ON quiz_games (first_player_id);
CREATE INDEX idx_quiz_games_second_player ON quiz_games (second_player_id);
CREATE INDEX idx_quiz_games_status        ON quiz_games (status);

-- ============================================================================
-- Table: quiz_game_questions
-- Purpose: 5 questions assigned to a game when second player joins
-- ============================================================================
CREATE TABLE quiz_game_questions (
    id          UUID    PRIMARY KEY DEFAULT gen_random_uuid(),
    game_id     UUID    NOT NULL REFERENCES quiz_games(id) ON DELETE CASCADE,
    question_id UUID    NOT NULL REFERENCES quiz_questions(id) ON DELETE CASCADE,
    order_index INTEGER NOT NULL CHECK (order_index BETWEEN 0 AND 4),
    UNIQUE (game_id, order_index)
);

CREATE INDEX idx_quiz_game_questions_game ON quiz_game_questions (game_id);

-- ============================================================================
-- Table: quiz_game_answers
-- Purpose: Answers submitted by players (one per question per player)
-- ============================================================================
CREATE TABLE quiz_game_answers (
    id          UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    game_id     UUID        NOT NULL REFERENCES quiz_games(id) ON DELETE CASCADE,
    player_id   UUID        NOT NULL REFERENCES users(id),
    question_id UUID        NOT NULL REFERENCES quiz_questions(id) ON DELETE CASCADE,
    answer      TEXT        NOT NULL,
    is_correct  BOOLEAN     NOT NULL,
    answered_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (game_id, player_id, question_id)
);

CREATE INDEX idx_quiz_game_answers_game_player ON quiz_game_answers (game_id, player_id);
