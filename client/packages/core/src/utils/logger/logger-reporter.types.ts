export interface LoggerReporter {
  /**
   * Reports an error
   * @param context
   * @param error
   */
  onError(context: string, error: unknown): Promise<void>
}
