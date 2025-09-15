 import { redisConnection } from "../utils/redis.ts"
 import rateLimit from "express-rate-limit";
import RedisStore from "rate-limit-redis";
import { logger } from "../utils/logger.ts";


const redisInstance = redisConnection()

console.log("redisInstance",redisInstance)
export const globalLimiter = rateLimit({
  windowMs: 1 * 60 * 1000, // 1 minute window
  max: 20, // 100 requests per IP
  standardHeaders: true,
  legacyHeaders: false,
  handler: (req, res, next) => {
    // 🔹 Log IP when limit exceeded
    logger.warn(`Rate limit exceeded by IP: ${req.ip}`);

    res.status(429).json({
      status: "failed",
      message: "Too many requests, please try again later.",
    });
  },
  store: new RedisStore({
    sendCommand: (...args: string[]) => redisInstance.sendCommand(args),
  }),
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