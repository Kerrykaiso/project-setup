import express from "express";
import type { Request, Response, NextFunction } from "express";
import dotenv from "dotenv";
dotenv.config();
import errorHandler from "./middlewares/errorHandler.ts";
import { globalLimiter } from "./middlewares/rateLimit.ts";
import proxy from "express-http-proxy";
import cookieParser from "cookie-parser";
import { redisConnection } from "./utils/redis.ts";
import { logger } from "./utils/logger.ts";
import AppError from "./utils/appError.ts";
import { notfound } from "./middlewares/not-found.ts";
const app = express();
const PORT = process.env.API_GATEWAY_PORT;

app.use(globalLimiter);

// app.use(endpointLimiter)
app.use(express.json());
app.use(cookieParser());

function createProxyOptions(serviceName: string) {
  return {
    proxyReqPathResolver: (req: Request) => {
      return req.originalUrl.replace(/^\/v1/, "/api");
    },

    proxyErrorHandler: (err: any, res: Response, next: NextFunction) => {
      logger.error(`${serviceName} error: ${err.message}`);

      if (err && err.code === "ECONNREFUSED") {
        logger.error(`proxy server error ${err.message}`);
        return next(
          new AppError(`${serviceName} is unavailable`, 503, "failed"),
        );
      }
      if (err && err.code === "ECONNRESET") {
        console.log("err", err);
        return next(new AppError("Internal server error", 500, "failed"));
      }
      return next(err);
    },
  };
}

//service health status
app.get("/api-gateway-health", (_, res) => {
  res
    .status(200)
    .json({ message: "api-gateway-service running", success: true });
});

app.use(
  "/v1/auth",
  proxy(process.env.AUTH_PATH_URL as string, {
    ...createProxyOptions("Users"),
    proxyReqOptDecorator: (proxyReqOpts, srcreq) => {
      proxyReqOpts.timeout = 5000;
      return proxyReqOpts;
    },
    userResDecorator: (proxyRes, proxyResData, userReq, userRes) => {
      return proxyResData;
    },
  }),
);

app.use(
  "/v1",
  proxy(process.env.DESIGNER_PATH_URL as string, {
    ...createProxyOptions("Designer"),
    proxyReqOptDecorator: (proxyReqOpts, srcreq) => {
      proxyReqOpts.timeout = 5000;
      return proxyReqOpts;
    },
    userResDecorator: (proxyRes, proxyResData, userReq, userRes) => {
      return proxyResData;
    },
  }),
);

app.use(notfound);
app.use(errorHandler);

const startServer = async () => {
  try {
    redisConnection();
    const server = app.listen(PORT, () => {
      logger.info(`Api gateway running on port ${PORT}`);
    });

    const shutdownSignal = ["SIGTERM", "SIGINT"];

    shutdownSignal.forEach((signal) => {
      process.on(signal, async () => {
        console.log("shutting down server");
        logger.info("shutting down server");
        try {
          server.close(() => {
            console.log("Server shut down");
            logger.info("Server shut down");
          });
        } catch (error) {
          console.log("Error shutting down", error);
          logger.error("Error shutting down", error);
        }
      });
    });
  } catch (error) {
    console.log("Error starting API Gateway Service:", error);
    logger.error("Error starting API Gateway Service:", error);
  }
};

startServer();
