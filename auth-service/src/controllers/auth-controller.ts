import type { NextFunction, Request, Response } from "express";
import { loginSchema, registerSchema } from "../utils/auth-schema.ts";
import {
  createAccountService,
  loginAccountService,
  logoutSevice,
  refreshTokenService,
} from "../services/auth-service.ts";
import AppError from "../utils/AppError.ts";
import { logger } from "../utils/logger.ts";
import { signToken, verifyToken } from "../utils/jwt.ts";
import { clearCookies, setAuthCookie } from "../utils/cookies.ts";
import type { JwtPayload } from "jsonwebtoken";
import type { AuthPayload } from "../types/jwtPayload.ts";

export const registerController = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  try {
    const request = registerSchema.parse({ ...req.body });
    const authResponse = await createAccountService(request);
    if (!authResponse) {
      throw new AppError("Error registering your account", 401, "failed");
    }
    res.status(201).json({
      message: "Account created successfully",
      status: "success",
      data: authResponse,
    });
  } catch (error) {
    next(error);
    logger.error(error);
  }
};

export const loginController = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  try {
    const request = loginSchema.parse({ ...req.body });
    const loginData = await loginAccountService(request, req);
    const { accessToken, refreshToken } = signToken(
      {
        userId: loginData.userWithoutPassword.userId,
        sessionId: loginData.session.sessionId,
      },
      "users"
    );
    await setAuthCookie(res, accessToken, refreshToken);
    res.status(201).json({
      message: "Account logged in successfully",
      status: "success",
      data: loginData.userWithoutPassword,
    });
  } catch (error) {
    next(error);
    logger.error(`controller ${error}`);
  }
};

export const logOutController = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  try {
    const sessionId = "e89837c1-27af-4514-ace7-2c4b31dc73fa";
    if (!sessionId) {
      throw new AppError("User not logged in", 401, "failed");
    }
    const logOut = await logoutSevice(sessionId);
    if (!logOut) {
      throw new AppError("failed to logout", 401, "failed");
    }
    clearCookies(res);
    res.status(200).json("Logout successful");
  } catch (error) {
    next(error);
    logger.error(`controller ${error}`);
  }
};

export const refreshTokenController = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  try {
    console.log(req.url);
    const refreshToken = req.cookies.refreshToken;
    console.log("reftreshhh", refreshToken);
    if (!refreshToken) {
      throw new AppError("No refresh token", 401, "failed");
    }
    const payload = verifyToken(refreshToken) as AuthPayload & JwtPayload;
    const { needRefreshToken, sessionData } = await refreshTokenService(
      payload.sessionId
    );
    if (needRefreshToken) {
      const { accessToken, refreshToken } = signToken(
        { userId: sessionData.userId, sessionId: sessionData.sessionId },
        "users"
      );
      await setAuthCookie(res, accessToken, refreshToken);
    }
    const { accessToken } = signToken(
      { userId: sessionData.userId, sessionId: sessionData.sessionId },
      "users",
      needRefreshToken
    );
    await setAuthCookie(res, accessToken, refreshToken);
    res.status(200).json("Token refreshed successfully");
  } catch (error) {
    next(error);
    logger.error(`controller ${error}`);
  }
};
