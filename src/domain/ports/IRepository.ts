interface ICreateRepository<TEntity> {
    Create(entity: TEntity): void
}

interface IReadSingleRepository<TKey, TEntity> {
    Read(key: TKey): TEntity
}

interface IReadAllRepository<TEntity> {
    ReadAll(): TEntity[]
}

interface IReadByExternalIDRepository<TExtID, TEntity> {
    ReadByExternalID(extID: TExtID): TEntity
}

interface IReadByParentIDRepository<TParentID, TEntity> {
    ReadByExternalID(parentID: TParentID): TEntity[]
}
