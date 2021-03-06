package messagehandle

import (
	"github.com/canopener/PongPlusPlus-Server/server/connection"
	"github.com/canopener/PongPlusPlus-Server/server/games"
	"github.com/canopener/PongPlusPlus-Server/server/messages"
	"github.com/canopener/serverlog"
)

// LeaveGame handles the case where a player sends a LeaveGame message
func LeaveGame(conn *connection.Conn, allGames map[string]*games.Game, message messages.LeaveGameMessage) {
	serverlog.General("Received LeaveGame message from", conn.Identification())

	if !conn.Registered {
		serverlog.General("Unregistered", conn.Identification(), "called LeaveGame")
		return
	}

	if !conn.InGame {
		serverlog.General(conn.Identification(), "attempted to leave a game but isn't in a game:", allGames[conn.GameID].Name)
		return
	}

	if allGames[conn.GameID].Ready {
		// will be dealt with by game object
		return
	}

	conn.InGame = false
	delete(allGames, conn.GameID)
}
