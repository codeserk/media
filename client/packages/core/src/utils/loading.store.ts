import { computed, useComputed } from '@preact-signals/safe-react'

import { Logger } from './logger/logger'
import { useMapStore } from './map/map.store'

// eslint-disable-next-line @typescript-eslint/ban-types
export type LoaderId = string | number | Function | { code: string }

const logger = new Logger('loading')

/**
 * Store to keep loaders.
 */
export function useLoadingStore() {
  // State

  const loadersMap = useMapStore<string, boolean>()

  // Getters

  const loaders = useComputed(() =>
    Object.entries(loadersMap.values.value)
      .filter(([, value]) => value)
      .map((key) => key),
  )

  /** Whether any loader is active. */
  const isLoading = useComputed(() => loaders.value.length > 0)
  const isLoadingKey = (id?: LoaderId) =>
    computed(() => (id ? (loadersMap.valueByKey(getId(id)).value ?? false) : false))

  // Actions

  const startLoading = (id: LoaderId) => {
    loadersMap.addOne(getId(id), true)
  }
  const stopLoading = (id: LoaderId) => {
    loadersMap.removeOne(getId(id))
  }
  const clearLoaders = () => {
    loadersMap.clear()
  }

  const getId = (id: LoaderId): string => {
    if (!id) {
      return 'unknown'
    }
    if (typeof id === 'string') {
      return id
    }
    if (typeof id === 'function') {
      if ((id as any).code) {
        return (id as any).code
      }
      if (id.name) {
        return id.name
      }
    }

    return id?.toString()
  }

  function operation<TArgs extends any[], TResult>(
    fn: Func<TArgs, Promise<TResult>>,
  ): (...args: TArgs) => Promise<TResult> {
    const id = getId(fn)
    const returnFn = async (...args: TArgs): Promise<TResult> => {
      startLoading(id)
      for (const arg of args) {
        if (typeof arg === 'string') {
          startLoading(arg as string)
        }
      }

      try {
        return await fn(...args)
      } catch (error) {
        logger.error(error)
        throw error
      } finally {
        stopLoading(id)
        for (const arg of args) {
          if (typeof arg === 'string') {
            stopLoading(arg as string)
          }
        }
      }
    }

    returnFn.code = id
    return returnFn
  }

  return {
    isLoading,
    isLoadingKey,
    loaders,
    startLoading,
    stopLoading,
    clearLoaders,

    operation,
  }
}

export type LoadingStore = ReturnType<typeof useLoadingStore>

type Func<TArgs extends any[], TResult> = (...args: TArgs) => TResult
