import Joi from "joi";
import { CreateUserRequest } from "../messages";
import { ILocalizer } from '../../shared/abstractions'
import { AbstractValidator } from "../../shared/abstractions/AbstractValidator";
import { CoreMesages } from "../resources";

export class CreateUserValidator extends AbstractValidator<CreateUserRequest> {

    constructor(localizer: ILocalizer) {
        super(
            {
                email: Joi.string()
                    .required()
                    .email()
                    .messages({
                        "string.empty": localizer.m(CoreMesages.USER_EMAIL_EMPTY),
                        "string.email": localizer.m(CoreMesages.USER_EMAIL_MATCH),
                    }),
                fullname: Joi.string()
                    .required()
                    .min(5)
                    .max(50)
                    .messages({
                        "string.min": localizer.m(CoreMesages.USER_FULLNAME_LENGTH),
                        "string.max": localizer.m(CoreMesages.USER_FULLNAME_LENGTH),
                        "string.empty": localizer.m(CoreMesages.USER_FULLNAME_EMPTY),
                    }),
                hashedPassword: Joi.string()
                    .required()
                    .pattern(new RegExp('^(?=.*[@$!%*#?&])(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$'))
                    .messages({
                        "string.empty": localizer.m(CoreMesages.USER_PASSWORD_EMPTY),
                        "string.pattern.base": localizer.m(CoreMesages.USER_PASSWORD_MATCH)
                    }),
            }
        );
    }
}
