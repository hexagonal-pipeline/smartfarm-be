-- 농장 구획 테이블
CREATE TABLE farm_plots (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    location VARCHAR(200),
    size_sqm INTEGER NOT NULL, -- 평방미터
    monthly_rent INTEGER NOT NULL, -- 월 임대료
    crop_type VARCHAR(50), -- 재배 작물 종류
    status VARCHAR(20) DEFAULT 'available', -- available, rented, maintenance
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 임대 정보 테이블
CREATE TABLE rentals (
    id SERIAL PRIMARY KEY,
    renter_nickname VARCHAR(50) NOT NULL,
    plot_id INTEGER REFERENCES farm_plots(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    monthly_rent INTEGER NOT NULL,
    status VARCHAR(20) DEFAULT 'active', -- active, completed, cancelled
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

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
    creator_nickname VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 레이드 참여 테이블
CREATE TABLE raid_participations (
    id SERIAL PRIMARY KEY,
    raid_id INTEGER REFERENCES raids(id),
    participant_nickname VARCHAR(50) NOT NULL,
    quantity INTEGER NOT NULL, -- 참여 수량 (kg)
    expected_revenue INTEGER NOT NULL, -- 예상 수익
    status VARCHAR(20) DEFAULT 'pending', -- pending, confirmed, paid
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(raid_id, participant_nickname) -- 한 레이드당 한 번만 참여
);

-- 게이미피케이션 사용자 통계
CREATE TABLE user_stats (
    nickname VARCHAR(50) PRIMARY KEY,
    level INTEGER DEFAULT 1,
    experience INTEGER DEFAULT 0,
    total_revenue INTEGER DEFAULT 0, -- 총 수익
    successful_raids INTEGER DEFAULT 0, -- 성공한 레이드 횟수
    plots_rented INTEGER DEFAULT 0, -- 임대한 구획 수
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 수익 기록 테이블 (정산용)
CREATE TABLE revenue_records (
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL,
    type VARCHAR(20) NOT NULL, -- 'rental', 'raid'
    amount INTEGER NOT NULL,
    source_id INTEGER, -- rental_id 또는 raid_participation_id
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 인덱스 생성 (성능용)
CREATE INDEX idx_rentals_nickname ON rentals(renter_nickname);
CREATE INDEX idx_rentals_plot_id ON rentals(plot_id);
CREATE INDEX idx_raid_participations_raid_id ON raid_participations(raid_id);
CREATE INDEX idx_raid_participations_nickname ON raid_participations(participant_nickname);
CREATE INDEX idx_raids_status ON raids(status);
CREATE INDEX idx_revenue_records_nickname ON revenue_records(nickname);

-- 더미 데이터 삽입 (데모용)
INSERT INTO farm_plots (name, location, size_sqm, monthly_rent, crop_type, status) VALUES 
('A구역-01', '경기도 파주시', 100, 150000, '상추', 'available'),
('A구역-02', '경기도 파주시', 150, 200000, '토마토', 'rented'),
('B구역-01', '충남 천안시', 200, 250000, '딸기', 'available'),
('B구역-02', '충남 천안시', 120, 180000, '오이', 'available');

INSERT INTO raids (title, description, crop_type, target_quantity, min_participation, max_participation, price_per_kg, deadline, creator_nickname) VALUES 
('대형마트 상추 납품', '대형마트에서 신선한 상추 500kg을 주문했습니다', '상추', 500, 10, 100, 3000, '2024-02-15 23:59:59', '농부김씨'),
('학교급식 토마토 공급', '인근 학교 급식용 토마토 대량 주문', '토마토', 300, 20, 80, 4000, '2024-02-20 23:59:59', '청년이');

INSERT INTO user_stats (nickname, level, experience, total_revenue, successful_raids, plots_rented) VALUES 
('농부김씨', 3, 250, 500000, 5, 2),
('청년이', 2, 150, 200000, 3, 1),
('스마트농부', 1, 50, 50000, 1, 1),
('레이드마스터', 4, 400, 800000, 8, 3); 