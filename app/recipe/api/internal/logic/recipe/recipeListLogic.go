package recipe

import (
	"context"

	"CookingMaster_Backend/app/recipe/api/internal/svc"
	"CookingMaster_Backend/app/recipe/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecipeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecipeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecipeListLogic {
	return &RecipeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetIngredientsByRecipe retrieves all ingredients for a given recipe.
func (m *RecipeDetailsModel) GetIngredientsByRecipe(ctx context.Context, recipeID int64) ([]Ingredient, error) {
	var ingredients []Ingredient
	query := "SELECT id, recipe_id, name, quantity, created_at FROM ingredients WHERE recipe_id = ?"
	err := m.conn.QueryRowsCtx(ctx, &ingredients, query, recipeID)
	return ingredients, err
}
}
