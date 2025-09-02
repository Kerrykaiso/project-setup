class AppError extends Error{
    public statusCode:number
    public status:string
    public isOperational: boolean
    constructor(message:string, statusCode: number, status:string){
    super(message)
    this.message = message
    this.status =status
    this.statusCode = statusCode
    this.isOperational = true
    Object.setPrototypeOf(this, new.target.prototype)
  }
}

export default AppError