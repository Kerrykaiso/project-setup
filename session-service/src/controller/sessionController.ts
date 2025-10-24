import { connectDB } from "../models/index.ts";
import { Sessions } from "../models/index.ts";
import grpc from "@grpc/grpc-js";
import { v4 as uuidv4 } from "uuid";
import type { ServerUnaryCall, sendUnaryData } from "@grpc/grpc-js";
import type { SessionRequest } from "../../proto/session/SessionRequest.ts";
import type { SessionResponse } from "../../proto/session/SessionResponse.ts";
import type { deleteSessionRequest } from "../../proto/session/deleteSessionRequest.ts";
import type { deleteSessionResponse } from "../../proto/session/deleteSessionResponse.ts";
import type { getSessionRequest } from "../../proto/session/getSessionRequest.ts";
import type { getSessionResponse } from "../../proto/session/getSessionResponse.ts";
import type { updateSessionRequest } from "../../proto/session/updateSessionRequest.ts";
import type { updateSessionResponse } from "../../proto/session/updateSessionResponse.ts";
import { logger } from "../utils/logger.ts";

const sequelize = await connectDB();
const thirtyDays = new Date(Date.now() + 30 * 24 * 60 * 1000);
export const createSession = async (
  call: ServerUnaryCall<SessionRequest, SessionResponse>,
  callback: sendUnaryData<SessionResponse>
) => {
  const sessionId = uuidv4();
  try {
    const result = sequelize.transaction(async (t) => {
      const sessionData = {
        userId: call.request.userId as string,
        sessionId: sessionId,
        expiresAt: thirtyDays, //30 days
      };
      const createdSession = await Sessions.create(sessionData, {
        transaction: t,
      });
      return createdSession;
    });
    const data = (await result).toJSON();
    const id = data.sessionId;
    const expiresAt = data.expiresAt;
    const expiresAtDate = {
      seconds: Math.floor(expiresAt.getTime() / 1000),
      nanos: (expiresAt.getTime() % 1000) * 1e6,
    };
    callback(null, { sessionId: id, expiresAt: expiresAtDate });
    logger.info("new session created");
  } catch (error) {
    logger.error(error);
    callback({ code: grpc.status.NOT_FOUND, message: "User not found" });
  }
};

export const deleteSession = async (
  call: ServerUnaryCall<deleteSessionRequest, deleteSessionResponse>,
  callback: sendUnaryData<deleteSessionResponse>
) => {
  try {
    sequelize.transaction(async (t) => {
      const findSession = await Sessions.findByPk(call.request.sessionId);
      if (!findSession) {
        callback({
          code: grpc.status.NOT_FOUND,
          message: "sessionId not found",
        });
        return;
      }
      await findSession.destroy({ transaction: t });
    });
    const success = true;
    callback(null, { success });
  } catch (error) {
    callback({
      code: grpc.status.NOT_FOUND,
      message: "Error deleting session",
    });
    logger.error(error);
  }
};

export const getSession = async (
  call: ServerUnaryCall<getSessionRequest, getSessionResponse>,
  callback: sendUnaryData<getSessionResponse>
) => {
  try {
    const findSession = await Sessions.findByPk(call.request.sessionId);
    if (!findSession) {
      callback({ code: grpc.status.NOT_FOUND, message: "sessionId not found" });
      return;
    }
    const { expiresAt, userId, sessionId } = findSession.toJSON();
    const expiresAtDate = {
      seconds: Math.floor(expiresAt.getTime() / 1000),
      nanos: (expiresAt.getTime() % 1000) * 1e6,
    };
    callback(null, { sessionId, expiresAt: expiresAtDate, userId });
  } catch (error) {
    callback({
      code: grpc.status.NOT_FOUND,
      message: "Error deleting session",
    });
    logger.error(error);
  }
};

export const updateSession = async (
  call: ServerUnaryCall<updateSessionRequest, updateSessionResponse>,
  callback: sendUnaryData<updateSessionResponse>
) => {
  try {
    sequelize.transaction(async (t) => {
      const updated = await Sessions.update(
        { expiresAt: thirtyDays },
        { where: { sessionId: call.request.sessionId }, transaction: t }
      );
      if (!updated) {
        callback({
          code: grpc.status.NOT_FOUND,
          message: "Error updating session",
        });
        logger.error("error updating session");
      }
      callback(null, { success: true });
    });
  } catch (error) {
    callback({
      code: grpc.status.NOT_FOUND,
      message: "Error updating session",
    });
    logger.error(error);
  }
};
