/* eslint-disable no-console */

import { getErrorMessage } from '../error'
import { LoggerReporter } from './logger-reporter.types'

let reporter: LoggerReporter

/**
 * Sets the logger reporter
 * @param newReporter
 */
export function setLoggerReporter(newReporter: LoggerReporter) {
  reporter = newReporter
}

export class Logger {
  /**
   * Constructor.
   * @param context
   */
  constructor(readonly context: string) {}

  /**
   * Debugs a message, only in development
   * @param message
   * @param data
   */
  debug(message: string, ...data: unknown[]) {
    if (process.env.NODE_ENV === 'development') {
      return this.info(message, ...data)
    }
  }

  /**
   * Logs a message
   * @param message
   * @param data
   */
  info(message: string, ...data: unknown[]) {
    console.log(this.context, message, ...data)
  }

  /**
   * Warns a message
   * @param message
   * @param data
   */
  warn(message: string, ...data: unknown[]) {
    console.warn(this.context, message, ...data)
  }

  /**
   * Logs an error
   * @param error
   */
  error(error: unknown, ...data: unknown[]) {
    try {
      const extra = this.extractExtraFromError(error)
      const message = getErrorMessage(error, this.context)
      console.error(message, extra, ...data)

      reporter?.onError(this.context, error)
    } catch (error) {
      //
    }
  }

  /**
   * Extracts extra information from the error
   * @param error
   * @returns extra information
   */
  protected extractExtraFromError(error: unknown): Record<string, any> {
    // Response data
    if (typeof error === 'object' && (error as any).response?.data) {
      return { response: (error as any).response.data }
    }

    return {}
  }
}
