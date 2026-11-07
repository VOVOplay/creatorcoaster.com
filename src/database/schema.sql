CREATE TABLE IF NOT EXISTS staff_list(
    user_id BIGINT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    pfp_link VARCHAR(500) NOT NULL,
    position VARCHAR(50) NOT NULL,
    position_pretty_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS wiki_categories(
    id INT AUTO_INCREMENT PRIMARY KEY,
    pretty_name VARCHAR(200) NOT NULL,
    sort_order INT NOT NULL DEFAULT 9999,
    parent_category_id INT DEFAULT NULL,
    FOREIGN KEY (parent_category_id) REFERENCES wiki_categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wiki_articles(
    name VARCHAR(200) PRIMARY KEY,
    pretty_name VARCHAR(200) NOT NULL,
    md_source TEXT NOT NULL,
    parent_category_id INT DEFAULT NULL,
    last_edited_at BIGINT NOT NULL,
    sort_order INT NOT NULL DEFAULT 9999,
    FOREIGN KEY (parent_category_id) REFERENCES wiki_categories(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS blog_posts(
    name VARCHAR(200) PRIMARY KEY,
    pretty_name VARCHAR(200) NOT NULL,
    md_source TEXT NOT NULL,
    date BIGINT NOT NULL,
    author VARCHAR(100) NOT NULL,
    author_pfp_link VARCHAR(500) NOT NULL,
    read_time VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS markdown_variables(
    variable_name VARCHAR(200) PRIMARY KEY,
    variable_contents TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS markdown_channel_variables(
    variable_name VARCHAR(200) PRIMARY KEY,
    channel_name VARCHAR(200) NOT NULL,
    channel_id BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS markdown_role_user_variables(
    variable_name VARCHAR(200) PRIMARY KEY,
    role_user_name VARCHAR(500) NOT NULL,
    role_user_id BIGINT NOT NULL,
    variable_type INT NOT NULL
);