import Joi from "joi";

export abstract class AbstractValidator<T> {
    private _schema: Joi.ObjectSchema<any>;

    constructor(r: any) {
        this._schema = Joi.object(r);
    }

    async ValidAndThrow(request: T): Promise<void> {
        await this._schema.validateAsync(request);
    }


}
