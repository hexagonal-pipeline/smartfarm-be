-- migrate:up
-- 커미션(위탁) 작업 테이블
CREATE TABLE commission_works (
    id SERIAL PRIMARY KEY,
    requester_nickname VARCHAR(50) NOT NULL REFERENCES user_stats(nickname),
    plot_id INTEGER NOT NULL REFERENCES farm_plots(id),
    task_type VARCHAR(50) NOT NULL, -- 작업 유형 (예: '해충 방제', '수확', '토지 정리')
    task_description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'requested', -- 작업 상태 (requested, in_progress, completed, cancelled)
    credit_cost INTEGER NOT NULL,
    requested_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP WITH TIME ZONE
);

-- 크레딧 거래 내역 테이블
CREATE TABLE credit_transactions (
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL REFERENCES user_stats(nickname),
    transaction_type VARCHAR(50) NOT NULL, -- 거래 유형 (예: '커미션_비용', '레이드_보상', '크레딧_충전')
    amount INTEGER NOT NULL, -- 거래량 (음수: 사용, 양수: 획득)
    related_id INTEGER, -- 관련 데이터 ID (예: commission_works의 ID)
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 인덱스 추가 (성능 최적화)
CREATE INDEX idx_commission_works_requester ON commission_works(requester_nickname);
CREATE INDEX idx_commission_works_plot_id ON commission_works(plot_id);
CREATE INDEX idx_credit_transactions_nickname ON credit_transactions(nickname); 

-- migrate:down
DROP TABLE IF EXISTS credit_transactions;
DROP TABLE IF EXISTS commission_works; 