import {DataTypes,Sequelize,Model} from "sequelize"
//import type {ModelDefined} from "sequelize"

interface Session {
    userId: string
    sessionId:string
    expiresAt:Date
}
export interface SessionInstance extends Model<Session> {} 

export default function Session(sequelize:Sequelize) {
  const Session = sequelize.define<SessionInstance>(
//  const Session: ModelDefined<Session, SessionInstance> = sequelize.define(
     "Session",{
        userId:{
            type:DataTypes.STRING,
           
        },
        sessionId:{
            type:DataTypes.STRING,
            allowNull: false,
            primaryKey:true
        },
       expiresAt:{
        type:DataTypes.DATE,
        allowNull:false
       }
       
     }
  )
  return Session
}