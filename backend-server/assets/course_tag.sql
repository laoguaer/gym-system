CREATE TABLE course_tag (
    id INT NOT NULL AUTO_INCREMENT,
    course_id INT NOT NULL,
    tag_id SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY k_tag_course (tag_id, course_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;