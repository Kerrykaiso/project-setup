'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface, Sequelize) {
    /**
     * Add altering commands here.
     *
     * Example:
     * await queryInterface.createTable('users', { id: Sequelize.INTEGER });
     */
    await queryInterface.createTable("Sessions",{
      sessionId:{
        type: Sequelize.DataTypes.STRING,
        unique:true,
        primaryKey:true,
        allowNull:false
      },
      userid:{
        type: Sequelize.DataTypes.STRING,
        allowNull:false
      },
      expiresAt:{
        type:Sequelize.DataTypes.DATE,
        allowNull:false
      }
    })
  },

  async down (queryInterface, Sequelize) {
      await queryInterface.dropTable('Sessions');
  }
};
