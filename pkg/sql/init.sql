-- expedition関連の制約
ALTER TABLE expedition_images
DROP CONSTRAINT fk_expeditions_expedition_images,
ADD CONSTRAINT fk_expeditions_expedition_images
FOREIGN KEY (expedition_id)
REFERENCES expeditions(id)
ON DELETE CASCADE;

ALTER TABLE payments
DROP CONSTRAINT fk_expeditions_payments,
ADD CONSTRAINT fk_expeditions_payments
FOREIGN KEY (expedition_id)
REFERENCES expeditions(id)
ON DELETE CASCADE;

ALTER TABLE visited_facilities
DROP CONSTRAINT fk_expeditions_visited_facilities,
ADD CONSTRAINT fk_expeditions_visited_facilities
FOREIGN KEY (expedition_id)
REFERENCES expeditions(id)
ON DELETE CASCADE;

ALTER TABLE games
DROP CONSTRAINT fk_expeditions_games,
ADD CONSTRAINT fk_expeditions_games
FOREIGN KEY (expedition_id)
REFERENCES expeditions(id)
ON DELETE CASCADE;

-- game関連の制約
ALTER TABLE game_scores
DROP CONSTRAINT fk_games_game_scores,
ADD CONSTRAINT fk_games_game_scores
FOREIGN KEY (game_id)
REFERENCES games(id)
ON DELETE CASCADE;

-- team関連の制約
ALTER TABLE favorite_teams
DROP CONSTRAINT fk_teams_favorite_teams,
ADD CONSTRAINT fk_teams_favorite_teams
FOREIGN KEY (team_id)
REFERENCES teams(id)
ON DELETE CASCADE;

-- expedition_likes関連の制約
ALTER TABLE expedition_likes
DROP CONSTRAINT fk_expedition_likes_expedition,
ADD CONSTRAINT fk_expedition_likes_expedition
FOREIGN KEY (expedition_id)
REFERENCES expeditions(id)
ON DELETE CASCADE;



-- スポーツの登録（野球）
INSERT INTO sports (id, name, created_at, updated_at) 
VALUES (1, '野球', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- スタジアムの登録
INSERT INTO stadia (id, name, description, address, capacity, image, file_id, created_at, updated_at) 
VALUES 
(1, '東京ドーム', '東京の象徴的なドーム球場', '東京都文京区後楽1-3-61', 45000, 'tokyo_dome.jpg', 'stadium_tokyo_dome_1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '横浜スタジアム', '横浜の海に面した球場', '神奈川県横浜市中区横浜公園', 34000, 'yokohama_stadium.jpg', 'stadium_yokohama_2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, '神宮球場', '都心の緑に囲まれた球場', '東京都新宿区霞ヶ丘町3-1', 37000, 'jingu_stadium.jpg', 'stadium_jingu_3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, '阪神甲子園球場', '高校野球の聖地', '兵庫県西宮市甲子園町1-82', 47000, 'koshien_stadium.jpg', 'stadium_koshien_4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, 'マツダスタジアム', '広島の新しいランドマーク', '広島県広島市南区南蟹屋2-3-1', 32000, 'mazda_stadium.jpg', 'stadium_mazda_5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- リーグの登録（セ・リーグ）
INSERT INTO leagues (id, name, sport_id, created_at, updated_at) 
VALUES (1, 'セ・リーグ', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- チームの登録
INSERT INTO teams (id, name, stadium_id, league_id, sport_id, created_at, updated_at) 
VALUES 
(1, '読売ジャイアンツ', 1, 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, '横浜DeNAベイスターズ', 2, 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, '東京ヤクルトスワローズ', 3, 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, '阪神タイガース', 4, 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(5, '広島東洋カープ', 5, 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);