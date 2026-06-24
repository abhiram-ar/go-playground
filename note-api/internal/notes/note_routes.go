package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterRouters(r *gin.Engine, db *mongo.Database) {
	// create repo and handler once at startup

	repo := NewRepo(db)
	h := NewHandler(repo)

	noteGroup := r.Group("/notes")
	{
		noteGroup.POST("", h.CreateNote)
	}
}
