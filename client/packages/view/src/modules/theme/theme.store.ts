import { createContext, useContext } from 'react'
import { ViewStyle } from 'react-native'

import { compileThemeViewStyle } from '../view/view.service'
import { StyleDimensions, Theme, ThemeStyle } from './theme.types'

interface Props<TTheme extends Theme> {
  readonly theme: TTheme
}

export function useThemeStore<TTheme extends Theme>(props: Props<TTheme>) {
  // State

  const compileViewStyle = (
    style: ThemeStyle<TTheme, ViewStyle, StyleDimensions<ViewStyle>>,
  ): ViewStyle => {
    return compileThemeViewStyle(props.theme, style)
  }

  return {
    theme: props.theme,

    compileViewStyle,
  }
}

export type ThemeStore = ReturnType<typeof useThemeStore>

export const ThemeStoreContext = createContext<ThemeStore>({
  theme: {},
  compileViewStyle: (style) => style as ViewStyle,
})

export function useThemeStoreContext() {
  return useContext(ThemeStoreContext)
}
