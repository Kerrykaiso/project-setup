import {Users} from "../models/index.ts"
import bcrypt from "bcrypt"
import AppError from "../utils/AppError.ts"
import {v4 as uuidv4} from "uuid"
import { logger } from "../utils/logger.ts"
import { connectDB } from "../models/index.ts"
import { type Request} from "express"
import { grpcClient } from "../utils/grpc.ts"
const sequelize = connectDB()

type registerData ={
    password:string
    name: string
    email: string
}
type LoginData ={
    password:string
    email: string
}



 export const createAccountService =async(data:registerData)=>{
  try {
    const result = (await sequelize).transaction(async (t)=>{
      const findUser = await Users.findOne({where:{email:data.email}})
    if (findUser) {
      throw new AppError("This email is already in use", 401,"Failed")
    }
    const hashPassword = await bcrypt.hash(data.password,10)
   const id = uuidv4()

    const userData ={
      userId:id,
      email:data.email,
      name: data.name,
      password: hashPassword
    }
    //Add verify email feature later
    const createdUser = await Users.create(userData,{transaction:t})
    logger.info(`New user ${id} created`)
    const { password, ...userWithoutPassword } = createdUser.toJSON()
    return userWithoutPassword
    })
    return result
  } catch (error) {
    logger.error(error)
    throw error
  }
  
}


export const loginAccountService = async(data:LoginData,req:Request)=>{
  try {
    const findUser = await Users.findOne({where: {email:data.email}})

    if (!findUser) {
      logger.warn(`login attempt with ${data.email} at ${req.ip}`)
      throw new AppError("Incorrect email or password",401,"failed")
    }
    const user = findUser.toJSON()
    const correctPassword = await bcrypt.compare(data.password,user.password)
    if (!correctPassword) {
       logger.warn(`login attempt with ${data.email} at ${req.ip}`)
      throw new AppError("Incorrect email or password",401,"failed")

    }
    const requestData ={
      userId:user.userId
    }
    //create session
     const session:any = await new Promise((resolve, reject) => {
      grpcClient.createSession(requestData, (error: any, response: any) => {
        if (error) {
          console.error("grpc client error:", error)
          reject(error)
           throw new AppError("error logging in",401,"failed")
        } else {
          // Check if response and the sessionId field exist
          if (response && response.sessionId) {
            console.log("Session created successfully. sessionId:", response)
            //console.log("Full gRPC response:", response);
            resolve(response)
          } else {
            console.warn("gRPC response is missing sessionId.")
            reject(new Error("Invalid gRPC response for session."))
          }
        }
      });
    })
    const expiresAt = new Date(session.expiresAt.seconds * 1000 + session.expiresAt.nanos/ 1e6)
    console.log(expiresAt.getTime() );
     const { password, ...userWithoutPassword } = user
     return {userWithoutPassword,session}
 
  } catch (error) {
    throw error
  }
}


export const logoutSevice = async(sessionId:string)=>{
  try {
    const id ={
      sessionId
    }
    const deleteSession = await new Promise((resolve,reject)=>{
        grpcClient.deleteSession(id,(error,response)=>{
         if (error) {
            console.error("grpc client error:", error);
            reject(error)
            throw new AppError("error logging out",401,"failed")
         }
         if (response) {
           resolve(response)
         }
       })
    }) 
    return deleteSession
  } catch (error) {
    throw error
  }
}

export const refreshTokenService =async(sessionId:string)=>{
  try {
    const session = {
        sessionId 
    }
    const sessionData:any = await new Promise((resolve,reject)=>{
           grpcClient.getSession(session,(error,response)=>{
              if (error) {
                console.error("grpc client error:", error);
               reject(error)
              throw new AppError("error getting session",401,"failed")
              }
              resolve(response)
           })
    }) 
   
   const expiry = new Date(sessionData.expiresAt.seconds * 1000 + sessionData.expiresAt.nanos/ 1e6)
   const now = Date.now()
   const needRefreshToken = expiry.getTime()-now <= 24*60*60*1000 //less than 24hrs

    if (needRefreshToken) {
      const id ={
        sessionId:sessionData.sessionId
      }
      await new Promise((resolve,reject)=>{
      grpcClient.updateSession(id,(error,response)=>{
          if (error) {
                console.error("grpc client error:", error);
               reject(error)
              throw new AppError("error updating session",401,"failed")
           }
           resolve(response)
      })
     })
    }
  return {needRefreshToken,sessionData}
  } catch (error) {
    throw error

  }
}