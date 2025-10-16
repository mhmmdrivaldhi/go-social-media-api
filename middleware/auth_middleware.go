package middleware

import(
	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/utils"
)

func JWTMiddleware(jt utils.JwtToken) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			utils.HandleError(ctx, &utils.UnauthorizedError{Message: "Unauthorized"})
			ctx.Abort()
			return
		}

		userId, err := jt.ValidateToken(tokenString)
		if err != nil {
			utils.HandleError(ctx, &utils.UnauthorizedError{Message: err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userId", userId)
		ctx.Next()
	
	}
}