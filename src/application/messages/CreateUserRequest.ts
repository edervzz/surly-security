export class CreateUserRequest {
    email: string
    fullname: string
    hashedPassword: string

    constructor(email: string, fullname: string, hashedPassword: string) {
        this.email = email;
        this.fullname = fullname;
        this.hashedPassword = hashedPassword;
    }

}