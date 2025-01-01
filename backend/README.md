# Backend Overview

This directory contains the backend logic for the **GalyCherryGame** project. The backend is built using **Go (Golang)** and provides APIs, data handling, and other server-side functionality.

---

## Directory Structure

```plaintext
backend/
├── models/         # Data models and schema definitions
├── pkg/            # Reusable utility packages
├── static/         # Static assets served by the backend
├── api.go          # API routes and handlers
├── main.go         # Entry point for the backend application
├── go.mod          # Go module dependencies
```

---

## Key Files

### `main.go`
- **Purpose:** The entry point for the backend application.
- **Responsibilities:**
  - Parses command-line flags (e.g., `-migrate` for database migrations).
  - Initializes the SQLite database using `db.InitDB()`.
  - Configures and starts the Gin web server.
  - Serves static assets and handles SPA routing.
- **Key Features:**
  - Supports running database migrations with the `-migrate` flag.
  - Uses Gin framework for HTTP routing and middleware.
  - Serves frontend assets from the `static` directory.
- **Environment Variables:**
  - `DB_PATH`: Path to the SQLite database file (default: `game.db`).
  - `PORT`: Port number for the web server (default: `8080`).

### `api.go`
- **Purpose:** Defines all API routes and their handler functions.
- **Responsibilities:**
  - Sets up HTTP endpoints using the Gin framework.
  - Handles core game mechanics, including combat, crafting, alchemy, and quests.
  - Responds to frontend requests with JSON data.

- **Key Components:**
  1. **Routes Setup (`SetupRoutes`):**
     - Maps HTTP endpoints to handler functions.
     - Groups endpoints by functionality:
       - Player: `/player`, `/player/attack`, `/player/use-item`
       - Crafting: `/craft`, `/brew`
       - Game: `/enemies`, `/quests`, `/shop`

  2. **Handler Examples:**
     - **`craftItem`:**
       - Validates player crafting skill and materials.
       - Adds the crafted item to inventory.
       - Awards experience and updates the player level.
     - **`attackEnemy`:**
       - Calculates damage dealt by the player and enemy.
       - Updates player and enemy health.
       - Grants experience and gold upon enemy defeat.

  3. **Data Models:**
     - **Skills:** Represents player abilities in combat, crafting, alchemy, etc.
     - **Enemy:** Models enemy stats, including health, damage, and special abilities.
     - **Inventory:** Categorizes player inventory (weapons, materials, consumables).

- **Environment Variables:**
  - None specific to `api.go`, but integrates with the database (`db`) for fetching game data like crafting recipes and alchemy formulas.


