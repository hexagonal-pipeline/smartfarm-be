-- migrate:up
-- 종합적인 더미데이터 추가

-- 1. 추가 농장 구획 데이터 (ID 명시적 지정)
INSERT INTO farm_plots (id, name, location, size_sqm, monthly_rent, crop_type, status, persona_prompt) VALUES 
(5, 'C구역-01', '전남 나주시', 180, 220000, '배추', 'available', 'The star of kimchi! I''m the cabbage king with white flesh and crispy texture.'),
(6, 'C구역-02', '전남 나주시', 160, 190000, '무', 'available', 'The white gem underground, radish! Feel my cool and sweet taste.'),
(7, 'D구역-01', '강원도 춘천시', 140, 170000, '감자', 'available', 'A potato grown in the cozy underground! My fluffy texture is my pride.'),
(8, 'D구역-02', '강원도 춘천시', 130, 160000, '당근', 'rented', 'The orange incarnation of vitamins! A healthy carrot is here.')
ON CONFLICT (id) DO NOTHING;

-- 시퀀스 값 업데이트
SELECT setval('farm_plots_id_seq', (SELECT MAX(id) FROM farm_plots));

-- 2. 추가 사용자 통계
INSERT INTO user_stats (nickname, level, experience, credit, total_revenue, successful_raids, plots_rented) VALUES 
('신입농부', 1, 30, 2000, 0, 0, 0),
('베테랑', 5, 650, 50000, 1200000, 12, 4),
('도시농부', 2, 180, 8000, 350000, 4, 2),
('친환경농부', 3, 320, 15000, 600000, 7, 3)
ON CONFLICT (nickname) DO NOTHING;

-- 3. 임대 데이터 (rentals)
INSERT INTO rentals (renter_nickname, plot_id, start_date, end_date, monthly_rent, status) VALUES 
('농부김씨', 2, '2025-01-01', '2025-12-31', 200000, 'active'),
('청년이', 8, '2025-02-01', '2026-01-31', 160000, 'active'),
('베테랑', 1, '2024-12-01', '2025-11-30', 150000, 'active'),
('도시농부', 3, '2025-01-15', '2026-01-14', 250000, 'active')
ON CONFLICT DO NOTHING;

-- 4. 추가 레이드 데이터
INSERT INTO raids (title, description, crop_type, target_quantity, min_participation, max_participation, price_per_kg, deadline, creator_nickname, status) VALUES 
('친환경 배추 대량주문', '유기농 배추 800kg 납품 프로젝트', '배추', 800, 15, 120, 2500, NOW() + INTERVAL '20 day', '친환경농부', 'open'),
('겨울 무 공급계약', '김장철 무 대량 공급', '무', 600, 12, 90, 2000, NOW() + INTERVAL '25 day', '베테랑', 'open'),
('감자칩 원료 공급', '가공업체 감자 납품', '감자', 1000, 20, 150, 1800, NOW() + INTERVAL '30 day', '농부김씨', 'open')
ON CONFLICT DO NOTHING;

-- 5. 레이드 참여 데이터
INSERT INTO raid_participations (raid_id, participant_nickname, quantity, expected_revenue, status) VALUES 
(1, '농부김씨', 100, 300000, 'confirmed'),
(1, '청년이', 80, 240000, 'confirmed'),
(1, '베테랑', 120, 360000, 'pending'),
(2, '스마트농부', 50, 200000, 'confirmed'),
(2, '도시농부', 70, 280000, 'pending'),
(3, '친환경농부', 90, 360000, 'confirmed'),
(4, '농부김씨', 150, 375000, 'confirmed'),
(4, '베테랑', 200, 500000, 'confirmed'),
(5, '청년이', 100, 200000, 'pending')
ON CONFLICT DO NOTHING;

-- 6. 위탁 작업 데이터 (commission_works)
INSERT INTO commission_works (requester_nickname, plot_id, task_type, task_description, credit_cost, status) VALUES 
('농부김씨', 1, '물주기', '매일 오전 8시 물주기 작업', 5000, 'requested'),
('청년이', 2, '잡초제거', '구역 전체 잡초 제거 및 정리', 15000, 'in_progress'),
('베테랑', 3, '수확', '딸기 수확 및 포장 작업', 25000, 'completed'),
('도시농부', 4, '비료주기', '유기농 비료 살포 작업', 8000, 'requested'),
('친환경농부', 1, '해충방제', '친환경 해충 방제 작업', 12000, 'in_progress')
ON CONFLICT DO NOTHING;

