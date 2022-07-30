-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user_logins` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `user_id` int unsigned DEFAULT NULL,
    `token` varchar(255) DEFAULT NULL,
    `client_ip` varchar(50) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_user_logins_deleted_at` (`deleted_at`),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_logins`;
-- +goose StatementEnd
