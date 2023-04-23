import { Langu, Messages } from "../../application/resources";
import { ILocalizer } from "./ILocalizer";

export class Localizer implements ILocalizer {
    definedLangu: Langu
    constructor(langu: Langu) {
        this.definedLangu = langu ?? Langu.EN;
    }

    m(key: string): string {
        const m = Messages.get(key)
        if (m == null) {
            return key + ": " + key
        } else {
            return key + ": " + m.find(e => e.langu == this.definedLangu)?.message
        }
    }

}