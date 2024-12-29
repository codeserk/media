import { Theme } from './modules/theme/theme.types'

// Default theme - should be overridden
export const defaultTheme = {} as const satisfies Theme
export type DefaultTheme = typeof defaultTheme
