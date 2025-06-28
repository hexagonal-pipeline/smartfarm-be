-- migrate:up
-- 임대 정보 테이블
CREATE TABLE rentals (
    id SERIAL PRIMARY KEY,
    renter_nickname VARCHAR(50) NOT NULL REFERENCES user_stats(nickname),
    plot_id INTEGER NOT NULL REFERENCES farm_plots(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    monthly_rent INTEGER NOT NULL,
    status VARCHAR(20) DEFAULT 'active', -- active, completed, cancelled
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 인덱스 생성 (성능용)
CREATE INDEX idx_rentals_nickname ON rentals(renter_nickname);
CREATE INDEX idx_rentals_plot_id ON rentals(plot_id);

-- migrate:down
DROP TABLE IF EXISTS rentals; 