export type FontWeight = 100 | 200 | 300 | 400 | 500 | 600 | 700 | 800 | 900

export interface ThemeFontForWeight {
  readonly normal: string
  readonly italic?: string
}

export interface ThemeFont {
  readonly weights: Partial<Record<FontWeight, ThemeFontForWeight>> & {
    readonly default: ThemeFontForWeight
  }
}
