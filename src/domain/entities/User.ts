export type User = {
    id: number
    username: string // must be unique
    email: string // must be unique
    phone: string
    fullname: string
    hashedPassword: string
    isConfirmed: boolean // user confirmed account creation
    isActive: boolean // user is available
    isLocked: boolean // user do a fault and should be locked
    refreshCode: string // code for refresh token
    recoveryCode: number // code to recovery your account (e.g. missing password)
    intentsRecover: number // number of intents to refresh/recover
}