import grpc from "@grpc/grpc-js"
import protoLoader from "@grpc/proto-loader"
import path from "path"
import dotenv from "dotenv"
dotenv.config()

import { fileURLToPath } from 'url';


const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);


import type { ProtoGrpcType } from "../proto/session.ts"
import { connectDB } from "./models/index.ts"
import { createSession ,deleteSession, getSession, updateSession} from "./controller/sessionController.ts"


const packageDefinition = protoLoader.loadSync(path.join(__dirname, "..", "proto", "session.proto"))
const grpcObj =(grpc.loadPackageDefinition(packageDefinition)) as unknown as ProtoGrpcType
const sessionProto = grpcObj.session

const startServer = async() =>{
try {
 await connectDB()
 const server = new grpc.Server()
const port ="127.0.0.1:5003"

server.addService(sessionProto.SessionService.service,{createSession:createSession,deleteSession:deleteSession,getSession:getSession,
    updateSession:updateSession
})
server.bindAsync(port,grpc.ServerCredentials.createInsecure(),(error,port)=>{
if (error) {
   console.log("failed to start grpc server", error) 
   return
}
console.log(`server running on ${port}`)
})
} catch (error) {
    console.log(error)
}
}
startServer()