import express from "express"
import dotenv from "dotenv"
dotenv.config()
import cookieParser from "cookie-parser"
import { logger } from "./utils/logger.ts"
import { connectDB,  } from "./models/index.ts"
import authRoute from "./routes/auth-route.ts"
import errorHandler from "./middlewares/erroHandler.ts"
import { notfound } from "./middlewares/notFound.ts"
const PORT = process.env.AUTH_SERVICE_PORT



const app = express()
app.use(express.json())
app.use(cookieParser())
app.use("/api/auth",authRoute)
app.use(notfound)
app.use(errorHandler)



 const  startServer =async ()=>{
    try {
       await connectDB()
    
       const server= app.listen(PORT,()=>{
        logger.info(`Auth service running on  ${PORT}`)
      })

       const shutdownSignal = ["SIGTERM","SIGINT"]

       shutdownSignal.forEach((signal)=>{
         process.on(signal,async()=>{
            console.log("shutting down server")
            logger.info("shutting down server")
             try {
            server.close(()=>{
                console.log("Server shut down")
                logger.info("Server shut down")

            })
        } catch (error) {
            console.log("Error shutting down", error)
              logger.error("Error shutting down", error)
        }
         })
       })
    } catch (error) {
        console.log("Error starting auth Service:", error)
         logger.error("Error starting auth Service:", error)

    }
}

startServer()

