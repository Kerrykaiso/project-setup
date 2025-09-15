import { Sequelize } from "sequelize";
import config from "../config/config.ts";
import { logger } from "../utils/logger.ts";
import Session from "./Sessions.ts"
type Env = "development" | "production" | "test"
const env:Env= process.env.NODE_ENV as Env
const dbConfig =(config as any)[env]


let sequelize: Sequelize | null = null
let Sessions: ReturnType<typeof Session>
 export const connectDB = async()=>{
  if (!sequelize) {
  sequelize = new Sequelize(
  dbConfig.database,
  dbConfig.username,
  dbConfig.password,
  dbConfig
)
    Sessions = Session(sequelize)
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

export {Sessions}