package main
import ("github.com/gin-gonic/gin"
"net/http"
"strconv"
)
func main(){
	router := gin.Default()
	router.GET("/test",func(context *gin.Context){
		a:=context.DefaultQuery("a","0")
		b:=context.DefaultQuery("b","0")
		a1,_ := strconv.ParseInt(a,0,64)
		b1,_ := strconv.ParseInt(b,0,64)
		res := a1+b1
		context.JSON(
			http.StatusOK,
			gin.H{
				"res":res,
			})
	})
	if err := router.Run(); err != nil {
		panic(err)
	}

}