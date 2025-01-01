CREATE TABLE combat_abilities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    damage_type TEXT NOT NULL,
    stamina_cost INTEGER NOT NULL,
    cooldown INTEGER NOT NULL,
    base_damage INTEGER NOT NULL,
    status_effect TEXT,
    required_level INTEGER NOT NULL,
    required_stat TEXT NOT NULL,
    required_stat_value INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE player_combat_abilities (
    player_id INTEGER NOT NULL,
    ability_id INTEGER NOT NULL,
    unlocked_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (player_id, ability_id),
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE,
    FOREIGN KEY (ability_id) REFERENCES combat_abilities(id) ON DELETE CASCADE
);

INSERT INTO combat_abilities (
    name, description, damage_type, stamina_cost, cooldown, base_damage, 
    status_effect, required_level, required_stat, required_stat_value
) VALUES 
('Power Strike', 'A powerful melee attack that deals extra damage', 'physical', 20, 5, 15, NULL, 1, 'strength', 5),
('Quick Shot', 'A fast ranged attack with a chance to hit twice', 'ranged', 15, 3, 10, NULL, 1, 'dexterity', 5),
('Fireball', 'A magical attack that deals fire damage over time', 'magic', 25, 8, 20, '{"type": "burn", "damage": 5, "duration": 3}', 1, 'magic', 5),
('Stunning Blow', 'A heavy attack that can stun the target', 'physical', 30, 10, 12, '{"type": "stun", "duration": 2}', 5, 'strength', 10),
('Poison Arrow', 'A poisoned arrow that deals damage over time', 'ranged', 25, 12, 8, '{"type": "poison", "damage": 3, "duration": 5}', 5, 'dexterity', 10),
('Ice Shard', 'A magical attack that slows the target', 'magic', 20, 6, 15, '{"type": "slow", "duration": 4}', 5, 'magic', 10);
