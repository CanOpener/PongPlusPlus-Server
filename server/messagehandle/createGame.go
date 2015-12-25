package messagehandle

import (
	"github.com/canopener/PongPlusPlus-Server/server/connection"
	"github.com/canopener/PongPlusPlus-Server/server/games"
	"github.com/canopener/PongPlusPlus-Server/server/messages"
	"github.com/canopener/serverlog"
)

// CreateGame handles the case where a player sends a CreateGame message
func CreateGame(conn *connection.Connection, allGames map[string]*games.Game, message messages.CreateGameMessage) {
	serverlog.General("Received CreateGame message from conn:", conn.Alias)

	if conn.InGame {
		serverlog.General("conn:", conn.Alias, " tried to create a new game but is already in game:", conn.InGameID)
		denied := messages.NewCreateGameDeniedMessage(message.GameName, "You are already in a game")
		conn.Write(denied.Bytes())
	}

	serverlog.General("Creating game:", message.GameName, "by conn:", conn.Alias)
	game := games.NewGame(conn, message.GameName)
	serverlog.General("conn:", conn.Alias, "setting InGame to true and Game to the game:", game.Name)
	conn.InGame = true
	conn.Game = game
	serverlog.General("Attatching game:", game.Name, "to games list")
	allGames[game.ID] = game

	approved := messages.NewCreateGameApprovedMessage(game.ID, game.Name)
	conn.Write(approved.Bytes())
}