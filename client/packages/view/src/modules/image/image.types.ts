import { ImageStyle as RNImageStyle } from 'react-native'

import { DefaultTheme } from '../../theme'
import { StyleDimensions, StyleDimensionsValue, Theme } from '../theme/theme.types'
import { flexStyleDimensions, shadowStyleDimensions } from '../view/view.types'

export const imageStyleDimensions = {
  ...flexStyleDimensions,
  ...shadowStyleDimensions,

  // Image
  borderRadius: 'borderRadius',
  borderTopLeftRadius: 'borderRadius',
  borderTopRightRadius: 'borderRadius',
  borderBottomLeftRadius: 'borderRadius',
  borderBottomRightRadius: 'borderRadius',
  backgroundColor: 'color',
  borderColor: 'color',
  overlayColor: 'color',
  tintColor: 'color',
} as const satisfies StyleDimensions<RNImageStyle>

export type ThemeViewStyle<TTheme extends Theme = DefaultTheme> = {
  [K in keyof RNImageStyle]?: StyleDimensionsValue<
    TTheme,
    RNImageStyle,
    typeof imageStyleDimensions,
    K
  >
}
