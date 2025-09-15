'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface, Sequelize) {
    await queryInterface.renameColumn("Sessions", "userid","userId")
  },

  async down (queryInterface, Sequelize) {
    await queryInterface.renameColumn("Sessions", "userId","userid")

  }
};
