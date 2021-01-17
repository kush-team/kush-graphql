package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"kush-graphql/app/models"
	"kush-graphql/helpers"
	"log"
	"net/http"

	"github.com/markbates/going/randx"
)

func (r *mutationResolver) CreateTheme(ctx context.Context, theme models.ThemeInput) (*models.ThemeResponse, error) {
	them := &models.Theme{
		Name:             theme.Name,
		AuthorID:         theme.AuthorID,
		LandingTemplate:  theme.LandingTemplate,
		LandingQuery:     theme.LandingQuery,
		ArticlesTemplate: theme.ArticlesTemplate,
		ArticlesQuery:    theme.ArticlesQuery,
		ArticleTemplate:  theme.ArticleTemplate,
		ArticleQuery:     theme.ArticleQuery,
	}

	if ok, errorString := helpers.ValidateInputs(*them); !ok {
		return &models.ThemeResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	themeCreated, err := r.ThemeService.CreateTheme(them)

	if err != nil {
		log.Println("Theme creation error: ", err)
		return &models.ThemeResponse{
			Message: "Error creating Theme",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.ThemeResponse{
		Message: "Successfully created Theme",
		Status:  http.StatusCreated,
		Data:    themeCreated,
	}, nil
}

func (r *mutationResolver) UpdateTheme(ctx context.Context, id string, theme models.ThemeInput) (*models.ThemeResponse, error) {
	them, err := r.ThemeService.GetThemeByID(id)
	if err != nil {
		log.Println("Error getting the Theme to update: ", err)
		return &models.ThemeResponse{
			Message: "Error getting the Theme",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	them.Name = theme.Name
	them.LandingTemplate = theme.LandingTemplate
	them.LandingQuery = theme.LandingQuery
	them.ArticlesTemplate = theme.ArticlesTemplate
	them.ArticlesQuery = theme.ArticlesQuery
	them.ArticleTemplate = theme.ArticleTemplate
	them.ArticleQuery = theme.ArticleQuery

	if ok, errorString := helpers.ValidateInputs(*them); !ok {
		return &models.ThemeResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	themeUpdated, err := r.ThemeService.UpdateTheme(them)
	if err != nil {
		log.Println("Theme updating error: ", err)
		return &models.ThemeResponse{
			Message: "Error updating Theme",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	for _, observer := range themePublishedChannel {
		observer <- them
	}

	return &models.ThemeResponse{
		Message: "Successfully updated Theme",
		Status:  http.StatusOK,
		Data:    themeUpdated,
	}, nil
}

func (r *mutationResolver) DeleteTheme(ctx context.Context, id string) (*models.ThemeResponse, error) {
	err := r.ThemeService.DeleteTheme(id)
	if err != nil {
		return &models.ThemeResponse{
			Message: "Something went wrong deleting the Theme.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.ThemeResponse{
		Message: "Successfully deleted Theme",
		Status:  http.StatusOK,
	}, nil
}

func (r *queryResolver) GetThemeByID(ctx context.Context, id string) (*models.ThemeResponse, error) {
	theme, err := r.ThemeService.GetThemeByID(id)
	if err != nil {
		log.Println("getting Theme error: ", err)
		return &models.ThemeResponse{
			Message: "Something went wrong getting the Theme.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.ThemeResponse{
		Message: "Successfully retrieved Theme",
		Status:  http.StatusOK,
		Data:    theme,
	}, nil
}

func (r *queryResolver) GetThemeByName(ctx context.Context, name *string) (*models.ThemeResponse, error) {
	theme, err := r.ThemeService.GetThemeByName(name)
	if err != nil {
		log.Println("getting Theme error: ", err)
		return &models.ThemeResponse{
			Message: "Something went wrong getting the Theme.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.ThemeResponse{
		Message: "Successfully retrieved Theme",
		Status:  http.StatusOK,
		Data:    theme,
	}, nil
}

func (r *queryResolver) GetAllThemes(ctx context.Context) (*models.ThemeResponse, error) {
	themes, err := r.ThemeService.GetAllThemes()
	if err != nil {
		log.Println("getting all themes error: ", err)
		return &models.ThemeResponse{
			Message: "Something went wrong getting all themes.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.ThemeResponse{
		Message:  "Successfully retrieved all themes",
		Status:   http.StatusOK,
		DataList: themes,
	}, nil
}

func (r *subscriptionResolver) ThemeChanged(ctx context.Context) (<-chan *models.Theme, error) {
	id := randx.String(8)

	themeEvent := make(chan *models.Theme, 1)
	go func() {
		<-ctx.Done()
	}()

	themePublishedChannel[id] = themeEvent

	return themeEvent, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var themePublishedChannel map[string]chan *models.Theme

func init() {
	themePublishedChannel = map[string]chan *models.Theme{}
}
