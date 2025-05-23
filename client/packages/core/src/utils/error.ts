import axios from 'axios'

export interface ErrorOptions {
  readonly code?: string
}

/** Base error class for all errors generated. */
export class BaseError extends Error {
  /** Error code. */
  readonly code: string = this.constructor.name

  /** Whether the error can be recovered */
  readonly canRecover?: boolean

  /**
   * Constructor.
   *
   * @param message the error message
   * @param previous the previous error
   * @param code
   * @param extra
   */
  constructor(
    message: string,
    readonly previous?: Error,
    code?: string,
    readonly extra?: Record<string, any>,
    canRecover?: boolean,
  ) {
    super(message)

    this.code = code ?? this.code
    this.canRecover = canRecover ?? false

    // Take properties from previous error if it is BaseError
    if (previous) {
      this.stack += '\nFrom previous ' + previous.stack
    }

    Object.setPrototypeOf(this, new.target.prototype)
  }
}

export function getErrorMessage(error: unknown, context?: string): string | any {
  const contextText = `[${context ?? '-'}] `

  if (typeof error === 'string') {
    return `${contextText}${error}`
  }

  if (axios.isAxiosError(error)) {
    if (error.response) {
      return `${contextText}Axios response error: ${error.message} -> ${
        error.response?.statusText
      } [${error.response?.statusText}] ${JSON.stringify(error.response?.data)}`
    }

    return `${contextText}${error.message}`
  }

  return error
}
