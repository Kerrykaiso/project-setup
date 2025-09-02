import type {ErrorRequestHandler} from "express"
import AppError from "../utils/appError.ts"
import { logger } from "../utils/logger.ts"

const errorHandler:ErrorRequestHandler = (err,req,res)=>{

  if (err instanceof AppError) {
    res.status(err.statusCode).json({
        message:err.message,
        status:err.status,
        data: null
    })
    logger.error(err.message)
  }
  logger.error(err.message)
  res.status(500).json({status:"failed",message: "Internal server error" }
  )
}

export default errorHandler