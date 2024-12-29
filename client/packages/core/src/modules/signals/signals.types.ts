import { ReadonlySignal } from '@preact-signals/safe-react'

export type DeepSignal<TObject extends object> = {
  [TKey in keyof TObject]: TObject[TKey] extends object
    ? DeepSignal<TObject[TKey]>
    : ReadonlySignal<TObject[TKey]>
} & {
  value: TObject
}

export type SignalOrValue<TValue> = ReadonlySignal<TValue> | TValue
