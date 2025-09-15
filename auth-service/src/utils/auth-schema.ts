import {z} from "zod"

const emailSchema = z.string().min(1).max(255)
const passwordSchema = z.string().min(6).max(255)

export const loginSchema = z.object({
    email: emailSchema,
    password: passwordSchema
})

export const registerSchema = loginSchema.extend({
    name: z.string()
})