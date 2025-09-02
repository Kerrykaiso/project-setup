import { AuthPayload } from "./jwtPayload.ts";

declare module "express-serve-static-core"{
    interface Request {
        user?: AuthPayload
    }
}