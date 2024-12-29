import { computed, Signal } from '@preact-signals/safe-react'
import { useMemo } from 'react'

import { combineObjectPath, getObjectValueFromPath, ObjectPath } from '../../utils/object'
import { DeepSignal, SignalOrValue } from './signals.types'

/**
 * Gets readonly signals for all the keys of a given object.
 * Useful when we want to extract getters from an original signal.
 * @param signal
 * @returns
 */
export function deepSignal<TObject extends object>(
  signal: SignalOrValue<TObject>,
  path?: ObjectPath<TObject>,
): DeepSignal<TObject> {
  const rootValue = path ? getObjectValueFromPath(signalValue(signal), path) : signalValue(signal)
  if (rootValue === undefined || typeof rootValue !== 'object') {
    return {} as DeepSignal<TObject>
  }

  const keys = Object.keys(rootValue as object)
  if (
    rootValue &&
    !['Object', 'Array'].includes(rootValue?.constructor.name) &&
    rootValue?.constructor.name
  ) {
    const constructorKeys = Object.getOwnPropertyNames(rootValue.constructor.prototype)

    for (const key of constructorKeys) {
      if (key === 'constructor' || keys.includes(key)) {
        continue
      }

      keys.push(key)
    }
  }

  const result = keys.reduce((result, key) => {
    const subPath = combineObjectPath<TObject>(path, key)
    const subValue = getObjectValueFromPath(signalValue(signal), subPath)

    if (typeof subValue === 'object' && !Array.isArray(subValue)) {
      result[key] = deepSignal(signal, subPath)
    } else {
      result[key] = computed(() => getObjectValueFromPath(signalValue(signal), subPath))
    }

    return result
  }, {} as any)

  return {
    ...result,
    get value() {
      return signalValue(signal)
    },
  }
}

export function useDeepSignal<TObject extends object>(
  signal: SignalOrValue<TObject>,
): DeepSignal<TObject> {
  return useMemo(() => deepSignal(signal), [])
}

export function signalValue<TValue>(signalOrValue: SignalOrValue<TValue>): TValue {
  if (signalOrValue instanceof Signal) {
    return signalOrValue.value
  }

  return signalOrValue as TValue
}
