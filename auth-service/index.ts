import express from "express"
import { logger } from "./src/utils/logger.ts"
const PORT = process.env.AUTH_SEERVICE_PORT



const app = express()



 const  startServer =async ()=>{
    try {
       const server= app.listen(PORT,()=>{
        logger.info(`Api gateway running on port ${PORT}`)
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
        console.log("Error starting API Gateway Service:", error)
         logger.error("Error starting API Gateway Service:", error)

    }
}

startServer()

