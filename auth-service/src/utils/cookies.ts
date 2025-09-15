
 import type {CookieOptions,Response} from "express"
  const fifteenMinutes = new Date(Date.now()+15*60*1000)
  const thirtyDays = new Date(Date.now()+30*24*60*60*1000)

 const REFRESH_PATH = "auth/refresh-token"
  const defaultCookieOptions:CookieOptions ={
    httpOnly: true,
    secure: process.env.NODE_ENV === "production",
    sameSite:"strict"
    }
export const setAuthCookie = async(res:Response,accessToken:string,refreshToken?:string)=>{
     const accessCookieExpiresIn = fifteenMinutes 
     const refreshCookieExpiresIn = thirtyDays
    if (refreshToken) {
      res.cookie("refreshToken",refreshToken,{...defaultCookieOptions,expires:refreshCookieExpiresIn,path:REFRESH_PATH})
    .cookie("accessToken",accessToken,{...defaultCookieOptions,expires:accessCookieExpiresIn})
    }
    res.cookie("accessToken",accessToken,{...defaultCookieOptions,expires:accessCookieExpiresIn})

}

export const clearCookies =(res:Response)=>{
 res.clearCookie("accessToken").clearCookie("refreshToken",{path:REFRESH_PATH})
}