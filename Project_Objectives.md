# GalyCherryGame - Menu-Based RPG

## Overview
GalyCherryGame is a fully-playable menu-based RPG featuring strategic progression through combat stats, diverse skills (crafting, alchemy), dungeon exploration, slayer quests, and long-term goals (earning high-tier capes).

## Tech Stack
- **Front End:** AppRun (lightweight library for reactive web apps)
- **Back End:** Go (Golang) for server-side logic
- **Database:** SQLite (file-based SQL)
- **Web Server:** Go's `net/http`
- **APIs & Framework:** Gin for building high-performance Go APIs
- **Infrastructure:** Docker for containerization
- **DevOps:** GitHub for CI/CD, Dependabot for dependencies, and OAuth2 for authentication/security

## Getting Started
1. Clone the repository
2. Run `docker-compose up` to start the application
3. Access the game at `http://localhost:8080`

## Project Structure
```
/project-root
├── /frontend  (AppRun + Vite)
├── /backend   (Go server)
├── /db        (SQLite + migrations)
├── /docker    (Dockerfile)
└── /docs      (Markdown docs)
```

### DETAILS
---

### 1. Context and Purpose
- Completely polished and market ready production game.
- Specify the platform: Web/Docker Containers

### 2. Game Concept
- **Genre:** Menu-based RPG
- **Theme:** Fantasy
- **Core Gameplay Loop:** Players level up their character by performing a variety of activities such as combat, fishing, cooking, farming, and crafting. They engage in dungeon exploration to fight mobs, complete slayer quests targeting specific enemies, and gather resources for crafting and alchemy. Players also manage their character's stats and equipment, trade with NPCs or other players, and work towards achieving milestone objectives.
- **Player Objective:** The ultimate goal is to maximize character progression by acquiring high-tier capes such as the Mini Max Cape, Max Cape, and No Life Cape, while completing quests, upgrading equipment, and mastering various skills.

### 3. Design Overview
- **Art Style:** Text-based with minimal visual enhancements (e.g., simple icons for menus and skills) or optional pixel-art style for nostalgic appeal.
- **Sound Design:** Relaxing background music to enhance immersion in the fantasy theme, distinct sound effects for actions like combat, crafting, and fishing, and subtle ambient sounds for dungeon exploration and the marketplace. No voiceovers are required.
- **Narrative:** Set in a fantastical world where the player assumes the role of an adventurer striving to master various skills and uncover hidden challenges. The world is filled with dungeons, mythical creatures, and opportunities for discovery, blending personal progression with an open-ended journey.
- **UI/UX:**
  - **Start Screen:** A simple menu offering options like "New Game," "Continue," "Settings," and "Exit."
  - **Main Interface:** Displays essential player stats (health, experience, gold) and an action menu for tasks like combat, crafting, and trading.
  - **Inventory:** Organized sections for equipment, consumables, and crafting materials, with tooltips for item descriptions.
  - **Quest Log:** Tracks active and completed quests, milestones, and objectives.
  - **Pause Menu:** Provides access to game settings, save/load options, and return to the main menu.

### 4. Mechanics and Features
- **Movement:** None (menu-based interactions; actions are chosen through menus rather than direct character movement).
- **Combat:**
  - **Melee:** Strength-based attacks for close-range combat.
  - **Ranged:** Archery-based attacks using equipped arrows.
  - **Abilities:** Combat modifiers like special blessings or equipped item effects (e.g., instant kills or critical hit bonuses).
  - **Magic:** Not explicitly stated but could include alchemy for combat enhancements (e.g., potions).
- **Progression:**
  - **Levels:** Gain experience in skills such as combat, fishing, cooking, farming, crafting, and alchemy.
  - **Experience:** Earned through performing actions, completing quests, and defeating enemies.
  - **Upgrades:** Equipment upgrades through crafting and trading; milestone-based objectives like acquiring high-tier capes.
- **Interactions:**
  - **NPC Dialogues:** Trade, accept quests, and gather information.
  - **Pickups:** Gather resources from farming, fishing, and dungeon exploration.
  - **Crafting:** Combine gathered resources to create items, equipment, and alchemical products.
  - **Quests:** Slayer tasks and milestone achievements guide progression.
- **Procedural Systems:**
  - **Randomized Drops:** Rare item drops from specific enemies.
  - **Skill Growth:** Skill-specific experience requirements vary, adding strategic depth to progression.
- **Multiplayer:**
  - **Online Trading:** Players can trade items with other players through a market system.
  - **Leaderboard (Optional):** Showcase achievements like high-tier capes and skill mastery.

### 5. Development Phases
#### a. Setup
- **Environment Setup:**
  - Initialize a GitHub repository and set up CI/CD pipelines using GitHub Actions.
  - Create Docker containers for the development environment, including Go, SQLite, and Vite.
  - Install required libraries and dependencies:
    - **Frontend:** AppRun, Vite, ESLint, Prettier.
    - **Backend:** Gin, Go's `crypto` package, OAuth2, Kong API Gateway.

