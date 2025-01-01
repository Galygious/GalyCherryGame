CREATE TABLE players_backup (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    health INTEGER NOT NULL,
    max_health INTEGER NOT NULL,
    level INTEGER NOT NULL,
    experience INTEGER NOT NULL,
    experience_to_level INTEGER NOT NULL,
    gold INTEGER NOT NULL,
    combat_skill INTEGER NOT NULL,
    fishing_skill INTEGER NOT NULL,
    cooking_skill INTEGER NOT NULL,
    farming_skill INTEGER NOT NULL,
    crafting_skill INTEGER NOT NULL,
    alchemy_skill INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

CREATE TABLE inventory_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    quantity INTEGER NOT NULL,
    attack INTEGER,
    defense INTEGER,
    magic_power INTEGER,
    durability INTEGER,
    type TEXT NOT NULL,
    FOREIGN KEY(player_id) REFERENCES players(id)
);

CREATE TABLE player_quests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    quest_id INTEGER NOT NULL,
    status TEXT NOT NULL,
    progress INTEGER NOT NULL,
    started_at DATETIME NOT NULL,
    completed_at DATETIME,
    FOREIGN KEY(player_id) REFERENCES players(id)
);
