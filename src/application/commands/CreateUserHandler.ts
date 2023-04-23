import { CreateUserRequest, CreateUserResponse } from "../messages";
import { ILocalizer } from "../../shared/abstractions";
import { CreateUserValidator } from "../validators"


export class CreateUserHandler implements IRequestHandler<CreateUserRequest, CreateUserResponse> {
    localizer: ILocalizer;
    constructor(
        localizer: ILocalizer
    ) {
        this.localizer = localizer;
    }

    async Handle(request: CreateUserRequest): Promise<CreateUserResponse> {
        // 1. request validator
        const validator = new CreateUserValidator(this.localizer);
        await validator.ValidAndThrow(request)

        return new CreateUserResponse();

    }
}