package models

type Enemy struct {
    Name           string          `json:"name"`
    Health         int             `json:"health"`
    MaxHealth      int             `json:"maxHealth"`
    Level          int             `json:"level"`
    MaxDamage      int             `json:"maxDamage"`
    Defense        int             `json:"defense"`
    AttackSpeed    int             `json:"attackSpeed"`
    SpecialAbility *SpecialAbility `json:"specialAbility,omitempty"`
}

type SpecialAbility struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Cooldown    int    `json:"cooldown"`
    Effect      string `json:"effect"`
}
