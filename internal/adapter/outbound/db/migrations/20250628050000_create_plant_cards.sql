-- migrate:up
-- 플랜트카드 테이블 (SNS 공유용 생성 콘텐츠)
CREATE TABLE plant_cards (
    id SERIAL PRIMARY KEY,
    farm_plot_id INTEGER NOT NULL,
    persona TEXT NOT NULL, -- AI 생성 페르소나
    image_url VARCHAR(255), -- 생성된 이미지 URL
    video_url VARCHAR(255), -- Veo3로 생성된 비디오 URL
    event_message TEXT, -- SNS 공유용 이벤트 메시지
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    
    -- 외래키 제약조건
    CONSTRAINT plant_cards_farm_plot_id_fkey 
        FOREIGN KEY (farm_plot_id) REFERENCES farm_plots(id) ON DELETE CASCADE
);

-- 인덱스 추가 (성능 최적화)
CREATE INDEX idx_plant_cards_farm_plot_id ON plant_cards(farm_plot_id);
CREATE INDEX idx_plant_cards_created_at ON plant_cards(created_at DESC);

-- migrate:down
DROP TABLE IF EXISTS plant_cards; 