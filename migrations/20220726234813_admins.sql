-- +goose Up
-- +goose StatementBegin
CREATE TABLE `admins` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `account` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `status` int DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `account` (`account`),
    UNIQUE KEY `password` (`password`),
    KEY `idx_admins_deleted_at` (`deleted_at`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `admins`
-- +goose StatementEnd
