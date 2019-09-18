package mathcomputation

import (
	"github.com/spf13/viper"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/iris-contrib/middleware/cors"
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

func SetupRoutes(ops *MathOps, monitor *APIMonitor){
	app := iris.New()
	var crs context.Handler
	if viper.GetBool("AllowCrossOrigin") {
		crs = cors.New(cors.Options{
			AllowedOrigins: []string{"*"},	//allows everything, use that to change the hosts.
			AllowedMethods: []string{"GET", "POST", "PUT"},
			AllowedHeaders: []string{"*"},
			AllowCredentials: true,
		})
	}

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	//TODO: write separate helper function to register all api endpoints
	//Register API endpoint for usage montioring
	endpointList := make([]string,0)
	endpointList = append(endpointList, "add")
	monitor.RegisterAPI(endpointList)
	//Running usage Montoring
	monitor.RunMonitoring()

	v1 := app.Party("/api/v1/math", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/add", func (ctx iris.Context){
			//Update Usage Count
			monitor.AddUsageCount("add")
			
			//Assuming if numbers is not provided in api then by default number value is 0
			first := ctx.URLParamInt64Default("first", 0)			//TODO: Pass query param keys from spearte variable.
			second := ctx.URLParamInt64Default("second", 0)
			log.Println("Add Two Number API Called for : ", first, second)
			res := ops.AddTwoNumbers(first, second)
			response , err := json.Marshal(res)
			//_ err := ctx.JSON(res)
			
			if err != nil {
				ctx.Values().Set("message", "Computation Failed")
				ctx.StatusCode(http.StatusInternalServerError)
			} else{
				ctx.StatusCode(http.StatusOK)
				ctx.Text(string(response))
				log.Println("Response: ", string(response))
			}
		})

		v1.Get("/usage", func (ctx iris.Context){
			epoint := ctx.URLParam("endpoint")

			count, err := monitor.GetUsageCount(epoint)

			if err != nil {
				ctx.Values().Set("message", err)
				ctx.StatusCode(http.StatusInternalServerError)
			}else {
				response , _ := json.Marshal(count)
				ctx.StatusCode(http.StatusOK)
				ctx.Text(string(response))
			}

		})

	}

	//Serving
	port := fmt.Sprintf(":%s", viper.GetString("MathComputationServer.Port"))
	app.Run(iris.Addr(port), iris.WithCharset(viper.GetString("CharSet")))
}

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("<b>404 Page Not Found<b>")
}