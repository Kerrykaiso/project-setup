import type {ErrorRequestHandler} from "express"
import AppError from "../utils/AppError.ts"
import { logger } from "../utils/logger.ts"
import {z} from "zod"

const errorHandler:ErrorRequestHandler = (err,req,res,next)=>{


  if (err instanceof z.ZodError) {
     err.issues.map((error)=>{
     return res.status(400).json({
      error : error.path.join("."),
      message: error.message
     })
    })

  }
  if (err instanceof AppError) {
    logger.error(`ehandler ${err.message}`)
   return res.status(err.statusCode).json({
        message:err.message,
        status:err.status,
        data: null
    })
  }
  logger.error(err.message)
 return res.status(500).json({status:"failed",message: "Internal server error" }
  )
}

export default errorHandler