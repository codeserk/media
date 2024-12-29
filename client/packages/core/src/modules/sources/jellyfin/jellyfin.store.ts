import { Api, Jellyfin } from '@jellyfin/sdk'
import { AudioApi } from '@jellyfin/sdk/lib/generated-client/api/audio-api'
import { UniversalAudioApi } from '@jellyfin/sdk/lib/generated-client/api/universal-audio-api'
import { BaseItemDto } from '@jellyfin/sdk/lib/generated-client/models'
import {
  getAudioApi,
  getImageApi,
  getItemsApi,
  getLibraryApi,
  getUniversalAudioApi,
} from '@jellyfin/sdk/lib/utils/api'
import { ImageUrlsApi } from '@jellyfin/sdk/lib/utils/api/image-urls-api'
import { useSignal } from '@preact-signals/safe-react'
import { createContext, useContext, useEffect } from 'react'

export function useJellyfinSourceStore() {
  const client = useSignal(
    new Jellyfin({
      clientInfo: { name: 'My Client Application', version: '1.0.0' },
      deviceInfo: { name: 'Device Name', id: 'unique-device-id' },
    }),
  )

  const servers = useSignal<any[]>([])
  const api = useSignal<Api | undefined>(undefined)
  const collections = useSignal<BaseItemDto[]>([])
  const items = useSignal<any[]>([])
  const audios = useSignal<UniversalAudioApi | undefined>(undefined)

  const images = useSignal<ImageUrlsApi | undefined>(undefined)

  const getServers = async () => {
    // const result = await client.value.discovery.getRecommendedServerCandidates('192.168.2.250')
    // const best = client.value.discovery.findBestServer(result)
    // if (!best) {
    //   return result
    // }

    // servers.value = result
    api.value = client.value.createApi('http://192.168.2.250:8096')

    // Fetch the public system info
    // const info = await getSystemApi(api.value).getPublicSystemInfo()
    // console.log('Info =>', info.data)

    // Fetch the list of public users
    // const users = getUserApi(api.value)

    // console.log('Users =>', users.data)

    // A helper method for authentication has been added to the SDK because
    // the default method exposed in the generated Axios client is rather
    // cumbersome to use.
    const auth = await api.value.authenticateUserByName('codeserk', '')

    // Authentication state is stored internally in the Api class, so now
    // requests that require authentication can be made normally
    const media = getLibraryApi(api.value)
    const libraries = await media.getMediaFolders()

    images.value = getImageApi(api.value)

    collections.value = libraries.data.Items?.filter((it) => it.CollectionType === 'music') ?? []
    // console.log('Libraries =>', libraries.data)

    if (collections.value[0]) {
      const itemsApi = getItemsApi(api.value)
      const response = await itemsApi.getItems({
        userId: auth.data.User?.Id,
        parentId: collections.value[0].Id,
        recursive: true,
        // mediaTypes: ['Audio'],
      })
      items.value = response.data.Items ?? []
    }

    audios.value = getUniversalAudioApi(api.value)

    return []
  }

  const getImageUrl = (item: BaseItemDto): string | undefined => {
    if (!images.value) {
      return
    }
    const response = images.value.getItemImageUrl(item)
    // console.log('response is', response)

    return response
  }

  const getAudioStream = async (id: string): Promise<File | undefined> => {
    if (!audios.value) {
      return
    }

    // audios.value.headUniversalAudioStream()
    const response = await audios.value.getUniversalAudioStream({
      itemId: id,
      maxStreamingBitrate: 140000000,
      container: [
        'opus',
        'webm|opus',
        'mp3',
        'aac',
        'm4a|aac',
        'm4b|aac',
        'flac',
        'webma',
        'webm|webma',
        'wav',
        'ogg',
      ],
      transcodingContainer: 'ts',
      transcodingProtocol: 'hls',
      audioCodec: 'ac',
      enableRedirection: true,
      enableRemoteMedia: false,
    })

    return response.data
  }

  useEffect(() => {
    getServers()
  }, [])

  return {
    client,
    servers,
    collections,
    items,

    getServers,
    getImageUrl,
    getAudioStream,
  }
}

export type JellyfinSourceStore = ReturnType<typeof useJellyfinSourceStore>

export const JellyfinSourceStoreContext = createContext<JellyfinSourceStore>({
  isLoading: true,
} as any)

export function useJellyfinSourceStoreContext() {
  return useContext(JellyfinSourceStoreContext)
}
