-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频标题',
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '视频描述',
  `cover` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '封面图片地址',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频URL',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态(1-公开 2-私密 3-草稿)',
  `is_delete` tinyint(1) NULL DEFAULT 0 COMMENT '是否删除',
  `category_id` bigint NULL DEFAULT NULL COMMENT '分类ID',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_category_id`(`category_id` ASC) USING BTREE,
  INDEX `idx_user_id`(`user_id` ASC) USING BTREE,
  INDEX `idx_title`(`title` ASC) USING BTREE,
  CONSTRAINT `fk_video_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_video_user` FOREIGN KEY (`user_id`) REFERENCES `user_auth` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, '2023-12-28 10:00:00.000', '2023-12-28 10:00:00.000', 'Go语言入门教程', 'Go语言基础知识讲解，适合初学者学习', 'https://example.com/covers/go-tutorial.jpg', 'https://example.com/videos/go-tutorial.mp4', 1, 0, 1, 1);
INSERT INTO `video` VALUES (2, '2023-12-28 11:30:00.000', '2023-12-28 11:30:00.000', 'Vue3项目实战', 'Vue3框架实战项目开发教程', 'https://example.com/covers/vue3-project.jpg', 'https://example.com/videos/vue3-project.mp4', 1, 0, 2, 2);
INSERT INTO `video` VALUES (3, '2023-12-29 09:15:00.000', '2023-12-29 09:15:00.000', 'Docker容器化部署', 'Docker基础知识与实战部署教程', 'https://example.com/covers/docker-deploy.jpg', 'https://example.com/videos/docker-deploy.mp4', 1, 0, 1, 1);
INSERT INTO `video` VALUES (4, '2023-12-29 14:20:00.000', '2023-12-29 14:20:00.000', 'React组件开发', 'React组件化开发与状态管理', 'https://example.com/covers/react-components.jpg', 'https://example.com/videos/react-components.mp4', 2, 0, 2, 3);
INSERT INTO `video` VALUES (5, '2023-12-30 08:45:00.000', '2023-12-30 08:45:00.000', '微服务架构设计', '企业级微服务架构设计与实现', 'https://example.com/covers/microservice.jpg', 'https://example.com/videos/microservice.mp4', 1, 0, 1, 2);
INSERT INTO `video` VALUES (6, '2023-12-30 16:10:00.000', '2023-12-30 16:10:00.000', 'CSS动画效果实现', 'CSS高级动画效果与交互设计', 'https://example.com/covers/css-animation.jpg', 'https://example.com/videos/css-animation.mp4', 3, 0, 2, 3);
INSERT INTO `video` VALUES (7, '2023-12-31 10:30:00.000', '2023-12-31 10:30:00.000', '数据库优化技巧', 'MySQL数据库性能优化与实战', 'https://example.com/covers/db-optimization.jpg', 'https://example.com/videos/db-optimization.mp4', 1, 0, 1, 1);
INSERT INTO `video` VALUES (8, '2023-12-31 13:45:00.000', '2023-12-31 13:45:00.000', '前端工程化实践', '现代前端工程化工具与最佳实践', 'https://example.com/covers/frontend-engineering.jpg', 'https://example.com/videos/frontend-engineering.mp4', 1, 0, 2, 2);
INSERT INTO `video` VALUES (9, '2024-01-01 09:00:00.000', '2024-01-01 09:00:00.000', '健身计划制定', '科学健身计划制定与训练方法', 'https://example.com/covers/fitness-plan.jpg', 'https://example.com/videos/fitness-plan.mp4', 1, 0, 4, 3);
INSERT INTO `video` VALUES (10, '2024-01-01 15:20:00.000', '2024-01-01 15:20:00.000', '项目管理方法论', '敏捷开发与项目管理实践', 'https://example.com/covers/project-management.jpg', 'https://example.com/videos/project-management.mp4', 2, 0, 3, 1);
INSERT INTO `video` VALUES (11, '2024-01-02 11:10:00.000', '2024-01-02 11:10:00.000', 'Linux服务器配置', 'Linux服务器环境配置与优化', 'https://example.com/covers/linux-server.jpg', 'https://example.com/videos/linux-server.mp4', 1, 0, 1, 2);
INSERT INTO `video` VALUES (12, '2024-01-02 16:40:00.000', '2024-01-02 16:40:00.000', 'JavaScript高级编程', 'JavaScript高级特性与设计模式', 'https://example.com/covers/js-advanced.jpg', 'https://example.com/videos/js-advanced.mp4', 3, 0, 2, 3);