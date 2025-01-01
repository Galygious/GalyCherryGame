import { app, Component } from 'apprun';
import './style.css';

interface StatusEffect {
  type: string;
  damage?: number;
  duration: number;
  endTime: Date;
}

interface CombatAbility {
  id: number;
  name: string;
  description: string;
  damageType: 'physical' | 'ranged' | 'magic';
  staminaCost: number;
  cooldown: number;
  baseDamage: number;
  statusEffect?: string;
  requiredLevel: number;
  requiredStat: 'strength' | 'dexterity' | 'magic';
  requiredStatValue: number;
  lastUsed?: Date;
}

interface Player {
  name: string;
  health: number;
  maxHealth: number;
  level: number;
  experience: number;
  experienceToLevel: number;
  gold: number;
  strength: number;
  dexterity: number;
  magic: number;
  stamina: number;
  maxStamina: number;
  statusEffects: StatusEffect[];
  combatAbilities: CombatAbility[];
  skills: {
    combat: number;
    fishing: number;
    cooking: number;
    farming: number;
    crafting: number;
    alchemy: number;
  };
  inventory: {
    weapons: string[];
    armor: string[];
    consumables: string[];
    materials: string[];
  };
  activeQuests: string[];
  completedQuests: string[];
}

interface Enemy {
  name: string;
  health: number;
  maxHealth: number;
  level: number;
  maxDamage: number;
  defense: number;
  statusEffects: StatusEffect[];
}

interface CraftingStation {
  id: number;
  name: string;
  description: string;
  type: string;
  skillLevel: number;
  location: string;
}

interface GameState {
  player: Player;
  currentView: string;
  currentEnemy?: Enemy;
  inCombat: boolean;
  combatLog: string[];
  availableQuests: string[];
  shopItems: string[];
  craftingStations: CraftingStation[];
}

class GameApp extends Component<GameState> {
  state: GameState = {
    player: {
      name: 'Adventurer',
      health: 100,
      maxHealth: 100,
      level: 1,
      experience: 0,
      experienceToLevel: 100,
      gold: 0,
      strength: 5,
      dexterity: 5,
      magic: 5,
      stamina: 100,
      maxStamina: 100,
      statusEffects: [],
      combatAbilities: [
        {
          id: 1,
          name: 'Power Strike',
          description: 'A powerful melee attack that deals extra damage',
          damageType: 'physical',
          staminaCost: 20,
          cooldown: 5,
          baseDamage: 15,
          requiredLevel: 1,
          requiredStat: 'strength',
          requiredStatValue: 5
        },
        {
          id: 2,
          name: 'Quick Shot',
          description: 'A fast ranged attack with a chance to hit twice',
          damageType: 'ranged',
          staminaCost: 15,
          cooldown: 3,
          baseDamage: 10,
          requiredLevel: 1,
          requiredStat: 'dexterity',
          requiredStatValue: 5
        },
        {
          id: 3,
          name: 'Fireball',
          description: 'A magical attack that deals fire damage over time',
          damageType: 'magic',
          staminaCost: 25,
          cooldown: 8,
          baseDamage: 20,
          statusEffect: '{"type": "burn", "damage": 5, "duration": 3}',
          requiredLevel: 1,
          requiredStat: 'magic',
          requiredStatValue: 5
        }
      ],
      skills: {
        combat: 1,
        fishing: 1,
        cooking: 1,
        farming: 1,
        crafting: 1,
        alchemy: 1
      },
      inventory: {
        weapons: ['Wooden Sword'],
        armor: ['Cloth Tunic'],
        consumables: ['Health Potion'],
        materials: []
      },
      activeQuests: [],
      completedQuests: []
    },
    currentView: 'main',
    inCombat: false,
    combatLog: [],
    availableQuests: ['Goblin Slayer', 'Wolf Hunter'],
    shopItems: ['Iron Sword', 'Leather Armor'],
    craftingStations: []
  };

  componentDidMount() {
    this.fetchCraftingStations();
  }

  async fetchCraftingStations() {
    try {
      const response = await fetch('/crafting-stations');
      if (response.ok) {
        const stations = await response.json();
        app.run('setCraftingStations', stations);
      }
    } catch (error) {
      console.error('Failed to fetch crafting stations:', error);
    }
  }

