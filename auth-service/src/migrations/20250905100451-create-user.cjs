'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface, Sequelize) {
    await queryInterface.createTable("Users",{
      userId:{
        primaryKey:true,
        type: Sequelize.DataTypes.STRING,
        unique:true
      },
      name:{
        type: Sequelize.DataTypes.STRING,
        allowNull:false
      },
      email:{
        type: Sequelize.DataTypes.STRING, 
        allowNull:false
      },
      password:{
         type: Sequelize.DataTypes.STRING,
         allowNull:false
      },
      createdAt:{
          type: Sequelize.DataTypes.DATE
      },
      updatedAt:{
           type: Sequelize.DataTypes.DATE
      }

    })
    
  },

  async down (queryInterface, Sequelize) {
   await queryInterface.dropTable("Users")
  }
};
