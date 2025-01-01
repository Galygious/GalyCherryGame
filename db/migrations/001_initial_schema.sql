CREATE TABLE players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    health INTEGER NOT NULL DEFAULT 100,
    mana INTEGER NOT NULL DEFAULT 100,
    strength INTEGER NOT NULL DEFAULT 10,
    dexterity INTEGER NOT NULL DEFAULT 10,
    intellect INTEGER NOT NULL DEFAULT 10,
    level INTEGER NOT NULL DEFAULT 1,
    experience INTEGER NOT NULL DEFAULT 0,
    gold INTEGER NOT NULL DEFAULT 0,
    crafting INTEGER NOT NULL DEFAULT 1,
    alchemy INTEGER NOT NULL DEFAULT 1,
    fishing INTEGER NOT NULL DEFAULT 1,
    cooking INTEGER NOT NULL DEFAULT 1,
    farming INTEGER NOT NULL DEFAULT 1,
    combat INTEGER NOT NULL DEFAULT 1,
    weapon_name TEXT,
    weapon_description TEXT,
    armor_name TEXT,
    armor_description TEXT,
    accessory_name TEXT,
    accessory_description TEXT,
    cape TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE inventory (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY(player_id) REFERENCES players(id)
);

CREATE TABLE quests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    quest_id INTEGER NOT NULL,
    status TEXT NOT NULL DEFAULT 'active',
    progress INTEGER NOT NULL DEFAULT 0,
    started_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME,
    FOREIGN KEY(player_id) REFERENCES players(id)
);

CREATE TABLE crafting_recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    required_level INTEGER NOT NULL,
    required_items TEXT NOT NULL,
    output_item TEXT NOT NULL
);