  view = (state: GameState) => {
    const mainContent =
      state.currentView === 'main' ? (
        <div className="main-menu">
          <button onclick={() => app.run('navigate', 'combat')}>Combat</button>
          <button onclick={() => app.run('navigate', 'crafting')}>Crafting</button>
          <button onclick={() => app.run('navigate', 'quests')}>Quests</button>
          <button onclick={() => app.run('navigate', 'inventory')}>Inventory</button>
          <button onclick={() => app.run('navigate', 'skills')}>Skills</button>
          <button onclick={() => app.run('navigate', 'shop')}>Shop</button>
        </div>
      ) : state.currentView === 'crafting' ? (
        <div className="crafting-view">
          <h2>Crafting Stations</h2>
          <div className="stations-list">
            {state.craftingStations.map(station => (
              <div key={station.id} className="station-card">
                <h3>{station.name}</h3>
                <p>{station.description}</p>
                <div className="station-details">
                  <span>Type: {station.type}</span>
                  <span>Location: {station.location}</span>
                  <span>Required Skill: {station.skillLevel}</span>
                </div>
                <button 
                  onclick={() => app.run('useCraftingStation', station)}
                  disabled={state.player.skills.crafting < station.skillLevel}
                >
                  Use Station
                </button>
              </div>
            ))}
          </div>
          <button onclick={() => app.run('navigate', 'main')}>Return to Main Menu</button>
        </div>
      ) : state.currentView === 'combat' ? (
        <div className="combat-view">
          {!state.inCombat ? (
            <div className="combat-menu">
              <h2>Combat Arena</h2>
              <button onclick={() => app.run('startCombat', 'Goblin')}>Fight Goblin (Level 1)</button>
              <button onclick={() => app.run('startCombat', 'Wolf')}>Fight Wolf (Level 2)</button>
              <button onclick={() => app.run('startCombat', 'Orc')}>Fight Orc (Level 3)</button>
              <button onclick={() => app.run('navigate', 'main')}>Return to Main Menu</button>
            </div>
          ) : (
            <div className="combat-interface">
              <div className="enemy-info">
                <h3>{state.currentEnemy?.name}</h3>
                <div>Health: {state.currentEnemy?.health}</div>
                <div>Level: {state.currentEnemy?.level}</div>
                <div className="progress-bar">
                  <div className="progress" style={{ width: `${((state.currentEnemy?.health || 0) / (state.currentEnemy?.maxHealth || 1)) * 100}%` }}></div>
                </div>
              </div>
              <div className="player-info">
                <div>Health: {state.player.health}/{state.player.maxHealth}</div>
                <div>Level: {state.player.level} (XP: {state.player.experience}/{state.player.experienceToLevel})</div>
                <div>Gold: {state.player.gold}</div>
                <div className="progress-bar health-bar">
                  <div className="progress" style={{ width: `${(state.player.health / state.player.maxHealth) * 100}%` }}></div>
                </div>
                <div className="progress-bar xp-bar">
                  <div className="progress" style={{ width: `${(state.player.experience / state.player.experienceToLevel) * 100}%` }}></div>
                </div>
              </div>
              <div className="status-effects">
                {state.player.statusEffects.map((effect, i) => (
                  <div key={i} className={`status-effect ${effect.type}`}>
                    {effect.type}
                    <span className="status-effect-timer">
                      {Math.ceil((new Date(effect.endTime).getTime() - Date.now()) / 1000)}s
                    </span>
                  </div>
                ))}
              </div>
              <div className="abilities-list">
                {state.player.combatAbilities.map(ability => {
                  const onCooldown = ability.lastUsed && 
                    (Date.now() - new Date(ability.lastUsed).getTime()) / 1000 < ability.cooldown;
                  const notEnoughStamina = state.player.stamina < ability.staminaCost;
                  const notEnoughStat = state.player[ability.requiredStat] < ability.requiredStatValue;
                  
                  return (
                    <button 
                      key={ability.id}
                      className="ability-button"
                      onclick={() => app.run('useAbility', ability.id)}
                      disabled={onCooldown || notEnoughStamina || notEnoughStat}
                    >
                      {ability.name}
                      <div className="ability-tooltip">
                        {ability.description}<br/>
                        Damage: {ability.baseDamage}<br/>
                        Cost: {ability.staminaCost} stamina<br/>
                        Requires: {ability.requiredStat} {ability.requiredStatValue}
                        {onCooldown && <span><br/>Cooldown: {Math.ceil(ability.cooldown - (Date.now() - (ability.lastUsed ? ability.lastUsed.getTime() : Date.now())) / 1000)}s</span>}
                      </div>
                    </button>
                  );
                })}
              </div>
              <div className="combat-actions">
                <button onclick={() => app.run('attack')}>Basic Attack</button>
                <button onclick={() => app.run('useItem', 'Health Potion')} disabled={!state.player.inventory.consumables.includes('Health Potion')}>Use Health Potion</button>
                <button onclick={() => app.run('defend')}>Defend</button>
                <button onclick={() => app.run('flee')}>Flee</button>
              </div>
              <div className="combat-log">
                {state.combatLog.map((log, i) => (
                  <div key={i} className={`log-entry ${log.includes('damage') ? 'damage-message' : log.includes('Level Up') ? 'level-up-message' : log.includes('defeated') ? 'victory-message' : ''}`}>
                    {log}
                  </div>
                ))}
              </div>
            </div>
          )}
        </div>
      ) : state.currentView === 'quests' ? (
        <div className="quests-view">
          <h2>Quests</h2>
          <div className="active-quests">
            <h3>Active Quests</h3>
            {state.player.activeQuests.length > 0 ? (
              state.player.activeQuests.map((quest, i) => (
                <div key={i} className="quest">
                  {quest}
                </div>
              ))
            ) : (
              <div>No active quests</div>
            )}
          </div>
          <div className="available-quests">
            <h3>Available Quests</h3>
            {state.availableQuests.map((quest, i) => (
              <div key={i} className="quest">
                {quest}
                <button onclick={() => app.run('acceptQuest', quest)}>Accept</button>
              </div>
            ))}
          </div>
          <button onclick={() => app.run('navigate', 'main')}>Return to Main Menu</button>
        </div>
      ) : null;

    return (
      <div className="game-container">
        <header>
          <h1>GalyCherryGame</h1>
          <div className="player-stats">
            <div>Name: {state.player.name}</div>
            <div>Health: {state.player.health}/{state.player.maxHealth}</div>
            <div>Level: {state.player.level} (XP: {state.player.experience}/{state.player.experienceToLevel})</div>
            <div>Gold: {state.player.gold}</div>
          </div>
        </header>
        <main>{mainContent}</main>
      </div>
    );
  };

