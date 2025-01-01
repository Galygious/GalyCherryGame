ALTER TABLE players ADD COLUMN skill_points INTEGER NOT NULL DEFAULT 0;
ALTER TABLE players ADD COLUMN skill_cap INTEGER NOT NULL DEFAULT 100;

CREATE TABLE achievements (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    completed_at DATETIME,
    reward TEXT NOT NULL,
    FOREIGN KEY(player_id) REFERENCES players(id)
);
