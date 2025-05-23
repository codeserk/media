import { ReadonlySignal } from '@preact-signals/safe-react'

export interface MapStore<TKey extends string | number, TValue> {
  readonly map: ReadonlySignal<Record<TKey, TValue>>
  readonly entries: ReadonlySignal<[TKey, TValue][]>
  readonly values: ReadonlySignal<TValue[]>
  readonly valueByKey: (key?: TKey) => ReadonlySignal<TValue | undefined>
  readonly valuesByKeys: (keys: TKey[]) => ReadonlySignal<TValue[]>

  addOne: (key: TKey, value: TValue, clear?: boolean) => void
  addMany: (values: Record<TKey, TValue>, clear?: boolean) => void
  removeOne: (key: TKey) => void
  clear: () => void
}