  update = {
    navigate: (state: GameState, view: string): GameState => ({
      ...state,
      currentView: view,
    }),
    setCraftingStations: (state: GameState, stations: CraftingStation[]): GameState => ({
      ...state,
      craftingStations: stations
    }),
    useCraftingStation: (_state: GameState): GameState => {
      // TODO: Implement crafting station interaction
      return _state;
    },
    startCombat: (state: GameState, enemyName: string): GameState => {
      const enemies: Record<string, Enemy> = {
        Goblin: { 
          name: 'Goblin', 
          health: 50, 
          maxHealth: 50, 
          level: 1, 
          maxDamage: 5, 
          defense: 2,
          statusEffects: []
        },
        Wolf: { 
          name: 'Wolf', 
          health: 75, 
          maxHealth: 75, 
          level: 2, 
          maxDamage: 8, 
          defense: 3,
          statusEffects: []
        },
        Orc: { 
          name: 'Orc', 
          health: 100, 
          maxHealth: 100, 
          level: 3, 
          maxDamage: 12, 
          defense: 5,
          statusEffects: []
        },
      };
      return {
        ...state,
        inCombat: true,
        currentEnemy: enemies[enemyName],
        combatLog: [`A wild ${enemyName} appears!`],
      };
    },
    useAbility: (state: GameState, abilityId: number): GameState => {
      if (!state.currentEnemy) return state;
      
      const ability = state.player.combatAbilities.find(a => a.id === abilityId);
      if (!ability) return state;

      // Check cooldown
      if (ability.lastUsed && 
          (Date.now() - new Date(ability.lastUsed).getTime()) / 1000 < ability.cooldown) {
        return state;
      }

      // Check stamina
      if (state.player.stamina < ability.staminaCost) {
        return state;
      }

      // Check stat requirement
      if (state.player[ability.requiredStat] < ability.requiredStatValue) {
        return state;
      }

      const newLog = [...state.combatLog];
      
      // Calculate damage based on ability type and relevant stat
      let statBonus = 0;
      switch (ability.damageType) {
        case 'physical':
          statBonus = state.player.strength * 2;
          break;
        case 'ranged':
          statBonus = state.player.dexterity * 2;
          break;
        case 'magic':
          statBonus = state.player.magic * 2;
          break;
      }
      
      const totalDamage = ability.baseDamage + statBonus;
      const enemyHealth = Math.max(0, state.currentEnemy.health - totalDamage);
      newLog.push(`You use ${ability.name} and deal ${totalDamage} damage to ${state.currentEnemy.name}!`);

      // Apply status effect if any
      if (ability.statusEffect) {
        const effect = JSON.parse(ability.statusEffect);
        const statusEffect = {
          type: effect.type,
          damage: effect.damage,
          duration: effect.duration,
          endTime: new Date(Date.now() + effect.duration * 1000)
        };
        state.currentEnemy.statusEffects.push(statusEffect);
        newLog.push(`${state.currentEnemy.name} is affected by ${effect.type}!`);
      }

      // Update ability cooldown
      const updatedAbilities = state.player.combatAbilities.map(a => 
        a.id === ability.id ? { ...a, lastUsed: new Date() } : a
      );

      // Use stamina
      const newStamina = state.player.stamina - ability.staminaCost;

      if (enemyHealth <= 0) {
        const goldEarned = state.currentEnemy.level * 10;
        const expEarned = state.currentEnemy.level * 25;
        const newExp = state.player.experience + expEarned;
        let newLevel = state.player.level;
        let newExpToLevel = state.player.experienceToLevel;
        let newMaxHealth = state.player.maxHealth;
        let newHealth = state.player.health;
        
        newLog.push(`You defeated ${state.currentEnemy.name} and earned ${goldEarned} gold and ${expEarned} XP!`);
        
        if (newExp >= state.player.experienceToLevel) {
          newLevel += 1;
          newExpToLevel = Math.floor(state.player.experienceToLevel * 1.5);
          newMaxHealth = state.player.maxHealth + 20;
          newHealth = newMaxHealth;
          newLog.push(`Level Up! You are now level ${newLevel}!`);
          newLog.push(`Your max health increased to ${newMaxHealth}!`);
        }
        
        return {
          ...state,
          inCombat: false,
          currentEnemy: undefined,
          combatLog: [],
          player: {
            ...state.player,
            level: newLevel,
            experience: newExp % newExpToLevel,
            experienceToLevel: newExpToLevel,
            maxHealth: newMaxHealth,
            health: newHealth,
            gold: state.player.gold + goldEarned,
            stamina: newStamina,
            combatAbilities: updatedAbilities
          },
        };
      }

      // Apply enemy status effects
      let enemyDamage = Math.floor(Math.random() * state.currentEnemy.maxDamage);
      let currentEnemyHealth = enemyHealth;
      const updatedEnemyStatusEffects = state.currentEnemy.statusEffects.filter(effect => {
        const now = new Date();
        if (now < new Date(effect.endTime)) {
          if (effect.damage) {
            currentEnemyHealth = Math.max(0, currentEnemyHealth - effect.damage);
            newLog.push(`${state.currentEnemy!.name} takes ${effect.damage} damage from ${effect.type}!`);
          }
          return true;
        }
        newLog.push(`${effect.type} effect on ${state.currentEnemy!.name} has worn off.`);
        return false;
      });

      // Enemy's turn
      return {
        ...state,
        currentEnemy: {
          ...state.currentEnemy,
          health: currentEnemyHealth,
          statusEffects: updatedEnemyStatusEffects
        },
        player: {
          ...state.player,
          stamina: newStamina,
          combatAbilities: updatedAbilities,
          health: Math.max(0, state.player.health - enemyDamage)
        },
        combatLog: [...state.combatLog, ...newLog].slice(-5)
      };
    },

    attack: (state: GameState): GameState => {
      if (!state.currentEnemy) return state;
      const newLog = [...state.combatLog];
      
      // Basic attack uses strength
      const baseDamage = 5 + state.player.strength;
      const totalDamage = Math.max(1, baseDamage - state.currentEnemy.defense);
      const enemyHealth = Math.max(0, state.currentEnemy.health - totalDamage);
      newLog.push(`You deal ${totalDamage} damage to ${state.currentEnemy.name}!`);

      if (enemyHealth <= 0) {
        const goldEarned = state.currentEnemy.level * 10;
        const expEarned = state.currentEnemy.level * 25;
        const newExp = state.player.experience + expEarned;
        let newLevel = state.player.level;
        let newExpToLevel = state.player.experienceToLevel;
        let newMaxHealth = state.player.maxHealth;
        let newHealth = state.player.health;
        
        newLog.push(`You defeated ${state.currentEnemy.name} and earned ${goldEarned} gold and ${expEarned} XP!`);
        
        if (newExp >= state.player.experienceToLevel) {
          newLevel += 1;
          newExpToLevel = Math.floor(state.player.experienceToLevel * 1.5);
          newMaxHealth = state.player.maxHealth + 20;
          newHealth = newMaxHealth;
          newLog.push(`Level Up! You are now level ${newLevel}!`);
          newLog.push(`Your max health increased to ${newMaxHealth}!`);
        }
        
        return {
          ...state,
          inCombat: false,
          currentEnemy: undefined,
          combatLog: [],
          player: {
            ...state.player,
            level: newLevel,
            experience: newExp % newExpToLevel,
            experienceToLevel: newExpToLevel,
            maxHealth: newMaxHealth,
            health: newHealth,
            gold: state.player.gold + goldEarned,
          },
        };
      }

      // Apply enemy status effects
      let enemyDamage = Math.floor(Math.random() * state.currentEnemy.maxDamage);
      let currentEnemyHealth = enemyHealth;
      const updatedEnemyStatusEffects = state.currentEnemy.statusEffects.filter(effect => {
        const now = new Date();
        if (now < new Date(effect.endTime)) {
          if (effect.damage) {
            currentEnemyHealth = Math.max(0, currentEnemyHealth - effect.damage);
            newLog.push(`${state.currentEnemy!.name} takes ${effect.damage} damage from ${effect.type}!`);
          }
          return true;
        }
        newLog.push(`${effect.type} effect on ${state.currentEnemy!.name} has worn off.`);
        return false;
      });
      
      return {
        ...state,
        currentEnemy: {
          ...state.currentEnemy,
          health: currentEnemyHealth,
          statusEffects: updatedEnemyStatusEffects
        },
        player: {
          ...state.player,
          health: Math.max(0, state.player.health - enemyDamage),
        },
        combatLog: [...state.combatLog, ...newLog].slice(-5),
      };
    },

    defend: (state: GameState): GameState => {
      if (!state.currentEnemy) return state;
      const newLog = [...state.combatLog];
      
      // Defending reduces damage taken by 50%
      const enemyDamage = Math.floor((Math.random() * state.currentEnemy.maxDamage) * 0.5);
      newLog.push(`You defend against ${state.currentEnemy.name}'s attack!`);
      newLog.push(`${state.currentEnemy.name} deals ${enemyDamage} damage to you!`);
      
      return {
        ...state,
        player: {
          ...state.player,
          health: Math.max(0, state.player.health - enemyDamage),
        },
        combatLog: [...state.combatLog, ...newLog].slice(-5),
      };
    },

    useItem: (state: GameState, item: string): GameState => {
      if (item === 'Health Potion') {
        const healAmount = 30;
        const newHealth = Math.min(state.player.maxHealth, state.player.health + healAmount);
        return {
          ...state,
          player: {
            ...state.player,
            health: newHealth,
            inventory: {
              ...state.player.inventory,
              consumables: state.player.inventory.consumables.filter(i => i !== item)
            }
          },
          combatLog: [...state.combatLog, `You used a Health Potion and healed ${healAmount} health!`]
        };
      }
      return state;
    },

    flee: (state: GameState): GameState => ({
      ...state,
      inCombat: false,
      currentEnemy: undefined,
      combatLog: [],
    }),

    acceptQuest: (state: GameState, quest: string): GameState => ({
      ...state,
      player: {
        ...state.player,
        activeQuests: [...state.player.activeQuests, quest]
      },
      availableQuests: state.availableQuests.filter(q => q !== quest)
    })
  };
}

const appElement = document.getElementById('app');
if (appElement) {
  try {
    const game = new GameApp();
    game.mount(appElement);
    app.start(appElement, game.state, game.view, game.update);
  } catch (error) {
    console.error('Error initializing game:', error);
  }
}
