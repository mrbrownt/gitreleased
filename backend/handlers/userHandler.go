package handlers

import (
	"context"
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
	r.POST("/subcribe", subscribeToRepo)
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

	conf, err := config.Get()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	conf.DB.Where("id = ?", id.(string)).First(&user)
	if user.ID == uuid.Nil {
		// Something really shitty has happened if you hit this!
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func subscribeToRepo(c *gin.Context) {
	conf, err := config.Get()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	repoQuery := c.Query("repo")
	if repoQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no querry value"})
		return
	}

	if slash := strings.Contains(repoQuery, "/"); !slash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repo query not formatted correctly"})
		return
	}

	userID, exists := c.Get("id")
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	repoSlice := strings.Split(repoQuery, "/")

	repo := models.Repository{}
	conf.DB.Where(
		models.Repository{
			Owner: repoSlice[0],
			Name:  repoSlice[1],
		},
	).First(&repo)

	userUUID, err := uuid.FromString(userID.(string))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	var NilUUID = uuid.UUID{}
	if repo.ID == NilUUID {
		repoModel, err := createGithubRepo(userID.(string), repoSlice, conf.DB)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"error": err.Error()},
			)
			logrus.Errorln(err.Error())
			return
		}

		err = createSubscription(userUUID, repoModel.ID, conf.DB)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
	} else {
		err = createSubscription(userUUID, repo.ID, conf.DB)
	}
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.Status(http.StatusCreated)
}

func createSubscription(userID, repoID uuid.UUID, db *gorm.DB) (err error) {
	sub := models.Subscriptions{}
	db.Where(models.Subscriptions{Repo: repoID, User: userID}).First(&sub)
	if sub.ID != 0 {
		return nil
	}

	return db.Create(&models.Subscriptions{
		User: userID,
		Repo: repoID,
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
	k, exists := c.Get("id")
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	conf, _ := config.Get()
	repos := []models.Repository{}

	err := conf.DB.
		Table("repositories").
		Select("*").
		Joins("INNER JOIN subscriptions ON repositories.id = subscriptions.repo AND subscriptions.\"user\" = ?", k.(string)).
		Scan(&repos).Error
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, repos)
}
