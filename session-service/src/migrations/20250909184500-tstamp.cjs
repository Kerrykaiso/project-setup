'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface, Sequelize) {
   await queryInterface.addColumn("Sessions","createdAt",{
      type:Sequelize.DataTypes.DATE
    
   }),
   await queryInterface.addColumn("Sessions","updatedAt",{
      type:Sequelize.DataTypes.DATE
    
   })
  },

  async down (queryInterface, Sequelize) {
   
    await queryInterface.removeColumn("Sessions","createdAt")
    await queryInterface.removeColumn("Sessions","updatedAt")
  }
};
