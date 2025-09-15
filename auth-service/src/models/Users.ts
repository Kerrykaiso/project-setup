import {DataTypes,Sequelize,Model} from "sequelize"
//import type {ModelDefined} from "sequelize"

interface User {
    userId: string
    name: string
    email: string
    password: string
    createdAt?: Date
    updatedAt?: Date
}
export interface UserInstance extends Model<User> {} 

export default function User(sequelize:Sequelize) {
  const User = sequelize.define<UserInstance>(
//  const User: ModelDefined<User, UserInstance> = sequelize.define(
     "User",{
        userId:{
            type:DataTypes.STRING,
            primaryKey:true
        },
        name:{
            type:DataTypes.STRING,
            allowNull: false
        },
        password:{
            type:DataTypes.STRING,
            allowNull:false
        },
        email:{
            type:DataTypes.STRING,
            unique: true
        }
     }
  )
  return User
}