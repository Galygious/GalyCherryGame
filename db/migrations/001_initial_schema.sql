-- Players table
CREATE TABLE IF NOT EXISTS players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    health INTEGER NOT NULL DEFAULT 100,
    level INTEGER NOT NULL DEFAULT 1,
    gold INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Inventory table
CREATE TABLE IF NOT EXISTS inventory (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    item_name TEXT NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY(player_id) REFERENCES players(id)
);

-- Skills table
CREATE TABLE IF NOT EXISTS skills (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    skill_name TEXT NOT NULL,
    level INTEGER NOT NULL DEFAULT 1,
    experience INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY(player_id) REFERENCES players(id)
);

-- Quests table
CREATE TABLE IF NOT EXISTS quests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    reward_gold INTEGER NOT NULL,
    reward_xp INTEGER NOT NULL,
    required_level INTEGER NOT NULL
);

-- Active Quests table
CREATE TABLE IF NOT EXISTS active_quests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    quest_id INTEGER NOT NULL,
    progress INTEGER NOT NULL DEFAULT 0,
    completed BOOLEAN NOT NULL DEFAULT 0,
    FOREIGN KEY(player_id) REFERENCES players(id),
    FOREIGN KEY(quest_id) REFERENCES quests(id)
);

-- Crafting Recipes table
CREATE TABLE IF NOT EXISTS crafting_recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    required_level INTEGER NOT NULL,
    required_items TEXT NOT NULL,
    output_item TEXT NOT NULL
);
