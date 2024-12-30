import { Theme } from 'view/modules/theme/theme.types'

export const theme = {
  fontFamily: {
    Hero: {
      weights: {
        default: { normal: '' },
      },
    },
  },
} as const satisfies Theme

export type AppTheme = typeof theme
