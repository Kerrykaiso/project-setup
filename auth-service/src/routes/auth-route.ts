import router from "express"
import { loginController, logOutController, refreshTokenController, registerController } from "../controllers/auth-controller.ts"
export const authRoute = router.Router()

authRoute.post("/register",registerController)
authRoute.post("/login",loginController)
authRoute.delete("/logout",logOutController)
authRoute.get("/refresh-token", refreshTokenController)


export default authRoute