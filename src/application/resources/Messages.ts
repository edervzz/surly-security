import { CoreMesages } from './CoreMessages'
import { Message } from './typings/Message'
import { Langu } from './typings/Langu'

export const Messages = new Map<string, Message[]>([
    [CoreMesages.USER_EMAIL_EMPTY, [
        { langu: Langu.EN, message: "'Email' must not be empty." },
        { langu: Langu.ES, message: "'Correo' es requerido." }],
    ],
    [CoreMesages.USER_EMAIL_MATCH, [
        { langu: Langu.EN, message: "'Email' dont match." },
        { langu: Langu.ES, message: "'Correo' no hace match." }],
    ],
    [CoreMesages.USER_FULLNAME_EMPTY, [
        { langu: Langu.EN, message: "'User' must not be empty." },
        { langu: Langu.ES, message: "'Usuario' es requerido." }],
    ],
    [CoreMesages.USER_FULLNAME_LENGTH, [
        { langu: Langu.EN, message: "'User' must have between 5 and 50 characters." },
        { langu: Langu.ES, message: "'Usuario' debe tener entre 5 a 50 carácteres." }],
    ],
    [CoreMesages.USER_PASSWORD_EMPTY, [
        { langu: Langu.EN, message: "'Password' must not be empty." },
        { langu: Langu.ES, message: "'Constraseña' es requerida." }],
    ],
    [CoreMesages.USER_PASSWORD_MATCH, [
        { langu: Langu.EN, message: "'Password' must have one upper, one lower, one number and one special character @$!%*#?&" },
        { langu: Langu.ES, message: "'Constraseña' debe tener una mayúscula, una minúscula, un número y un carácter especial @$!%*#?&" }],
    ],
]);
