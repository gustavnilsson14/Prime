package main
import (
	"fmt"
	"math/big"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"strconv"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	//router.Static("/web", "./web")
	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	//router.StaticFile("/web/index.html", "./index.html")

	api := router.Group("/api")
	{
		api.GET("/GetPreviousPrime", func(ctx *gin.Context) {
			
			ctx.JSON(200, gin.H{
				"previousPrime": GetAndPrintClosestPrimeNumber(ctx.Request.URL.Query().Get("number")),
			})
		})
	}
	router.Run(":80")
}

func GetAndPrintClosestPrimeNumber(inputNumber string) string{
	i, err := strconv.ParseInt(inputNumber, 10, 64)
	if i < 2 || err != nil{
		return "0"
	}

	var bignum, ok = new(big.Int).SetString(inputNumber, 0)
	if ok == false{
		return "0";
	}
	primeNumber := GetPrimeLowerThan(bignum)
	//primeNumber.SetString(inputNumber, 10) //base 10
	return fmt.Sprintf("%d", primeNumber)
}

func GetPrimeLowerThan(number *big.Int) *big.Int {
    isPrime := false
	for isPrime == false{
		number.Sub(number,big.NewInt(1))
		isPrime = number.ProbablyPrime(1)
	}
    return number
}