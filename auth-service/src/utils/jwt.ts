import jwt  from "jsonwebtoken"
import type {  SignOptions } from "jsonwebtoken"
import type { AuthPayload } from "../types/jwtPayload.ts"
type Payload ={
    userId: string
    sessionId: string
}


export const signToken = (payload:Payload,audience:string,signAll=true):{accessToken:string,refreshToken?:string}=>{
    const accesOptions:SignOptions={
        expiresIn:"15m",
        audience: audience
    }
    const refreshOptions:SignOptions={
        expiresIn:"30d",
        audience: audience
    }
    if (signAll) {
         const accessToken= jwt.sign(payload,process.env.JWT_ACCESS_SECRET as string,accesOptions)
        const refreshToken= jwt.sign({sessionId:payload.sessionId},process.env.JWT_REFRESH_SECRET as string,refreshOptions)
        return {accessToken,refreshToken}
    }
    else{
       const accessToken= jwt.sign(payload,process.env.JWT_ACCESS_SECRET as string,accesOptions)
       return {accessToken}
    }
  
}


export const verifyToken = (token:string):AuthPayload=>{
  const payload=  jwt.verify(token,process.env.JWT_REFRESH_SECRET! ) as AuthPayload
  return payload
}