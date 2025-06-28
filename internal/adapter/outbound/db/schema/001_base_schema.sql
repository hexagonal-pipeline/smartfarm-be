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

-- 게이미피케ATION 사용자 통계
CREATE TABLE user_stats (
    nickname VARCHAR(50) PRIMARY KEY,
    level INTEGER DEFAULT 1,
    experience INTEGER DEFAULT 0,
    credit INTEGER DEFAULT 0, -- 크레딧 잔액
    total_revenue INTEGER DEFAULT 0, -- 총 수익
    successful_raids INTEGER DEFAULT 0, -- 성공한 레이드 횟수
    plots_rented INTEGER DEFAULT 0, -- 임대한 구획 수
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 