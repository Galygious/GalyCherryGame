import { app, Component } from 'apprun';
import './style.css';

interface Player {
  name: string;
  health: number;
  level: number;
  gold: number;
}

interface GameState {
  player: Player;
  currentView: string;
}

class GameApp extends Component<GameState> {
  state: GameState = {
    player: {
      name: 'Adventurer',
      health: 100,
      level: 1,
      gold: 0,
    },
    currentView: 'main',
  };

  view = (state: GameState) => {
    return `
      <div class="game-container">
        <header>
          <h1>GalyCherryGame</h1>
          <div class="player-stats">
            <div>Name: ${state.player.name}</div>
            <div>Health: ${state.player.health}</div>
            <div>Level: ${state.player.level}</div>
            <div>Gold: ${state.player.gold}</div>
          </div>
        </header>
        <main>
          ${this.getCurrentView(state)}
        </main>
      </div>
    `;
  };

  getCurrentView = (state: GameState) => {
    switch (state.currentView) {
      case 'main':
        return `
          <div class="main-menu">
            <button onclick="app.run('navigate', 'combat')">Combat</button>
            <button onclick="app.run('navigate', 'crafting')">Crafting</button>
            <button onclick="app.run('navigate', 'quests')">Quests</button>
          </div>
        `;
      default:
        return '<div>Coming Soon</div>';
    }
  };

  update = {
    navigate: (state: GameState, view: string) => ({
      ...state,
      currentView: view,
    }),
  };
}

const appElement = document.getElementById('app');
if (appElement) {
  const game = new GameApp();
  game.mount(appElement);
}
