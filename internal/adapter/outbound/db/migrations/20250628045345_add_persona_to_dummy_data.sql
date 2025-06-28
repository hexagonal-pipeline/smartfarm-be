-- migrate:up
-- 더미 데이터 삽입 (데모용)
INSERT INTO farm_plots (name, location, size_sqm, monthly_rent, crop_type, status) VALUES 
('A구역-01', '경기도 파주시', 100, 150000, '상추', 'available'),
('A구역-02', '경기도 파주시', 150, 200000, '토마토', 'rented'),
('B구역-01', '충남 천안시', 200, 250000, '딸기', 'available'),
('B구역-02', '충남 천안시', 120, 180000, '오이', 'available');

INSERT INTO user_stats (nickname, level, experience, credit, total_revenue, successful_raids, plots_rented) VALUES 
('농부김씨', 3, 250, 10000, 500000, 5, 2),
('청년이', 2, 150, 5000, 200000, 3, 1),
('스마트농부', 1, 50, 0, 50000, 1, 1),
('레이드마스터', 4, 400, 20000, 800000, 8, 3);

INSERT INTO raids (title, description, crop_type, target_quantity, min_participation, max_participation, price_per_kg, deadline, creator_nickname) VALUES 
('대형마트 상추 납품', '대형마트에서 신선한 상추 500kg을 주문했습니다', '상추', 500, 10, 100, 3000, NOW() + INTERVAL '10 day', '농부김씨'),
('학교급식 토마토 공급', '인근 학교 급식용 토마토 대량 주문', '토마토', 300, 20, 80, 4000, NOW() + INTERVAL '15 day', '청년이');

-- 페르소나 프롬프트 추가
UPDATE farm_plots SET persona_prompt = '나는 겉절이계의 아이돌, 아삭함으로 세상을 평정할 상추王子' WHERE crop_type = '상추';
UPDATE farm_plots SET persona_prompt = '태양의 맛을 가득 품은 정열의 토마토. 내 과즙은 멈추지 않아!' WHERE crop_type = '토마토';
UPDATE farm_plots SET persona_prompt = '새콤달콤한 첫사랑의 기억, 그게 바로 나, 딸기 공주님이야.' WHERE crop_type = '딸기';
UPDATE farm_plots SET persona_prompt = '수분감 200%! 쿨하고 시크한 오이. 내면은 촉촉해.' WHERE crop_type = '오이';

-- migrate:down
DELETE FROM raids;
DELETE FROM user_stats;
DELETE FROM farm_plots;
-- persona_prompt는 farm_plots 삭제 시 함께 사라지므로 별도 down 마이그레이션 불필요

