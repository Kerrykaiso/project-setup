 import { redisConnection } from "../utils/redis.ts"
 import { RateLimiterRedis} from "rate-limiter-flexible"
 import rateLimit from "express-rate-limit";
import RedisStore from "rate-limit-redis";
import { logger } from "../utils/logger.ts";


const redisInstance = redisConnection()


export const rateLimiter  = new RateLimiterRedis({
  storeClient:redisInstance,
  keyPrefix: "middleware",
  points: 10,
  duration:1
 })


  export const endpointLimiter = rateLimit({
  windowMs: 10*60*1000,
  max:50,
  standardHeaders:true,
  legacyHeaders:false,
  handler:(req,res)=>{
    logger.warn(`exccessive request from IP :${req.ip}`)
    res.status(429).json({success:"failed",message:"Too many request"})
  },
  store: new RedisStore({
      sendCommand: (command: string, ...args: string[])=>redisInstance.sendCommand([command, ...args]) 
  })

  })