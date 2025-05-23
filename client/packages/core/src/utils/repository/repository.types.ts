import { ReadonlySignal } from '@preact-signals/safe-react'

export interface Entity<K> {
  readonly id: K
}
export interface RepositoryStore<TEntity extends Entity<string>> {
  readonly items: ReadonlySignal<TEntity[]>
  readonly itemById: (id?: string) => ReadonlySignal<TEntity | undefined>
  readonly itemsByIds: (ids: string[]) => ReadonlySignal<TEntity[]>

  addOne: (item: TEntity, clear?: boolean) => void
  addMany: (items: TEntity[] | Record<string, TEntity>, clear?: boolean) => void
  removeOne: (id: string) => void
  clear: () => void
}