#### b. Prototyping
- **Prototype Tasks:**
  - Set up basic API endpoints in Gin for player stats, inventory, and actions.
  - Build a simple frontend interface to interact with the backend (e.g., menus for combat, crafting, and fishing).
  - Use SQLite to store mock data for player progression and items.
  - Ensure communication between frontend and backend is functional using dummy data.

#### c. Core Development
- **Tasks:**
  - Expand the backend to handle core mechanics: combat calculations, skill progression, and item crafting.
  - Add visual elements to the frontend for UI enhancements (e.g., progress bars, item icons).
  - Implement the game loop:
    - Backend handles actions (e.g., fighting mobs, completing quests).
    - Frontend updates dynamically based on backend responses.
  - Integrate OAuth2 for user authentication.

#### d. Content Creation
- **Tasks:**
  - Add data to the SQLite database for dungeons, mobs, items, and quests.
  - Develop additional UI components for trading, crafting, and slayer tasks.
  - Write and integrate milestone quests and narrative dialogues.
  - Create templates for capes and other high-tier equipment progression.

#### e. Final Polish
- **Optimization:**
  - Compress assets and optimize frontend performance using Vite.
  - Fine-tune database queries and server response times.
- **Sound and Music:**
  - Add ambient soundtracks and sound effects for actions like fishing, combat, and crafting.
- **Bug Fixes:**
  - Perform thorough testing for edge cases and fix all identified issues.

#### f. Packaging
- **Tasks:**
  - Use Docker to containerize the application for deployment.
  - Ensure compatibility with self-hosted infrastructure.
  - Provide an easy setup guide in the Markdown documentation for installation and running.
  - Export a final build of the frontend using Vite, ready to be served via the backend.

### 6. Documentation
- **Code Comments and Explanation:**
  - Include inline comments throughout the code to explain key logic, especially for:
    - Core mechanics (e.g., combat, crafting, skill progression).
    - API endpoints in the backend (e.g., what each route does and what data it returns).
    - Database schema and queries for clarity on how data is stored and accessed.
  - Add JSDoc-style comments for frontend functions and GoDoc comments for backend functions and types.

- **Game Guide:**
  - Provide a Markdown file (`/docs/game_guide.md`) with the following sections:
    - **Overview:** Explain the game concept, genre, and mechanics.
    - **Getting Started:** Describe how to create a new character and start playing.
    - **Gameplay:** Detail the actions players can perform (e.g., combat, crafting, fishing).
    - **Progression:** Outline how to level up, unlock new quests, and earn milestone capes.
    - **Tips and Tricks:** Provide strategies for maximizing efficiency in different skills.

- **Setup Instructions:**
  - Include a clear `README.md` file in the root of the project with the following:
    - **Prerequisites:**
      - Docker and Docker Compose installed.
      - Git installed to clone the repository.
    - **Setup Steps:**
      1. Clone the repository:
         ```bash
         git clone https://github.com/your-repo-name.git
         cd your-repo-name
         ```
      2. Build and run the Docker containers:
         ```bash
         docker-compose up --build
         ```
      3. Access the game: Open a web browser and go to `http://localhost:8080`.

    - **Environment Variables:** Provide examples for configuring API keys, database paths, or OAuth2 settings in a `.env` file:
      ```
      DATABASE_PATH=./db/game.db
      OAUTH_CLIENT_ID=your-client-id
      OAUTH_SECRET=your-secret
      ```

- **API Documentation:**
  - Use Swagger or a simple Markdown file (`/docs/api_reference.md`) to list all available API endpoints, their parameters, and response formats.
  - Example:
    ```
    POST /api/combat/attack
    Request Body:
    {
        "playerId": "12345",
        "enemyId": "67890",
        "attackType": "melee"
    }

    Response:
    {
        "success": true,
        "damageDealt": 25,
        "playerHealth": 75,
        "enemyHealth": 50
    }
    ```

- **Troubleshooting Guide:**
  - Add a section to the `README.md` for common errors and their solutions (e.g., "Database connection error" or "API Gateway not responding").

### 9. Output Expectations
- **Code:**
  - Fully functional and well-documented game code for both frontend and backend, adhering to the provided tech stack.
  - Organized project structure with clear separation of concerns (e.g., UI, API, and database logic).
  - Inline comments and documentation for all major functions and systems to ensure maintainability.

- **Assets:**
  - Basic visual assets for menus and UI elements (e.g., icons for skills, progress bars, and equipment).
  - Sound assets, including ambient music, sound effects for combat and crafting, and subtle audio cues for user actions.
  - Optional: Placeholder or minimal pixel art for any aesthetic enhancement.

- **Build:**
  - A Dockerized build of the game, ready to deploy on self-hosted infrastructure.
  - A web-based application accessible through a browser, with a fully bundled and optimized frontend (via Vite).
  - Backend API fully integrated and functional with all game mechanics.

- **Guide:**
  - Comprehensive Markdown documentation, including:
    - Setup instructions for local and production environments.
    - Game mechanics guide explaining how the game works, key actions, and progression systems.
    - API reference for developers interested in extending or maintaining the backend.
    - Troubleshooting tips for common issues.
  - Inline code comments and structured README files in each project directory.
