package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"crud/cache"
)

var redisCache = cache.NewRedisCache("localhost:6379", 0, 1)

func main() {
	router := gin.Default()

	router.POST("/movies", func(ctx *gin.Context){
		var movie cache.Movie

		if err := ctx.ShouldBind(&movie); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		res, err := redisCache.CreateMovie(ctx, &movie)
		if err != nil {
           ctx.JSON(http.StatusBadRequest, gin.H{
               "error": err.Error(),
           })
		   return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"movie" : res,
		})
	})


	router.GET("/movies/:id", func(ctx *gin.Context){
		id := ctx.Param("id")
		movie, err := redisCache.GetMovie(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "movie not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"movie": movie,
		})
	})


	router.GET("/movies", func(ctx *gin.Context){
		movies, err := redisCache.GetMovies(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"movies": movies,
		})
	})


	router.PUT("/movies/:id", func(ctx *gin.Context){
		id := ctx.Param("id")

		res, err := redisCache.GetMovie(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		var movie cache.Movie

		if err := ctx.ShouldBind(&movie); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		res.Title = movie.Title
		res.Description = movie.Description
		res, err = redisCache.UpdateMovie(ctx, res)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"movie": res,
		})
	})

	router.DELETE("/movies/:id", func(ctx *gin.Context){
		id := ctx.Param("id")
		err := redisCache.DeleteMovie(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "movie deleted sucessufully",
		})
	})

	fmt.Println(router.Run(":5000"))
}