export interface AuthPayload {
    userId:string
    sessionId:string
    iat?: number
    exp?: number
}