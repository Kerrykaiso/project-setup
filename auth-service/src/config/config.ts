import type { Options } from "sequelize";
import dotenv from "dotenv"
dotenv.config()

type Env = "development" | "production" | "test"
type Iconfig = Record<Env, Options & {database: string; password?: string; username?: string }>

// const env: Env = (process.env.NODE_ENV as Env) || "development"
const config : Iconfig = {
  development:{
    username:process.env.DB_USER || "postgres",
    database:process.env.DB || "postgres",
    password: process.env.DB_PASSWORD || "postgres",
    host: process.env.DB_HOST || "127.0.0.1",
    dialect: "postgres",
    port: 5432
  },
  production:{
    username:process.env.DB_USER || "postgres",
    database:process.env.DB || "postgres",
    password: process.env.DB_PASSWORD || "postgres",
    host: process.env.DB_HOST || "127.0.0.1",
    port: 5432
  },
  test:{
     username:process.env.DB_USER || "postgres",
    database:process.env.DB || "postgres",
    password: process.env.DB_PASSWORD || "postgres",
    host: process.env.DB_HOST || "127.0.0.1",
    port: 5432
  }
}

export default config