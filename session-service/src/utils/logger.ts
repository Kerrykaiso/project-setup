import dotenv from "dotenv"
dotenv.config()
import winston, {format,transports} from "winston";
const {combine} = format
const logFormat = format.printf(({level,message,timestamp})=>{
   return ` [${level.toUpperCase()}]: ${message} ${timestamp} `
})
 export const logger = winston.createLogger({
    level: process.env.NODE_ENV ==="development" ? "debug":"info",
    defaultMeta: {service:"api-gateway-service"},
    format:combine(
        format.timestamp(),
        process.env.NODE_ENV ==="production"? format.json() : logFormat
    ),
    transports:[
        new transports.Console({format:combine()}),
        new transports.File({filename:"src/logs/error.log", level:"error"})
    ]
})