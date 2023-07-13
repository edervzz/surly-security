package migrations

import "github.com/edervzz/maya"

var SecurityDBTables = maya.Migration{
	ID: "0000000001",
	Up: []string{
		users,
		token_users,
	},
	Down: []string{
		"DROP TABLE IF EXISTS token_users",
		"DROP TABLE IF EXISTS users",
	},
}

const users = `CREATE TABLE
    users (
        id BIGINT auto_increment NOT NULL COMMENT 'User ID',
        email varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Email address for user',
        fullname varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Full name',
        password_hash varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'Password hashed',
        email_confirmed tinyint (1) DEFAULT NULL COMMENT 'User confirmed via email',
        phone_confirmed tinyint (1) DEFAULT NULL COMMENT 'User confirmed via phone',
        is_locked tinyint (1) DEFAULT NULL COMMENT 'User is locked',
        is_active tinyint (1) DEFAULT NULL COMMENT 'User is active',
        intents tinyint DEFAULT NULL COMMENT 'Failed intents',
        created_by varchar(100) NOT NULL,
        created_at datetime NOT NULL,
        updated_by varchar(100) NOT NULL,
        updated_at datetime NOT NULL,
        PRIMARY KEY (id),
        UNIQUE KEY User_UN (email)
    ) AUTO_INCREMENT = 1;
    `

const token_users = `CREATE TABLE
    token_users (
        token varchar(1024) CHARACTER SET latin1 COLLATE latin1_general_cs NOT NULL COMMENT 'Token assign to only one user.',
        type varchar(20) NOT NULL CHECK (type in ("access","refresh","resetPassword","confirmEmail")) COMMENT 'Token type: access, refresh, resetPassword, confirmEmail.' ,
        user_id bigint NOT NULL COMMENT 'User identification',
        expires_at datetime NOT NULL COMMENT 'Expiration time in seconds.',
        is_active tinyint (1) NOT NULL COMMENT 'User is active',
        created_by varchar(100) NOT NULL,
        created_at datetime NOT NULL,
        updated_by varchar(100) NOT NULL,
        updated_at datetime NOT NULL,
        PRIMARY KEY (token),
        KEY token_users_FK (user_id),
        CONSTRAINT token_users_FK FOREIGN KEY (user_id) REFERENCES users (id)
    );
    `
