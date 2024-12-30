import { FC } from 'react'

import { View, ViewProps } from './View'

export const HStack: FC<ViewProps> = (props) => {
  return <View flexDirection="row" {...props} />
}
