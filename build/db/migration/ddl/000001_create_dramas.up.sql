USE `quiz` ;

SET CHARSET utf8mb4;

CREATE TABLE
    IF NOT EXISTS quiz (
        id INT(11) AUTO_INCREMENT COMMENT "ID",
        problem VARCHAR(20) NOT NULL COMMENT "problem",
        answer VARCHAR(20) NOT NULL COMMENT "answer",
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "created date",
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "updated date",
        deleted_at DATETIME COMMENT "quiz deleted date",
        PRIMARY KEY (id)
    ) ENGINE = INNODB COMMENT = 'quiz';