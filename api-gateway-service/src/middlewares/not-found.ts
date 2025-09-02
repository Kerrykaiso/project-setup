import type { Request,Response,NextFunction } from "express"
export const notfound = (req:Request,res:Response,next:NextFunction)=>{
    res.status(404).json({message:"endpoint not found",success:"failed"})
}