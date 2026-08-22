package notes

import (
	"fmt"
	"net/http" 
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote(c *gin.Context){
	var req CreateNoteRequest;
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Json",
		})
		return ;
	}
	now := time.Now();
	note := Note{
		ID : primitive.NewObjectID(),
		Title : req.Title,
		Content : req.Content,
		Pinned:  req.Pinned,

		CreatedAt : now,
		UpdatedAt : now,
	}
	fmt.Printf("Creating Note: %+v\n", note)
	created , err := h.repo.Create(c.Request.Context(),note)
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, created)
}


func (h *Handler) ListNotes(c *gin.Context){
	notes , err := h.repo.List(c.Request.Context())
	// err =  ;
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error: Custom Messsage @ ": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, notes)
}

func (h *Handler) GetNoteByID(c *gin.Context){
	id, err := primitive.ObjectIDFromHex(c.Param("id")) 
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	note , err :=  h.repo.GetByID(c.Request.Context(),id)
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, note)
}

func (h *Handler) UpdateNoteByID(c *gin.Context){
	id, err := primitive.ObjectIDFromHex(c.Param("id")) 
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	var req UpdateNoteRequest;
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ 
			"error": err.Error(),
		})
		return
	} 
	updated , err := h.repo.UpdateByID(c.Request.Context(),id,req)
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteNoteByID(c *gin.Context){
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}	
	deleted, err := h.repo.DeleteByID(c.Request.Context(), id)
	if err != nil {
		fmt.Println("Handler received error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
}