package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/cheatsnake/emojihub/emojistore"
	"github.com/cheatsnake/emojihub/server"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	port := "4000"
	store := emojistore.New()
	server := server.New(store)
	router := httprouter.New()

	router.GET("/api/all", server.Emojis)
	router.GET("/api/all/category/:category", server.EmojisByCategory)
	router.GET("/api/all/group/:group", server.EmojisByGroup)

	router.GET("/api/random", server.RandomEmoji)
	router.GET("/api/random/category/:category", server.RandomEmojiByCategory)
	router.GET("/api/random/group/:group", server.RandomEmojiByGroup)

	router.GET("/api/categories", server.Categories)
	router.GET("/api/groups", server.Groups)

	router.GET("/api/search", server.Search)
	router.GET("/api/similar/:name", server.Similar)

	fmt.Printf("Server is running on the port %s... \n", port)
	log.Fatal(http.ListenAndServe(":"+port, cors.Default().Handler(router)))
}
