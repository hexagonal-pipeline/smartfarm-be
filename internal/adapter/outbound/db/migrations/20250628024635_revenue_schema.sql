-- 수익 기록 테이블 (정산용)
CREATE TABLE revenue_records (
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL REFERENCES user_stats(nickname),
    type VARCHAR(20) NOT NULL, -- 'rental', 'raid'
    amount INTEGER NOT NULL,
    source_id INTEGER, -- rental_id 또는 raid_participation_id
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 인덱스 생성 (성능용)
CREATE INDEX idx_revenue_records_nickname ON revenue_records(nickname); 