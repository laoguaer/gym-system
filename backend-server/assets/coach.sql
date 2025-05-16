-- ----------------------------
-- Table structure for coach
-- ----------------------------
DROP TABLE IF EXISTS `coach`;
CREATE TABLE `coach`  (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(50) NOT NULL,
  `phone` VARCHAR(15) NOT NULL,
  `desc` TEXT NOT NULL,
  `avatar` VARCHAR(255) NOT NULL,
  `occupation` VARCHAR(1024) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  UNIQUE KEY (phone)
) ENGINE=InnoDB COMMENT='教练信息表';

-- ----------------------------
-- Records of coach
-- ----------------------------
INSERT INTO `coach` (`name`, `phone`, `desc`, `avatar`, `occupation`) VALUES
('李教练', '13800138001', '国家健身教练认证，专注于帮助客户实现健康减脂和肌肉塑造目标。拥有5年教练经验，已帮助128位学员达成健身目标，开设15门专业课程。擅长力量训练、有氧运动，平均评分4.8分。', 'https://randomuser.me/api/portraits/men/32.jpg', '力量训练、有氧运动'),
('王教练', '13800138002', '资深瑜伽教练，擅长帮助学员提高身体柔韧性和核心力量。拥有7年教练经验，已指导156位学员，开设12门专业课程。专注于瑜伽、普拉提教学，平均评分4.9分。', 'https://randomuser.me/api/portraits/men/33.jpg', '瑜伽、普拉提'),
('张教练', '13800138003', '专注于功能性训练和运动康复，帮助客户恢复运动机能和预防伤害。拥有4年教练经验，已服务98位学员，开设8门专业课程。平均评分4.7分。', 'https://randomuser.me/api/portraits/men/34.jpg', '功能性训练、康复训练'),
('刘教练', '13800138004', '前职业拳击运动员，现专注于搏击和拳击训练，提供高强度的有氧训练课程。拥有6年教练经验，已培训112位学员，开设10门专业课程。平均评分4.6分。', 'https://randomuser.me/api/portraits/men/35.jpg', '搏击、拳击'),
('陈教练', '13800138005', '前国家游泳队队员，擅长游泳技术指导和水中有氧训练。拥有8年教练经验，已指导145位学员，开设14门专业课程。专业水中训练指导，平均评分4.9分。', 'https://randomuser.me/api/portraits/men/36.jpg', '游泳、水中有氧'),
('赵教练', '13800138006', '健美比赛冠军，专注于肌肉增长和体型塑造训练。拥有5年教练经验，已指导132位学员，开设11门专业课程。专业增肌训练指导，平均评分4.8分。', 'https://randomuser.me/api/portraits/men/37.jpg', '健美、增肌训练');