-- 7. 수익 기록 데이터 (revenue_records)
INSERT INTO revenue_records (nickname, type, amount, source_id, description) VALUES 
('농부김씨', 'raid', 300000, 1, '상추 레이드 참여 수익'),
('청년이', 'raid', 240000, 1, '상추 레이드 참여 수익'),
('베테랑', 'commission', 25000, 3, '딸기 수확 위탁 작업 완료'),
('스마트농부', 'raid', 200000, 2, '토마토 레이드 참여 수익'),
('도시농부', 'rental', -250000, 3, '농장 구획 임대료 지출'),
('친환경농부', 'raid', 360000, 3, '감자 레이드 참여 수익')
ON CONFLICT DO NOTHING;

-- 8. 크레딧 거래 데이터 (credit_transactions)
INSERT INTO credit_transactions (nickname, transaction_type, amount, related_id, description) VALUES 
('농부김씨', 'earn', 300000, 1, '레이드 참여 보상'),
('농부김씨', 'spend', -5000, 1, '위탁 작업 의뢰'),
('청년이', 'earn', 240000, 1, '레이드 참여 보상'),
('청년이', 'spend', -15000, 2, '위탁 작업 의뢰'),
('베테랑', 'earn', 25000, 3, '위탁 작업 완료'),
('스마트농부', 'earn', 200000, 2, '레이드 참여 보상'),
('도시농부', 'spend', -8000, 4, '위탁 작업 의뢰'),
('친환경농부', 'earn', 360000, 3, '레이드 참여 보상'),
('친환경농부', 'spend', -12000, 5, '위탁 작업 의뢰')
ON CONFLICT DO NOTHING;

-- 9. 플랜트 카드 데이터 (plant_cards)
INSERT INTO plant_cards (farm_plot_id, persona, image_url, video_url, event_message) VALUES
(1, 'I am the Lettuce Prince, the idol of the fresh salad world, destined to conquer all with my crispiness! I am growing fresh and vibrant today~', 'https://example.com/lettuce_card.jpg', 'https://example.com/lettuce_video.mp4', 'I drank plenty of water today and became even fresher!'),
(2, 'A tomato of passion, full of the sun''s flavor. My juice is unstoppable! Currently ripening to a brilliant red.', 'https://example.com/tomato_card.jpg', 'https://example.com/tomato_video.mp4', 'I''m becoming sweeter under the sunlight!'),
(3, 'The memory of a sweet and sour first love, that''s me, the Strawberry Princess. Growing in the shape of a heart!', 'https://example.com/strawberry_card.jpg', 'https://example.com/strawberry_video.mp4', 'My flowers have started blooming! Fruits will be coming soon.'),
(4, '200% hydration! A cool and chic cucumber, moist to the core. I''m busy growing long and straight!', 'https://example.com/cucumber_card.jpg', 'https://example.com/cucumber_video.mp4', 'I grew 5cm more today!')
ON CONFLICT DO NOTHING;

-- migrate:down
-- 역순으로 데이터 삭제
DELETE FROM plant_cards WHERE farm_plot_id IN (1, 2, 3, 4);
DELETE FROM credit_transactions WHERE nickname IN ('농부김씨', '청년이', '베테랑', '스마트농부', '도시농부', '친환경농부', '신입농부');
DELETE FROM revenue_records WHERE nickname IN ('농부김씨', '청년이', '베테랑', '스마트농부', '도시농부', '친환경농부');
DELETE FROM commission_works WHERE requester_nickname IN ('농부김씨', '청년이', '베테랑', '도시농부', '친환경농부');
DELETE FROM raid_participations WHERE participant_nickname IN ('농부김씨', '청년이', '베테랑', '스마트농부', '도시농부', '친환경농부');
DELETE FROM raids WHERE creator_nickname IN ('친환경농부', '베테랑');
DELETE FROM rentals WHERE renter_nickname IN ('농부김씨', '청년이', '베테랑', '도시농부');
DELETE FROM user_stats WHERE nickname IN ('신입농부', '베테랑', '도시농부', '친환경농부');
DELETE FROM farm_plots WHERE name LIKE 'C구역-%' OR name LIKE 'D구역-%'; 