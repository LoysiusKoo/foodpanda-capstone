package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/loysiuskoo/foodpanda-capstone/database/sqlc"
)

// Server will serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	server.setupRouter()
	return server
}

// setRouter setup routing
func (server *Server) setupRouter() {
	router := gin.Default()

	// add routes to router
	router.POST("/users", server.createUser)
	router.POST("/playlists", server.createPlaylist)
	router.GET("/playlists/:playlistid", server.getPlaylist)
	router.GET("/playlists", server.listPlaylist)
	router.GET("/playlists/current/:playlistid", server.getPlaylistCurrent)
	router.GET("/playlists-latest", server.getPlaylistLatest)
	router.GET("/playlists-category", server.getPlaylistCategory)
	router.GET("/publicplaylists", server.getPublicPlaylist)
	router.GET("/userplaylists/:userid", server.getUserPlaylist)

	router.POST("/playlist/:userid/:playlistid", server.updateUserPlaylistStatus) // This post request is just to update the status to Pending for testing purpose.
	router.PUT("/playlist/:userid/:playlistid", server.updateUserPlaylistStatus)

	router.GET("/playlists-latestv2", server.getPlaylistLatestV2)
	router.GET("/playlists-user", server.getPlaylistUser)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
