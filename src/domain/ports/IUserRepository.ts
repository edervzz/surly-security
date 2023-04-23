import { User } from "../entities";

interface IUserRepository extends
    ICreateRepository<User>,
    IReadSingleRepository<number, User>,
    IReadByExternalIDRepository<string, User> {
}