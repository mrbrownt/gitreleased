package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/shurcooL/githubv4"
	"github.com/sirupsen/logrus"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"
	"golang.org/x/oauth2"
)

// UserHandler propigates routes for users and whatnot
func UserHandler(r *gin.RouterGroup) {
	r.Use(authMiddleware())
	r.GET("/", getUser)
	r.POST("/subscribe", subscribeToRepo)
	r.GET("/subscriptions", getSubscriptions)
}

func getUser(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		// Not sure if this is what we want to do here
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user := models.User{}
	db := config.Get().DB

	db.Where("id = ?", id.(uuid.UUID).String()).First(&user)
	if user.ID == uuid.Nil {
		// Something really shitty has happened if you hit this!
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func subscribeToRepo(c *gin.Context) {
	conf := config.Get()

	repoQuery := c.Query("repo")
	if repoQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no querry value"})
		return
	}

	id, exists := c.Get("id")
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	userID, ok := id.(uuid.UUID)
	if !ok {
		internalServerError(c, errors.New("Invalid user ID"))
	}

	tx := conf.DB.Begin()
	user := models.User{ID: userID}
	if !user.Valid(tx) {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "unknown user"},
		)
	}

	owner, name, err := parseRepoQuery(repoQuery)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	repo := models.Repository{
		Owner: owner,
		Name:  name,
	}

	if repo.Exists(tx) {
		subscription := models.Subscriptions{RepoID: repo.ID, UserID: user.ID}
		err := subscription.Create(tx)
		if err != nil {
			tx.Rollback()
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.Status(http.StatusCreated)
		return
	}

	repoSlice := []string{owner, name}
	if repo.ID == uuid.Nil {
		repoModel, err := createGithubRepo(userID.String(), repoSlice, tx)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": err.Error()},
			)
			logrus.Errorln(err.Error())
			return
		}

		err = createSubscription(userID, repoModel.ID, tx)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
	} else {
		err = createSubscription(userID, repo.ID, tx)
	}
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	err = tx.Commit().Error
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.Status(http.StatusCreated)
}

func parseRepoQuery(query string) (owner string, name string, err error) {
	if !strings.Contains(query, "/") {
		return "", "", errors.New("repo string should be formatted \"owner/name\"")
	}

	splitQuery := strings.Split(query, "/")
	owner = splitQuery[0]
	name = splitQuery[1]
	return owner, name, nil
}

func createSubscription(userID, repoID uuid.UUID, db *gorm.DB) (err error) {
	sub := models.Subscriptions{}
	err = db.Where(models.Subscriptions{RepoID: repoID, UserID: userID}).First(&sub).Error
	if !gorm.IsRecordNotFoundError(err) {
		return err
	}

	if sub.RepoID != uuid.Nil {
		return nil
	}

	return db.Create(&models.Subscriptions{
		UserID: userID,
		RepoID: repoID,
	}).Error
}

func createGithubRepo(userID string, repo []string, db *gorm.DB) (r models.Repository, err error) {
	user := models.User{}
	db.Where("id = ?", userID).First(&user)
	client := setupGithubClient(user.AccessToken)

	variables := map[string]interface{}{
		"owner": githubv4.String(repo[0]),
		"name":  githubv4.String(repo[1]),
	}

	var query struct {
		Repository struct {
			Description string
			URL         githubv4.URI
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	err = client.Query(context.Background(), &query, variables)
	if err != nil {
		return r, err
	}

	r.Owner = repo[0]
	r.Name = repo[1]
	r.Description = query.Repository.Description
	r.URL = query.Repository.URL.String()

	err = db.Create(&r).Error
	if err != nil {
		return r, err
	}

	return r, err
}

func setupGithubClient(token string) (client *githubv4.Client) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	return githubv4.NewClient(httpClient)
}

func getSubscriptions(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	userID, ok := id.(uuid.UUID)
	if !ok {
		internalServerError(c, errors.New("user id"))
	}

	conf := config.Get()
	repos := []models.Repository{}

	err := conf.DB.
		Table("repositories").
		Select("*").
		Joins("INNER JOIN subscriptions ON repositories.id = subscriptions.repo_id AND subscriptions.user_id = ?", userID.String()).
		Scan(&repos).
		Error
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, repos)
}
