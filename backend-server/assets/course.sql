CREATE TABLE `course` (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `description` text NOT NULL,
    `tags` varchar(255) NOT NULL,
    `start_time` datetime NOT NULL,
    `end_time` datetime NOT NULL,
    `coach_id` int NOT NULL,
    `max_capacity` int NOT NULL DEFAULT '30',
    `is_single` TINYINT NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci

INSERT INTO course (title, description, tags, start_time, end_time, coach_id, max_capacity, is_single)
VALUES
('AI体态评估私教课', '结合3D动作捕捉技术实时分析体态数据，自动生成矫正训练方案，含肩颈平衡/骨盆矫正等12种评估维度', '康复,科技,数据分析', '2025-06-01 18:30:00', '2025-08-31 20:00:00', 201, 15, b'0'),
('HIIT智能燃脂团课', '通过可穿戴设备实时监测心率变化，动态调整训练强度，配备LED屏幕显示全员能量消耗排行榜', '减脂,竞技,物联网', '2025-05-15 19:00:00', '2025-05-15 20:30:00', 202, 25, b'1'), 
('VR搏击竞技课程', '使用HTC Vive Pro 2头显实现沉浸式格斗训练，支持多人在线对战模式，自动记录出拳速度与力量数据', '游戏化,VR技术,心肺', '2025-07-01 14:00:00', '2025-07-01 15:30:00', 203, 10, b'1'),
('产后修复云指导', '通过智能压力传感垫采集盆底肌收缩数据，结合APP远程评估生成阶段性恢复计划，含中医理疗方案', '康复,远程医疗,女性', '2025-09-01 09:00:00', '2025-12-01 11:00:00', 204, 20, b'0'),
('代谢灵活性训练', '采用连续血糖监测仪+CGM数据分析，定制碳水循环方案，同步监测静息代谢率变化', '营养,生物数据,进阶', '2025-08-10 07:00:00', '2025-08-10 08:30:00', 205, 12, b'1');
