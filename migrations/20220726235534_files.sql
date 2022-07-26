-- +goose Up
-- +goose StatementBegin
CREATE TABLE `files` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `hash_id` varchar(255) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `reference` varchar(255) DEFAULT NULL,
    `system` varchar(255) DEFAULT NULL,
    `type` varchar(255) DEFAULT NULL,
    `path` varchar(255) DEFAULT NULL,
    `name` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `hash_id` (`hash_id`),
    KEY `idx_files_deleted_at` (`deleted_at`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `files`
-- +goose StatementEnd
