import { Sequelize } from "sequelize";
import config from "../config/config.ts";
import { logger } from "../utils/logger.ts";
import User from "./Users.ts"
type Env = "development" | "production" | "test"
const env:Env= process.env.NODE_ENV as Env
const dbConfig =(config as any)[env]


let sequelize: Sequelize | null = null
let Users: ReturnType<typeof User>
 export const connectDB = async()=>{
  if (!sequelize) {
  sequelize = new Sequelize(
  dbConfig.database,
  dbConfig.username,
  dbConfig.password,
  dbConfig
)
    Users = User(sequelize)
  }
 try {
await sequelize.authenticate()
console.log("db connected")
logger.info("db connected")
 } catch (error) {
  logger.error(error)
 }
 return sequelize
}

export {Users}
// // export const models = (sequelize:Sequelize)=>{
// //  return {
// //   User: User(sequelize)
// //  } 
// }
