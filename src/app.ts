import { ValidationError } from "joi";
import { CreateUserHandler } from "./application/commands";
import { CreateUserRequest } from "./application/messages";
import { Langu } from "./application/resources/typings/Langu";
import { Localizer } from "./shared/abstractions";

const localizer = new Localizer(Langu.ES);
let handler: CreateUserHandler = new CreateUserHandler(localizer)

let request = new CreateUserRequest("eder@gmail.com", "eder velazquez", "Eder123");
handler.Handle(request)
    .then(() => console.log("OK"))
    .catch((err: ValidationError) => console.log(err));

console.log("end_test");
