import type {RequestHandler} from "express"
import AppError from "../utils/appError.ts"
import jwt from "jsonwebtoken"
import type { AuthPayload } from "../types/jwtPayload.ts"

const verifyToken:RequestHandler = (req,res,next) =>{
    const accessToken = req.cookies.accessToken as string | undefined
     if (!accessToken) {
        throw new AppError("No access token",401,"failed")
     }
     const token = accessToken.split('')[1]
      if (!token) {
        throw new AppError("Token missing",401,"failed")
     }
        jwt.verify(token,process.env.JWT_SECRET as string,(err,decoded)=>{
            if (err) {
               if (err.name === "TokenExpiredError") {
                throw new AppError("Token expired",401,"failed")
               }
               if (err.name === "JsonWebTokenError") {
                throw new AppError("Invalid token",401,"failed")
               }
                throw new AppError(err.message,401,"failed")
            }
            req.user = decoded as AuthPayload
            next()
        })
    
}



export default verifyToken