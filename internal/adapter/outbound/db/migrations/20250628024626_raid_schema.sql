-- migrate:up
-- 유통 레이드 테이블 (핵심 기능)
CREATE TABLE raids (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    crop_type VARCHAR(50) NOT NULL,
    target_quantity INTEGER NOT NULL, -- 목표 수량 (kg)
    min_participation INTEGER NOT NULL, -- 최소 참여 수량
    max_participation INTEGER NOT NULL, -- 최대 참여 수량
    price_per_kg INTEGER NOT NULL, -- kg당 가격
    deadline TIMESTAMP NOT NULL,
    status VARCHAR(20) DEFAULT 'open', -- open, closed, completed
    creator_nickname VARCHAR(50) NOT NULL REFERENCES user_stats(nickname),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 레이드 참여 테이블
CREATE TABLE raid_participations (
    id SERIAL PRIMARY KEY,
    raid_id INTEGER NOT NULL REFERENCES raids(id),
    participant_nickname VARCHAR(50) NOT NULL REFERENCES user_stats(nickname),
    quantity INTEGER NOT NULL, -- 참여 수량 (kg)
    expected_revenue INTEGER NOT NULL, -- 예상 수익
    status VARCHAR(20) DEFAULT 'pending', -- pending, confirmed, paid
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(raid_id, participant_nickname) -- 한 레이드당 한 번만 참여
);

-- 인덱스 생성 (성능용)
CREATE INDEX idx_raids_status ON raids(status);
CREATE INDEX idx_raid_participations_raid_id ON raid_participations(raid_id);
CREATE INDEX idx_raid_participations_nickname ON raid_participations(participant_nickname);

-- migrate:down
DROP TABLE IF EXISTS raid_participations;
DROP TABLE IF EXISTS raids; 