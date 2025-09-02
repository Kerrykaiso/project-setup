import  {createClient} from "redis"
import type  {RedisClientType} from "redis"
import AppError from "./appError.ts"

   let connectRedis :RedisClientType
 export const redisConnection = ():RedisClientType => {

    if(!connectRedis){
        connectRedis =   createClient({
            url: "redis://127.0.0.1:6379"
        })
        connectRedis.on("error",(err)=>{
            console.log("Redis connection error")
        })
        connectRedis.on("connect",()=>{
            console.log("Connected to Redis")
        })
        connectRedis.on("close",()=>{
          console.log("disconnected from Redis")
        }
    )
      connectRedis.on("end",()=>{
          console.log("ended from Redis")
        })
        process.on("SIGINT",async()=>{
            console.log("SIGINT closing redis")
            await connectRedis.quit()
        })
          process.on("SIGTERM",async()=>{
            console.log("SIGTERM closing redis")
            await connectRedis.quit()
        })
        connectRedis.connect().catch((err)=>new AppError(err.message,400,"failed"))
    }
   return connectRedis
}
